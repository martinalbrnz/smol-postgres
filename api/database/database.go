package database

import (
	"log"
	"os"
	"server/api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error handling .env file")
	}

	dsn := os.Getenv("DB_URI")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to DB. \n", err)
	}

	log.Println("Connected!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	db.AutoMigrate(&models.Movement{})

	DB = Dbinstance{
		Db: db,
	}
}
