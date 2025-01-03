package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
	"github.com/jnestormg/GoLang.git/routes"
)

func main(){
	api := fiber.New()

	database.Connect()
	database.DB.AutoMigrate(&models.Medicos{}, &models.Pacientes{}, &models.Especialidades{})

	routes.PacientesRoutes(api)

	api.Listen(":3000")
}