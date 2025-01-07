package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
)
func GetCitas(c *fiber.Ctx) error {
	db := database.DB
	var citas []models.Citas
	db.Find(&citas)
	return c.JSON(citas)
}

func GetCitaById(c *fiber.Ctx) error {	
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var cita models.Citas
	buscar := database.DB.First(&cita, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cita no encontrada",
		})
	}
	return c.JSON(cita)
}

func CreateCita(c *fiber.Ctx) error {		
	var cita = new(models.Citas)	
	if err := c.BodyParser(cita); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error al parsear el body",
		})
	}
	fmt.Println("cita",cita)
	guardar := database.DB.Create(&cita)
	if  guardar.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "No se pudo crear la cita",
		})
		
	}
	return c.JSON(cita)
}

func DeleteCita(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "id invalido",
		})
	}
	var cita models.Citas
	buscar := database.DB.First(&cita, id)
	if buscar.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cita no encontrada",
		})
	}
	database.DB.Delete(&cita)
	return c.JSON(fiber.Map{
		"message": "Cita eliminada",
	})
}