// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	// "github.com/lib/pq"
)

// SupplierSpecialty ..
type SupplierSpecialty struct {
	UserID        uint `json:"user_id"`
	MainFactoryID uint `json:"mainFactoryID"`
	// CarMakeIDs    string `json:"carMakeIDs"`
	CarMakeIDs pq.Int64Array `json:"carMakeIDs" gorm:"type:integer[]"`
	// FuelsIDs string `json:"fuelsIDs"`
	FuelsIDs      pq.Int64Array `json:"fuelsIDs" gorm:"type:integer[]"`
	SpecialtyID   uint          `json:"specialtyID"`
	SpecialtyType uint          `json:"specialtyType"`
	gorm.Model
}
