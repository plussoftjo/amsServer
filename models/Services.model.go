// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Services ..
type Services struct {
	Title string `json:"title"`
	// Sections string `json:"sections"`
	Sections     pq.Int64Array `json:"sections" gorm:"type:integer[]"`
	Image        string        `json:"image"`
	RegisterType int           `json:"registerType"`
	gorm.Model
}
