// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// Countries ..
type Countries struct {
	Title     string `json:"title"`
	Code      string `json:"code"`
	Currency  string `json:"currency"`
	Language  string `json:"language"`
	PhoneCode string `json:"phoneCode"`
	gorm.Model
}
