// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name                    string                  `json:"name"`
	Phone                   string                  `json:"phone" gorm:"unique"`
	PhoneCode               string                  `json:"phoneCode"`
	Password                string                  `json:"password"`
	RolesID                 uint                    `json:"roles_id"`
	UserType                uint                    `json:"user_type"` // 01 -> Supplier , 02 -> User, 03 -> TowTrack, 04 -> Controller
	ServicesID              uint                    `json:"servicesID" gorm:"default:0"`
	CountryID               uint                    `json:"countryID"`
	EndingRegister          bool                    `json:"endingRegister" gorm:"default:false"`
	Roles                   Roles                   `json:"roles" gorm:"foreignKey:RolesID;references:ID"`
	Country                 Countries               `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	Service                 Services                `json:"service" gorm:"foreignKey:ServicesID;references:ID"`
	SupplierPersonalDetails SupplierPersonalDetails `json:"supplierPersonalDetails" gorm:"foreignKey:UserID;references:ID" `
}

// Login ...
type Login struct {
	Phone    string `json:"phone" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}
