package models

import (
	"time"

	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Name              string
	Brand             string
	Description       string
	Status            string
	NeedsMaintenance  bool
	SerialNumber      string
	MaintenancePeriod int              // in days
	EmployeeID        uint             // Foreign key for Employee
	Employee          Employee         // Relationship with Employee
	AssignedAt        *time.Time       // Optional: Track when the tool was assigned to the employee
	SupplierID        *uint            // Optional foreign key for Supplier
	Supplier          *Supplier        // Relationship with Supplier
	MaintenanceLogs   []MaintenanceLog `gorm:"polymorphic:Item;polymorphicValue:Tool"`
}
