package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Description     string
	Contact         string
	Email           string `gorm:"uniqueIndex"`
	Phone           string
	Address         string
	Tools           []Tool           `gorm:"foreignKey:SupplierID"` // Link tools to supplier (optional relationship)
	MaintenanceLogs []MaintenanceLog `gorm:"polymorphic:Item;polymorphicValue:Supplier"`
}
