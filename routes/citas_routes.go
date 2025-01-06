package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/controllers"
)
func CitasRoutes(app *fiber.App) {
	app.Post("/citas", controllers.CreateCita)
	app.Get("/citas", controllers.GetCitas)
	app.Get("/citas/:id", controllers.GetCitaById)
	app.Delete("/citas/:id", controllers.DeleteCita)
}