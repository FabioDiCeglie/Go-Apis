package models

import (
	"github.com/FabioDiCeglie/Learning-Go/tree/main/bookstorage_management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name       string `gorm: ""json:"name"`
	Author     string `json:"author"`
	Pubication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
