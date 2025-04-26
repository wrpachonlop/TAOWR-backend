package models

import "gorm.io/gorm"

type EmergencyContact struct {
	gorm.Model
	EmployeeID   uint
	Name         string
	Phone        string
	Address      string
	Relationship string
}
