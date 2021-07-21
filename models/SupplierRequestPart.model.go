// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// SupplierRequestPartWithDetails ..
func SupplierRequestPartWithDetails(db *gorm.DB) *gorm.DB {
	return db.Preload("MainFactory").Preload("CarMake").Preload("User").Preload("TakenOffer", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SupplierPersonalDetails")
		})
	}).Preload("ForOwnerSupplierOfferPrice", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("SupplierPersonalDetails")
		})
	}).Preload("CarModels").Preload("City").Preload("Area")
}

// SupplierRequestPart ..
type SupplierRequestPart struct {
	UserID                     uint                 `json:"user_id"`
	MainFactoryID              uint                 `json:"mainFactoryID"`
	CarMakeID                  uint                 `json:"carMakeID"`
	CarModelID                 uint                 `json:"carModelID"`
	CarFuelID                  uint                 `json:"carFuelID"`
	MadeYear                   uint                 `json:"madeYear"`
	CarPartText                string               `json:"carPartText"`
	CarPartQty                 string               `json:"carPartQty"`
	CarPartDescription         string               `json:"carPartDescription"`
	CarPartType                uint                 `json:"carPartType"` // 1 => New And Used , 2 => New, 3 => Used
	Image                      string               `json:"image"`
	OfferPrice                 bool                 `json:"offerPrice"`
	CityID                     uint                 `json:"cityID"`
	AreaID                     uint                 `json:"areaID"`
	UserMakeOffer              bool                 `json:"userMakeOffer"`
	TakenOfferID               uint                 `json:"takenOfferID" gorm:"default:0"`
	Ending                     bool                 `json:"ending" gorm:"default:false"`
	MainFactory                MainFactories        `json:"mainFactory" gorm:"foreignKey:MainFactoryID;references:ID"`
	CarMake                    CarsMake             `json:"carMake" gorm:"foreignKey:CarMakeID;references:ID"`
	CarModels                  CarsModels           `json:"carModel" gorm:"foreignKey:CarModelID;references:ID"`
	City                       Cites                `json:"city" gorm:"foreignKey:CityID;references:ID"`
	Area                       Areas                `json:"area" gorm:"foreignKey:AreaID;references:ID"`
	User                       User                 `json:"user" gorm:"foreignKey:UserID;references:ID"`
	SupplierOfferPrice         SupplierOfferPrice   `json:"supplierOfferPrice"`
	TakenOffer                 SupplierOfferPrice   `json:"takenOffer" gorm:"foreignKey:TakenOfferID;references:ID"`
	CarFuel                    CarFuels             `json:"carFuel" gorm:"foreignKey:CarFuelID;references:ID"`
	ForOwnerSupplierOfferPrice []SupplierOfferPrice `json:"forOwnerSupplierOfferPrice"`
	gorm.Model
}
