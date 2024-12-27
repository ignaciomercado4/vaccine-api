package handlers

import (
	"net/http"
	"vaccine-api/middleware"
	"vaccine-api/models"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(h *AppHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request data: " + err.Error(),
			})
			return
		}

		if result := h.DB.Create(&newUser); result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error creating user",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    newUser,
		})
	}
}

func Login(h *AppHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginReq LoginRequest
		if err := ctx.ShouldBindJSON(&loginReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var user models.User
		result := h.DB.Where("email = ?", loginReq.Email).First(&user)
		if result.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "User matching provided email not found",
			})
			return
		}

		if loginReq.Password != user.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}

		token, err := middleware.GenerateToken(user.ID, user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error generating token",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
			"user": gin.H{
				"id":    user.ID,
				"email": user.Email,
			},
		})
	}
}

func Logout(h *AppHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"token":   "",
			"message": "Logout successful",
		})
	}
}
