package config

import (
	"fmt"

	"github.com/fabio/go-fiber/controllers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=mouse.db.elephantsql.com user=ydgmnplr password=jMKkqF5eKoRj5MrTnJRAw6zpXGGQPMcn dbname=ydgmnplr port=5432"
	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection opened to database")
	db = d
}

// func GetDB() *gorm.DB {
// 	return db
// }

func initDatabase() {
	Connect()
	db.AutoMigrate(&controllers.Lead{})
	fmt.Println("Database Migrated")
}
