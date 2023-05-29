package routes

import (
	"github.com/gofiber/fiber"
)

func SetUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
