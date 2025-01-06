package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/database"
	"github.com/jnestormg/GoLang.git/models"
	"github.com/jnestormg/GoLang.git/routes"
)

func main() {
    api := fiber.New()

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

    api.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello World",
        })
    })

    routes.PacientesRoutes(api)
    routes.MedicosRoutes(api)
    routes.EspecialidadesRoutes(api)

    api.Listen(":3000")
}