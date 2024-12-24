package main

import (
	"log"
	"vaccine-api/database"
	"vaccine-api/handlers"
	"vaccine-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	db := database.ConnectDatabase()

	database.MigrateModels(db)

	routes.GetRoutes(router, handlers.AppHandler{DB: db})

	router.Run()
}
