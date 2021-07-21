// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierPersonalDetailsWithDetails ..
func SupplierPersonalDetailsWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("City").Preload("Area")
}

// SupplierPersonalDetails ..
type SupplierPersonalDetails struct {
	UserID       uint   `json:"user_id"`
	SupplierName string `json:"supplierName"`
	Description  string `json:"description"`
	CityID       uint   `json:"cityID"`
	AreaID       uint   `json:"areaID"`
	Image        string `json:"image"`
	Location     string `json:"location"`
	Facebook     string `json:"facebook"`
	City         Cites  `json:"city" gorm:"foreignKey:CityID;references:ID"`
	Area         Areas  `json:"area" gorm:"foreignKey:AreaID;references:ID"`
	gorm.Model
}
