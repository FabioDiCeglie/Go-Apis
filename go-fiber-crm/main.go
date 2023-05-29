package main

import (
	"fmt"

	"github.com/fabio/go-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setUpRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var db *gorm.DB
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate()
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	fmt.Print("Starting server to http://localhost:3000")
	defer database.DBConn.Close()
}
