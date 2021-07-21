// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// AutoMobileRequestWithDetails ..
func AutoMobileRequestWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("User").Preload("SubService").Preload("Supplier", func(db *gorm.DB) *gorm.DB {
		return db.Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(SupplierPersonalDetailsWithDetails)
		})
	}).Preload("ClientCar", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(ClientsCarWithDetails)
	}).Preload("OffersList", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
				return db.Scopes(SupplierPersonalDetailsWithDetails)
			})
		})
	}).Preload("TakenOffer")
}

// AutoMobileRequest ..
type AutoMobileRequest struct {
	UserID            uint               `json:"userID"`
	CarID             uint               `json:"carID"`
	SubServiceID      uint               `json:"subServiceID"`
	ServicesOptions   string             `json:"serviceOptions"`
	Note              string             `json:"note"`
	Location          string             `json:"location"`
	SupplierUserID    uint               `json:"supplierUserID" gorm:"default:0"`
	TakenOfferID      uint               `json:"takenOfferID" gorm:"default:0"`
	End               int64              `json:"end" gorm:"default:0"`
	BookingType       int64              `json:"bookingType" gorm:"default:0"`
	BookingDate       string             `json:"bookingDate"`
	BookingTime       string             `json:"bookingTime"`
	SupplierMakeOffer bool               `json:"supplierMakeOffer" gorm:"default:0"`
	SupplierOffer     SupplierAmrOffer   `json:"supplierOffer"`
	TakenOffer        SupplierAmrOffer   `json:"takenOffer" gorm:"foreignKey:takenOfferID;references:ID"`
	User              User               `json:"user" gorm:"foreignKey:UserID;references:ID"`
	ClientCar         ClientsCar         `json:"clientCar" gorm:"foreignKey:CarID;references:ID"`
	SubService        SubServices        `json:"subService" gorm:"foreignKey:SubServiceID;references:ID"`
	Supplier          User               `json:"supplier" gorm:"foreignKey:SupplierUserID;references:ID"`
	OffersList        []SupplierAmrOffer `json:"offerList" gorm:"foreignKey:AutoMobileRequestID;references:ID"`
	gorm.Model
}
