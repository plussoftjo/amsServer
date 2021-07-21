package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierOfferPriceForClients ..
type SupplierOfferPriceForClients struct {
	UserID               uint   `json:"user_id"`
	ClientsPartRequestID uint   `json:"clientsPartRequestID"`
	Price                string `json:"price"`
	Note                 string `json:"note"`
	Qty                  string `json:"qty"`
	Seen                 bool   `json:"seen" gorm:"default:false"`
	User                 User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	gorm.Model
}
