package handlers

import (
	"net/http"
	"vaccine-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserHandler struct {
	DB *gorm.DB
}

func CreateUser(h *UserHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User

		if err := ctx.ShouldBind(&newUser); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		h.DB.Create(&newUser)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "user created",
			"user":   newUser,
		})
	}
}
