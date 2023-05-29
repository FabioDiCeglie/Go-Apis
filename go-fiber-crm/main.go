package main

import (
	"fmt"

	"github.com/fabio/go-fiber/database"
	"github.com/fabio/go-fiber/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	database.initDatabase()
	routes.SetUpRoutes(app)
	app.Listen(3000)
	fmt.Println("Starting server to http://localhost:3000")
	defer database.DBConn.Close()
}
