package handlers

import (
	"net/http"
	"vaccine-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DrugHandler struct {
	DB *gorm.DB
}

func CreateDrug(h *DrugHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newDrug models.Drug

		if err := ctx.ShouldBind(&newDrug); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		result := h.DB.Create(&newDrug)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error,
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"drug": newDrug,
		})
	}
}
