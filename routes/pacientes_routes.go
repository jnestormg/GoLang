package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jnestormg/GoLang.git/controllers"
)

// PacientesRoutes define las rutas relacionadas con los pacientes
func PacientesRoutes(app *fiber.App) {
    // @Summary Get all pacientes
    // @Description Get all pacientes from the database
    // @Tags Pacientes
    // @Accept json
    // @Produce json
    // @Success 200 {array} models.Paciente "List of pacientes"
    // @Failure 500 {object} map[string]string "Internal Server Error"
    // @Router /pacientes [get]
    app.Get("/pacientes", controllers.GetPacientes)

    // @Summary Create a new paciente
    // @Description Create a new paciente in the database
    // @Accept json
    // @Produce json
    // @Param paciente body models.Paciente true "Paciente object"
    // @Success 201 {object} models.Paciente "Paciente created"
    // @Failure 400 {object} map[string]string "Bad Request"
    // @Failure 500 {object} map[string]string "Internal Server Error"
    // @Router /pacientes [post]
    app.Post("/pacientes", controllers.CreatePaciente)

    // @Summary Get a paciente by ID
    // @Description Get a specific paciente by their ID
    // @Accept json
    // @Produce json
    // @Param id path int true "Paciente ID"
    // @Success 200 {object} models.Paciente "Paciente found"
    // @Failure 404 {object} map[string]string "Paciente not found"
    // @Failure 500 {object} map[string]string "Internal Server Error"
    // @Router /pacientes/{id} [get]
    app.Get("/pacientes/:id", controllers.GetPAcienteById)

    // @Summary Delete a paciente by ID
    // @Description Delete a paciente from the database by their ID
    // @Accept json
    // @Produce json
    // @Param id path int true "Paciente ID"
    // @Success 200 {string} string "Paciente deleted"
    // @Failure 404 {object} map[string]string "Paciente not found"
    // @Failure 500 {object} map[string]string "Internal Server Error"
    // @Router /pacientes/{id} [delete]
    app.Delete("/pacientes/:id", controllers.DeletePaciente)
}
