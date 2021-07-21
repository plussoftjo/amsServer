// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Areas ..
type Areas struct {
	Title     string    `json:"title"`
	CountryID uint      `json:"countryID"`
	CityID    uint      `json:"cityID"`
	Country   Countries `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	City      Cites     `json:"city" gorm:"foreignKey:CityID;references:ID"`

	gorm.Model
}
