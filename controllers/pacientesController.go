package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
)


func GetPacientes(c *fiber.Ctx) error {
	var pacientes []models.Pacientes
	database.DB.Preload("Citas").Find(&pacientes)
	return c.JSON(pacientes)
}

func GetPAcienteById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var paciente models.Pacientes
	buscar := database.DB.Preload("Citas").First(&paciente, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Paciente no encontrado",
		})

	}
	return c.JSON(paciente)
}

func CreatePaciente(c *fiber.Ctx) error {
	var paciente = new(models.Pacientes)

	parsear := c.BodyParser(paciente)
	if parsear != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al parsear el body",
		})
	}
	guardar := database.DB.Create(&paciente)
	if guardar.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "No se pudo crear el paciente",
		})
	}
	return c.JSON(paciente)
}

func DeletePaciente(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var paciente models.Pacientes
	buscar := database.DB.First(&paciente, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Paciente no encontrado",
		})
	}
	database.DB.Delete(&paciente)
	return c.JSON(fiber.Map{
		"message": "Paciente eliminado",
	})
}
