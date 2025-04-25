package models

import "gorm.io/gorm"

type Truck struct {
	gorm.Model
	LicensePlate     string `gorm:"uniqueIndex"`
	TruckModel       string
	Status           string
	NeedsMaintenance bool
	EmployeeID       uint
	Employee         Employee
	MaintenanceLogs  []MaintenanceLog `gorm:"polymorphic:Item;polymorphicValue:Truck"`
}
