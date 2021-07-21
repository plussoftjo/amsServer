package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierAmrOffer ..
type SupplierAmrOffer struct {
	UserID              uint   `json:"user_id"`
	AutoMobileRequestID uint   `json:"autoMobileRequestID"`
	Price               string `json:"price"`
	Note                string `json:"note"`
	Seen                bool   `json:"seen" gorm:"default:false"`
	User                User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	gorm.Model
}
