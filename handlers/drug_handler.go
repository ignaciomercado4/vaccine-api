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
			return
		}

		result := h.DB.Create(&newDrug)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"drug": newDrug,
		})
	}
}

func GetDrugs(h *DrugHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var drugs []models.Drug

		result := h.DB.Find(&drugs)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"drugs": drugs,
		})
	}
}

func DeleteDrug(h *DrugHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		drugId := ctx.Param("id")

		var drug models.Drug
		result := h.DB.First(&drug, drugId)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error,
			})
			return
		}

		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Drug not found",
			})
			return
		}

		h.DB.Delete(&drug)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Drug deleted succesfully",
		})
	}
}

func UpdateDrug(h *DrugHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		drugId := ctx.Param("id")

		var updatedDrug models.Drug
		if err := ctx.ShouldBindJSON(&updatedDrug); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		var existingDrug models.Drug

		result := h.DB.First(&existingDrug, drugId)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Drug not found",
			})
			return
		}

		existingDrug.Name = updatedDrug.Name
		existingDrug.Approved = updatedDrug.Approved
		existingDrug.MinDose = updatedDrug.MinDose
		existingDrug.MaxDose = updatedDrug.MaxDose
		existingDrug.AvailableAt = updatedDrug.AvailableAt

		saveResult := h.DB.Save(&existingDrug)
		if saveResult.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": saveResult.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"drug": existingDrug,
		})
	}
}
