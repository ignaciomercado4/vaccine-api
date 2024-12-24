package handlers

import "gorm.io/gorm"

type AppHandler struct {
	DB *gorm.DB
}
