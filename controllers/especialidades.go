package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
)

func GetEspecialidades(c *fiber.Ctx) error {
	var especialidades []models.Especialidades
	database.DB.Find(&especialidades)
	return c.JSON(especialidades)
}

func GetEspecialidadById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var especialidad models.Especialidades
	buscar := database.DB.First(&especialidad, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Especialidad no encontrada",
		})
	}
	return c.JSON(especialidad)
}

func CreateEspecialidad(c *fiber.Ctx) error {
	var especialidad = new(models.Especialidades)	

	parsear := c.BodyParser(especialidad)
	if parsear != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al parsear el body",
		})
	}
	guardar := database.DB.Create(&especialidad)
	if guardar.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "No se pudo crear la especialidad",
		})
	}
	return c.JSON(especialidad)
}

func DeleteEspecialidad(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var especialidad models.Especialidades
	buscar := database.DB.First(&especialidad, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Especialidad no encontrada",
		})
	}
	database.DB.Delete(&especialidad)
	return c.JSON(fiber.Map{
		"message": "Especialidad eliminada",
	})
}