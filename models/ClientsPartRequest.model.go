// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

func ClientsPartRequestWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("SupplierOfferPriceForClients", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SupplierPersonalDetails")
		})
	}).Preload("TakenOfferPrice", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SupplierPersonalDetails")
		})
	}).Preload("City").Preload("Area")
}

// ClientsPartRequest ..
type ClientsPartRequest struct {
	UserID                       uint                           `json:"user_id"`
	CarID                        uint                           `json:"carID"`
	CarMakeID                    uint                           `json:"carMakeID"`
	CarFuelID                    uint                           `json:"carFuelID"`
	CarMakeText                  string                         `json:"carMakeText"`
	CarModelText                 string                         `json:"carModelText"`
	CarMadeYear                  string                         `json:"carMadeYear"`
	CarFuelText                  string                         `json:"carFuelText"`
	CarPartText                  string                         `json:"carPartText"`
	CarPartQty                   string                         `json:"carPartQty"`
	CarPartDescription           string                         `json:"carPartDescription"`
	CarPartType                  int64                          `json:"carPartType"`
	Image                        string                         `json:"image"`
	CityID                       uint                           `json:"cityID"`
	AreaID                       uint                           `json:"areaID"`
	TakenOfferID                 uint                           `json:"takenOfferID" gorm:"default:0"`
	SupplierTakeOffer            bool                           `json:"supplierTakeOffer"`
	BookingType                  int64                          `json:"bookingType" gorm:"default:0"`
	BookingDate                  string                         `json:"bookingDate"`
	BookingTime                  string                         `json:"bookingTime"`
	Ending                       bool                           `json:"ending" gorm:"default:false"`
	User                         User                           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	City                         Cites                          `json:"city" gorm:"foreignKey:CityID;references:ID"`
	Area                         Areas                          `json:"area" gorm:"foreignKey:AreaID;references:ID"`
	SupplierOfferPriceForClients []SupplierOfferPriceForClients `json:"supplierOfferPriceForClients" gorm:"foreignKey:clientsPartRequestID;references:ID"`
	TakenOfferPrice              SupplierOfferPriceForClients   `json:"takenOfferPrice" gorm:"foreignKey:TakenOfferID;references:ID"`
	SupplierOfferPrice           SupplierOfferPriceForClients   `json:"supplierOfferPrice"`
	gorm.Model
}
