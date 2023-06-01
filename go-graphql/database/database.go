package database

import (
	"fmt"

	"github.com/fabio/graphql/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Db *gorm.DB
)

func Connect() {
	dsn := "host=mouse.db.elephantsql.com user=ydgmnplr password=jMKkqF5eKoRj5MrTnJRAw6zpXGGQPMcn dbname=ydgmnplr port=5432"
	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection opened to database")
	Db = d
}

func InitDB() {
	Connect()
	Db.AutoMigrate(&model.Link{}, &model.User{})
	fmt.Println("Database Migrated")
}

func CloseDB() error {
	return Db.Close()
}
