package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
)

func GetMedicos(c *fiber.Ctx) error {
	var medicos []models.Medicos
	database.DB.Preload("Especialidades").Preload("Citas").Find(&medicos)
	return c.JSON(medicos)
}

func GetMedicoById(c *fiber.Ctx) error {	
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var medico models.Medicos
	buscar := database.DB.Preload("Especialidades").Preload("Citas").First(&medico, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Medico no encontrado",
		})
	}
	return c.JSON(medico)
}

func CreateMedico(c *fiber.Ctx) error {		
	var medico = new(models.Medicos)

	parsear := c.BodyParser(medico)
	if parsear != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al parsear el body",
		})
	}
	 guardar := database.DB.Create(&medico)
	 if guardar.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "No se pudo crear el medico",
		})
	 }
	return c.JSON(medico)
}

func DeleteMedico(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var medico models.Medicos
	buscar := database.DB.First(&medico, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Medico no encontrado",
		})
	}
	database.DB.Delete(&medico)
	return c.JSON(fiber.Map{
		"message": "Medico eliminado",
	})
}