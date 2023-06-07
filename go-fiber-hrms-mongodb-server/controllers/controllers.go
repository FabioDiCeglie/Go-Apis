package controllers

import (
	"github.com/fabio/go-fiber-mongo/database"
	"github.com/fabio/go-fiber-mongo/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetEmployees(c *fiber.Ctx) {
	query := bson.D{{}}

	cursor, err := database.Mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var employees []models.Employee = make([]models.Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}
