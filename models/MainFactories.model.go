// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// MainFactories ..
type MainFactories struct {
	Title string `json:"title"`
	gorm.Model
}
