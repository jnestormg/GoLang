package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnestormg/GoLang.git/controllers"
)

func EspecialidadesRoutes(app *fiber.App) {
	app.Post("/especialidades", controllers.CreateEspecialidad)
	app.Get("/especialidades", controllers.GetEspecialidades)
	app.Get("/especialidades/:id", controllers.GetEspecialidadById)
	app.Delete("/especialidades/:id", controllers.DeleteEspecialidad)
}