package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/controllers"
)

func MedicosRoutes(app *fiber.App) {
	app.Get("/medicos", controllers.GetMedicos)
	app.Post("/medicos", controllers.CreateMedico)	
	app.Get("/medicos/:id", controllers.GetMedicoById)
	app.Delete("/medicos/:id", controllers.DeleteMedico)
	
}