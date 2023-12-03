package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func ConnectDatabase() {
	LoadEnv()

	var err error
	dbURI := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		panic(err)
	}

	DB.AutoMigrate(&Student{})
	fmt.Println("Database connected")
}
