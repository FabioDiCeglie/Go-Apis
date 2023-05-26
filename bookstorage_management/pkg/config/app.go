package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db = d
}

func GetDB() *gorm.DB {
	return db
}
