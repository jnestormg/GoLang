package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/controllers"
)
func CitasRoutes(app *fiber.App) {
	// @Summary Get all pacientes
	// @Description Get all pacientes from the database
	// @Tags Citas
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.Paciente "List of pacientes"
	// @Failure 500 {object} map[string]string "Internal Server Error"
	// @Router /citas [get]
	app.Post("/citas", controllers.CreateCita)
	app.Get("/citas", controllers.GetCitas)
	app.Get("/citas/:id", controllers.GetCitaById)
	app.Delete("/citas/:id", controllers.DeleteCita)
}