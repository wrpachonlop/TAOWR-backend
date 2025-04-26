package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FullName       string
	Email          string `gorm:"uniqueIndex;not null"`
	Phone          string
	Address        string
	PostalCode     string
	DriversLicense string `gorm:"uniqueIndex;not null"`
	SIN            string `gorm:"uniqueIndex;not null"`
	BirthDate      time.Time

	IsAdmin  bool `gorm:"default:false"`
	IsActive bool `gorm:"default:true"`

	StartDate    time.Time
	JobTitle     string
	TypeContract string
	Salary       float64

	InstitutionNo   string
	AccountNo       string
	TransitNo       string
	BankAccountName string

	Tools  []Tool
	Trucks []Truck

	EmergencyContacts []EmergencyContact `gorm:"foreignKey:EmployeeID"`
}
