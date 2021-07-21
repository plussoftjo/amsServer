// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CarsModels ..
type CarsModels struct {
	Title         string        `json:"title"`
	MainFactoryID uint          `json:"mainFactoryID"`
	CarsMakeID    uint          `json:"carsMakeID"`
	MainFactory   MainFactories `json:"mainFactory" gorm:"foreignKey:MainFactoryID;references:ID"`
	CarsMake      CarsMake      `json:"carsMake" gorm:"foreignKey:CarsMakeID;references:ID"`
	gorm.Model
}
