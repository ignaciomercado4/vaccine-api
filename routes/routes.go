package routes

import (
	"net/http"
	"vaccine-api/handlers"
	"vaccine-api/middleware"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, appHandler handlers.AppHandler) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	api := router.Group("/api")
	{
		api.POST("/signup", handlers.CreateUser(&appHandler))
		api.POST("/login", handlers.Login(&appHandler))

		drugs := api.Group("/drugs")
		{
			drugs.POST("/", middleware.AuthMiddleware(), handlers.CreateDrug(&appHandler))
			drugs.GET("/", middleware.AuthMiddleware(), handlers.GetDrugs(&appHandler))
			drugs.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteDrug(&appHandler))
			drugs.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateDrug(&appHandler))
		}

		vaccination := api.Group("/vaccination")
		{
			vaccination.POST("/vaccination", middleware.AuthMiddleware(), handlers.CreateVaccination(&appHandler))
			vaccination.GET("/vaccination", middleware.AuthMiddleware(), handlers.GetVaccinations(&appHandler))
			vaccination.DELETE("/vaccinaction/:id", middleware.AuthMiddleware(), handlers.DeleteVaccination(&appHandler))
		}
	}
}
