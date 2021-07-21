package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierOfferPrice ..
type SupplierOfferPrice struct {
	UserID                uint   `json:"user_id"`
	SupplierRequestPartID uint   `json:"supplierRequestPartID"`
	Price                 string `json:"price"`
	Note                  string `json:"note"`
	Qty                   string `json:"qty"`
	Seen                  bool   `json:"seen" gorm:"default:false"`
	User                  User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	gorm.Model
}
