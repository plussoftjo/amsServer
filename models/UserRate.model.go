// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// UserRate ..
type UserRate struct {
	UserID uint   `json:"userID"`
	Rate   int64  `json:"rate"`
	Note   string `json:"note"`
	gorm.Model
}
