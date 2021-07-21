// Package config ...
package config

import (
	"server/models"

	"github.com/jinzhu/gorm"
	// Connect mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// models
)

// SetupDB ...

// DB ..
var DB *gorm.DB

// SetupDB ..
func SetupDB() {
	database, err := gorm.Open("mysql", "root:00962s00962S!@tcp(127.0.0.1:3306)/car?charset=utf8mb4&parseTime=True&loc=Local")

	// If Error in Connect
	if err != nil {
		panic(err)
	}
	// User Models Setup
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.AuthClients{})
	database.AutoMigrate(&models.AuthTokens{})
	database.AutoMigrate(&models.Roles{})

	// Services
	database.AutoMigrate(&models.Services{})
	database.AutoMigrate(&models.SubServices{})
	database.AutoMigrate(&models.ServicesOptions{})
	// Countries And things
	database.AutoMigrate(&models.Countries{})
	database.AutoMigrate(&models.Cites{})
	database.AutoMigrate(&models.Areas{})

	// Cars
	database.AutoMigrate(&models.MainFactories{})
	database.AutoMigrate(&models.CarsMake{})
	database.AutoMigrate(&models.CarsModels{})
	database.AutoMigrate(&models.CarParts{})
	database.AutoMigrate(&models.CarFuels{})

	// Suppliers
	database.AutoMigrate(&models.SupplierSpecialty{})
	database.AutoMigrate(&models.SupplierPersonalDetails{})

	// SupplierRequest
	database.AutoMigrate(&models.SupplierRequestPart{})
	database.AutoMigrate(&models.SupplierOfferPrice{})

	// Clients
	database.AutoMigrate(&models.ClientsCar{})
	database.AutoMigrate(&models.ClientsPartRequest{})
	database.AutoMigrate(&models.SupplierOfferPriceForClients{})

	// AutoMobileRequest
	database.AutoMigrate(&models.AutoMobileRequest{})
	database.AutoMigrate(&models.SupplierAmrOffer{})

	// User Rateing
	database.AutoMigrate(&models.UserRate{})

	DB = database
}
