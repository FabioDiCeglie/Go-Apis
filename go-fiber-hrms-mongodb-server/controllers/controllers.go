package controllers

import (
	"github.com/fabio/go-fiber-mongo/database"
	"github.com/fabio/go-fiber-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetEmployees(c *fiber.Ctx) {
	query := bson.D{bson.E{}}

	cursor, err := database.Mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		c.Status(500).Send(err.Error())
		return
	}

	var employees []models.Employee = make([]models.Employee, 0)

	if err := cursor.All(c.Context(), &employees); err != nil {
		c.Status(500).Send(err.Error())
		return
	}

	c.JSON(employees)
}

func UpdateEmployee(c *fiber.Ctx) {
	idParam := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.SendStatus(400)
		return
	}

	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		c.Status(400).SendString(err.Error())
		return
	}

	query := bson.D{bson.E{Key: "_id", Value: employeeID}}
	update := bson.D{
		bson.E{Key: "$set",
			Value: bson.D{
				bson.E{Key: "name", Value: employee.Name},
				bson.E{Key: "age", Value: employee.Age},
				bson.E{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = database.Mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.SendStatus(400)
			return
		}
		c.SendStatus(500)
		return
	}

	employee.ID = idParam

	c.Status(200).JSON(employee)
}

func CreateEmployee(c *fiber.Ctx) {
	collection := database.Mg.Db.Collection("employees")

	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		c.Status(400).SendString(err.Error())
		return
	}

	employee.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		c.Status(500).SendString(err.Error())
		return
	}

	filter := bson.D{bson.E{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	c.Status(201).JSON(createdEmployee)
}

func DeleteEmployee(c *fiber.Ctx) {
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		c.SendStatus(400)
		return
	}

	query := bson.D{bson.E{Key: "_id", Value: employeeID}}
	result, err := database.Mg.Db.Collection("employees").DeleteOne(c.Context(), &query)

	if err != nil {
		c.SendStatus(500)
		return
	}

	if result.DeletedCount < 1 {
		c.SendStatus(404)
		return
	}

	c.Status(200).JSON("record deleted")
}
