package main

import (
	"log"
	"net/http"
	"vaccine-api/database"
	"vaccine-api/handlers"
	"vaccine-api/middleware"

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
	drugHandler := handlers.DrugHandler{DB: db}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	router.POST("/signup", handlers.CreateUser(&userHandler))
	router.POST("/login", handlers.Login(&userHandler))

	router.POST("/drugs", middleware.AuthMiddleware(), handlers.CreateDrug(&drugHandler))
	router.GET("/drugs", middleware.AuthMiddleware(), handlers.GetDrugs(&drugHandler))
	router.DELETE("/drugs/:id", middleware.AuthMiddleware(), handlers.DeleteDrug(&drugHandler))

	router.Run()
}
