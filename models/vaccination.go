package models

import (
	"time"

	"gorm.io/gorm"
)

type Vaccination struct {
	gorm.Model
	Name   string    `gorm:"not null" json:"name"`
	DrugId uint32    `gorm:"not null" json:"drugId"`
	Dose   uint32    `gorm:"not null" json:"dose"`
	Date   time.Time `gorm:"not null" json:"date"`
}
