package db

import (
	"log"
	user_model "smol-api/src/models/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClient(uri string) *gorm.DB {
	if uri == "" {
		log.Fatal("Database URI not provided")
	}
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	db.AutoMigrate(&user_model.User{})

	return db
}
