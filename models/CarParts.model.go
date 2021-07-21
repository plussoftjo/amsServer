// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// CarParts ..
type CarParts struct {
	Title string `json:"title"`
	gorm.Model
}
