package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Db *gorm.DB
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func Connect() {
	dsn := "host=mouse.db.elephantsql.com user=ydgmnplr password=jMKkqF5eKoRj5MrTnJRAw6zpXGGQPMcn dbname=ydgmnplr port=5432"
	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection opened to database")
	Db = d
}

// func GetDB() *gorm.DB {
// 	return db
// }

func InitDatabase() {
	Connect()
	Db.AutoMigrate(&Lead{})
	fmt.Println("Database Migrated")
}
