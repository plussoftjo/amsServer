// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Cites ..
type Cites struct {
	Title     string    `json:"title"`
	CountryID uint      `json:"countryID"`
	Country   Countries `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	gorm.Model
}
