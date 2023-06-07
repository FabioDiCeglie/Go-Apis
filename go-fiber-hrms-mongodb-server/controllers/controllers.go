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

func CreateEmployee(c *fiber.Ctx) {
	collection := database.Mg.Db.Collection("employees")

	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}
