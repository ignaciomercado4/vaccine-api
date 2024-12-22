package models

import (
	"time"

	"gorm.io/gorm"
)

type Drug struct {
	gorm.Model
	name         string
	approved     bool
	min_dose     uint32
	max_dose     uint32
	available_at time.Time
}
