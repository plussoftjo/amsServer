// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ClientsCarWithDetails ..
func ClientsCarWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("MainFactory").Preload("CarMake").Preload("User").Preload("CarModel").Preload("CarFuel")
}

// ClientsCar ..
type ClientsCar struct {
	UserID        uint          `json:"userID"`
	MainFactoryID uint          `json:"mainFactoryID"`
	CarMakeID     uint          `json:"carMakeID"`
	CarModelID    uint          `json:"carModelID"`
	CarMadeYear   int64         `json:"carMadeYear"`
	CarFuelID     uint          `json:"carFuelID"`
	MainFactory   MainFactories `json:"mainFactory" gorm:"foreignKey:MainFactoryID;references:ID"`
	CarMake       CarsMake      `json:"carMake" gorm:"foreignKey:CarMakeID;references:ID"`
	User          User          `json:"user" gorm:"foreignKey:User;references:ID"`
	CarModel      CarsModels    `json:"carModel" gorm:"foreignKey:CarModelID;references:ID"`
	CarFuel       CarFuels      `json:"carFuel" gorm:"foreignKey:CarFuelID;references:ID"`
	gorm.Model
}
