// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SubServices ..
type SubServices struct {
	Title      string   `json:"title"`
	Section    int      `json:"section"`
	ServicesID uint     `json:"servicesID"`
	Multiple   bool     `json:"multiple"`
	Services   Services `json:"services"`
	gorm.Model
}
