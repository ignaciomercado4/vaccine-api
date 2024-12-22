package models

import (
	"time"

	"gorm.io/gorm"
)

type Drug struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name"`
	Approved    bool      `gorm:"not null" json:"approved"`
	MinDose     uint32    `gorm:"not null" json:"minDose"`
	MaxDose     uint32    `gorm:"not null" json:"maxDose"`
	AvailableAt time.Time `gorm:"not null" json:"availableAt"`
}
