package main

import (
	"fmt"

	"github.com/fabio/go-fiber/controllers"
	"github.com/fabio/go-fiber/database"
	"github.com/fabio/go-fiber/routes"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func initDatabase() {
	database.Connect()
	var db *gorm.DB = database.GetDB()
	db.AutoMigrate(&controllers.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.SetUpRoutes(app)
	app.Listen(3000)
	fmt.Println("Starting server to http://localhost:3000")
	defer database.DBConn.Close()
}
