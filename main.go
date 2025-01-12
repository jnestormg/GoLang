package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jnestormg/GoLang.git/database"
	_ "github.com/jnestormg/GoLang.git/docs"
	"github.com/jnestormg/GoLang.git/models"
	"github.com/jnestormg/GoLang.git/routes"
)

// @title API de citas médicas
// @version 1.0
// @description API para gestionar citas médicas.
// @host localhost:3000
// @contact.name API Support
// @contact.email fiber@swagger.io
// @BasePath /

func main() {
	api := fiber.New()

	api.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, DELETE, PUT, OPTIONS",
	}))

	api.Get("/swagger/*", swagger.HandlerDefault)

	database.Connect()
	err := database.DB.AutoMigrate(
		&models.Medicos{},
		&models.Pacientes{},
		&models.Especialidades{},
		&models.Citas{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	routes.PacientesRoutes(api)
	routes.MedicosRoutes(api)
	routes.EspecialidadesRoutes(api)
	routes.CitasRoutes(api)

	api.Listen(":3000")
}
