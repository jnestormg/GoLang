package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/controllers"
)

func PacientesRoutes(app *fiber.App) {
	app.Get("/pacientes", controllers.GetPacientes)
	app.Post("/pacientes", controllers.CreatePaciente)
	app.Get("/pacientes/:id", controllers.GetPAcienteById)
	app.Delete("/pacientes/:id", controllers.DeletePaciente)
}