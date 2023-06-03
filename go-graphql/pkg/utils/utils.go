package utils

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/fabio/graphql/database"
	"github.com/fabio/graphql/graph/model"
	"github.com/jinzhu/gorm"
)

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIdByUsername checks if a user exists in the database by the given username
func GetUserIdByUsername(username string) (int, error) {
	db := database.Db
	var user model.User
	if err := db.Where("Name = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("No user found with name: ", username)
			return 0, nil
		}
		return 0, err
	}

	num, err := strconv.Atoi(user.ID)
	if err != nil {
		return 0, err
	}

	return num, nil
}
