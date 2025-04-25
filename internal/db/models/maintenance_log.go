package models

import (
	"time"

	"gorm.io/gorm"
)

type MaintenanceLog struct {
	gorm.Model
	ItemID        uint
	ItemType      string // "Tool" or "Truck"
	Description   string
	Date          time.Time
	PerformedByID uint
	PerformedBy   Employee
}
