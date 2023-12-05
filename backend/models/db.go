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
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(dbURI)

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		panic(err)
	}

	DB.AutoMigrate(&Student{})
	DB.AutoMigrate(&Hostel{})
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Log{})
	DB.AutoMigrate(&Room{})
	fmt.Println("Database connected")
}
