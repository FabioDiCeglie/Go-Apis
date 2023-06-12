package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(os.Getenv("DB_NAME"))

	if err != nil {
		return err
	}

	fmt.Println("Connection opened to database")
	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func InitDatabase() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database initialized!")
}
