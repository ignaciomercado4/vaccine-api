package database

import (
	"fmt"
	"log"
	"os"
	"vaccine-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	fmt.Println("DB_PORT:", DB_PORT)

	if DB_PORT == "" {
		log.Fatal("DB_PORT is not set in environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Argentina/Buenos_Aires",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to initialize database, got error: ", err.Error())
	}

	return db
}

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
