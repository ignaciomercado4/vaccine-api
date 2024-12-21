package main

import (
	"log"
	"net/http"
	"vaccine-api/database"
	"vaccine-api/handlers"

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

	userHandler := handlers.UserHandler{DB: db}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	router.POST("/createUser", handlers.CreateUser(&userHandler))

	router.Run()
}
