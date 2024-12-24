package handlers

import (
	"net/http"
	"vaccine-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VaccinationHandler struct {
	DB *gorm.DB
}

func CreateVaccination(h *VaccinationHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newVaccination models.Vaccination

		if err := ctx.ShouldBind(&newVaccination); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var selectedDrug models.Drug
		result := h.DB.Find(&selectedDrug, newVaccination.DrugId)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
		if selectedDrug.ID == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Drug not found",
			})
			return
		}

		if newVaccination.Dose > selectedDrug.MaxDose ||
			newVaccination.Dose < selectedDrug.MinDose {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Selected drug dose invalid",
				"minDrugDose": selectedDrug.MinDose,
				"maxDrugDose": selectedDrug.MaxDose,
			})
			return
		}

		if !newVaccination.Date.UTC().After(selectedDrug.AvailableAt.UTC()) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":           "Drug unavailable on the selected date",
				"drugAvailableAt": selectedDrug.AvailableAt,
			})
			return
		}

		createResult := h.DB.Save(&newVaccination)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": createResult.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"vaccination": newVaccination,
		})
	}
}

func GetVaccinations(h *VaccinationHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vaccinations []models.Vaccination

		result := h.DB.Find(&vaccinations)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"drugs": vaccinations,
		})
	}
}
