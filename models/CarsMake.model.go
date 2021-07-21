// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CarsMake ..
type CarsMake struct {
	Title         string        `json:"title"`
	MainFactoryID uint          `json:"mainFactoryID"`
	MainFactory   MainFactories `json:"mainFactory" gorm:"foreignKey:MainFactoryID;references:ID"`
	gorm.Model
}
