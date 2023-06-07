package main

import (
	"fmt"

	"github.com/fabio/go-fiber-mongo/database"
	"github.com/fabio/go-fiber/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	database.InitDatabase()
	routes.SetUpRoutes(app)
	fmt.Println("Starting server to http://localhost:3000")
	app.Listen(3000)
}
