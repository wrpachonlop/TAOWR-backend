package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	IsAdmin  bool   `gorm:"default:false"`
	IsActive bool   `gorm:"default:true"`
	Tools    []Tool
	Trucks   []Truck
}
