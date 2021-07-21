// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CarFuels ..
type CarFuels struct {
	Title string `json:"title"`
	gorm.Model
}
