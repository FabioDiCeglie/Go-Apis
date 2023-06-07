package routes

import (
	"github.com/fabio/go-fiber-mongo/controllers"
	"github.com/gofiber/fiber"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/employee", controllers.GetEmployees)
	app.Post("/employee", controllers.CreateEmployee)
	app.Put("/employee/:id", controllers.UpdateEmployee)
	app.Delete("/employee/:id", controllers.DeleteEmployee)
}
