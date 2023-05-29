package routes

import (
	"github.com/fabio/go-fiber/controllers"
	"github.com/gofiber/fiber"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", controllers.GetLeads)
	app.Get("/api/v1/lead/:id", controllers.GetLead)
	app.Post("/api/v1/lead", controllers.NewLead)
	app.Delete("/api/v1/lead/:id", controllers.DeleteLead)
	app.Patch("/api/v1/lead/:id", controllers.UpdateLead)
}
