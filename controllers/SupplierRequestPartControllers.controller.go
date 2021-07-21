// Package controllers ...
package controllers

import (
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	// "strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// IndexMainFactoryForSupplierRequest ..
func IndexMainFactoryForSupplierRequest(c *gin.Context) {
	var mainFactories []models.MainFactories
	config.DB.Find(&mainFactories)

	var carsMake []models.CarsMake
	var carsModels []models.CarsModels
	if len(mainFactories) >= 1 {
		firstMainFactory := mainFactories[0]

		config.DB.Where("main_factory_id = ?", firstMainFactory.ID).Find(&carsMake)

		if len(carsMake) >= 1 {
			firstCarsMake := carsMake[0]
			config.DB.Where("cars_make_id = ?", firstCarsMake.ID).Find(&carsModels)
		}
	}

	// CarPartsAlsow
	var carParts []models.CarParts
	config.DB.Find(&carParts)

	// Fuels
	var fuels []models.CarFuels
	config.DB.Find(&fuels)

	c.JSON(200, gin.H{
		"mainFactories": mainFactories,
		"carsMake":      carsMake,
		"carsModels":    carsModels,
		"carParts":      carParts,
		"fuels":         fuels,
	})
}

// StoreSupplierRequestPart ..
func StoreSupplierRequestPart(c *gin.Context) {
	var data models.SupplierRequestPart

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var thisRequest models.SupplierRequestPart
	config.DB.Where("id = ?", data.ID).Scopes(models.SupplierRequestPartWithDetails).First(&thisRequest)

	c.JSON(200, thisRequest)
}

// IndexSupplierRequestPartForSuppliers ..
func IndexSupplierRequestPartForSuppliers(c *gin.Context) {
	// REGISTER USERID GET IT FROM THE URL
	UserID := c.Param("id")

	// GET SUPPLIER SPECIALTY WITH USERID
	var supplierSpecialty models.SupplierSpecialty
	config.DB.Where("user_id = ?", UserID).First(&supplierSpecialty)

	// GET PERSONAL DETAILS FOR THE SUPPLIER WITH USERID
	var supplierPersonalDetails models.SupplierPersonalDetails
	config.DB.Where("user_id = ?", UserID).First(&supplierPersonalDetails)

	// THREE DETAILS IS VERY IMPORTANT => :CARMAKE :CITYID :AREAID
	carMakesIDs := supplierSpecialty.CarMakeIDs

	// RETURN ERROR IF NOT HAVE CARS MAKE
	if len(carMakesIDs) == 0 {
		c.JSON(500, gin.H{
			"message": "NOT HAVE CARSMAKE",
			"code":    101,
		})
		return
	}

	// NOW WE MOST CONVERT THE CARMAKES IDS TO THE STRING
	// REGISTER THE LIST WITH STRING
	var carsMakeIDList []string
	// NOW APPEND THE INT WITH CONVERT IT TO THE STRING TO LIST
	for _, carMakeID := range carMakesIDs {
		// carsMakeIDList = append(carsMakeIDList, strconv.FormatInt(carMakeID, 10))
		fmt.Println(carMakeID)
	}

	// NOW MAKE IT STRING
	carMakeIDsString := strings.Join(carsMakeIDList, ",")
	// AND REGISTER THE CITY AND AREA ID
	cityID := supplierPersonalDetails.CityID
	areaID := supplierPersonalDetails.AreaID

	// CAR FUELS CHAECK

	var carFuelsIDLIST []string
	for _, fuelID := range supplierSpecialty.FuelsIDs {
		// carFuelsIDLIST = append(carFuelsIDLIST, strconv.FormatInt(fuelID, 10))
		fmt.Println(fuelID)
	}
	carFuelIDsString := strings.Join(carFuelsIDLIST, ",")

	// REGISTER SUPPLIER REQUEST
	var supplierRequestPart []models.SupplierRequestPart
	// NOW MAKE QUERY
	err := config.DB.
		Where("city_id = ? OR city_id = 0", cityID).
		Where("area_id = ? OR area_id = 0", areaID).
		Where("car_make_id IN "+"("+carMakeIDsString+")").
		Where("car_fuel_id IN "+"("+carFuelIDsString+")").
		Where("ending = ?", false).
		Not("user_id = ?", UserID).
		Scopes(models.SupplierRequestPartWithDetails).
		Order("id desc").
		Limit(15).
		Find(&supplierRequestPart).Error
	if err != nil {
		c.JSON(500, err)
		return
	}

	// Register The full supplier request part with user
	var fullSupplierRequestPart []models.SupplierRequestPart

	for _, srp := range supplierRequestPart {

		// Check if have offer price
		var supplierOfferPrice models.SupplierOfferPrice
		err := config.DB.Where("user_id = ?", UserID).Where("supplier_request_part_id = ?", srp.ID).Find(&supplierOfferPrice).Error
		if err != nil {
			srp.UserMakeOffer = false
		} else {
			srp.UserMakeOffer = true
			srp.SupplierOfferPrice = supplierOfferPrice

			supplierOfferPrice.Seen = true
			config.DB.Save(&supplierOfferPrice)
		}
		fullSupplierRequestPart = append(fullSupplierRequestPart, srp)

	}

	c.JSON(200, gin.H{
		"supplierRequestPartList": fullSupplierRequestPart,
	})
}

// StoreSupplierOfferPrice ..
func StoreSupplierOfferPrice(c *gin.Context) {
	var supplierOfferPrice models.SupplierOfferPrice
	if err := c.ShouldBindJSON(&supplierOfferPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&supplierOfferPrice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, supplierOfferPrice)

}

// UpdateTakenOfferPriceForRequestPartType ..
type UpdateTakenOfferPriceForRequestPartType struct {
	SupplierRequestPartID uint `json:"supplierRequestPartID"`
	OfferPriceID          uint `json:"offerPriceID"`
}

// UpdateTakenOfferPriceForRequestPart ..
func UpdateTakenOfferPriceForRequestPart(c *gin.Context) {
	var data UpdateTakenOfferPriceForRequestPartType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var supplierRequestPart models.SupplierRequestPart
	config.DB.Where("id = ?", data.SupplierRequestPartID).Find(&supplierRequestPart)
	supplierRequestPart.TakenOfferID = data.OfferPriceID
	config.DB.Save(&supplierRequestPart)

	var thisSupplierRequestPart models.SupplierRequestPart
	config.DB.Where("id = ?", data.SupplierRequestPartID).Scopes(models.SupplierRequestPartWithDetails).First(&thisSupplierRequestPart)

	c.JSON(200, gin.H{
		"message":                 "success",
		"code":                    200,
		"thisSupplierRequestPart": thisSupplierRequestPart,
	})

}

// IndexSupplierMyRequestPart ..
func IndexSupplierMyRequestPart(c *gin.Context) {
	UserID := c.Param("id")
	var SupplierRequestPart []models.SupplierRequestPart

	config.DB.Where("user_id = ?", UserID).Where("ending = ?", false).Scopes(models.SupplierRequestPartWithDetails).Order("id desc").Find(&SupplierRequestPart)

	c.JSON(200, SupplierRequestPart)

}

// TestIndexSupplierRequestPart ..
func TestIndexSupplierRequestPart(c *gin.Context) {
	var SupplierRequestPart []models.SupplierRequestPart
	err := config.DB.Where("user_id = ?", 66).Find(&SupplierRequestPart).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, SupplierRequestPart)
}

// EndSupplierRequestPart ..
func EndSupplierRequestPart(c *gin.Context) {
	ID := c.Param("id")

	var supplierRequestPart models.SupplierRequestPart
	config.DB.Where("id = ?", ID).Scopes(models.SupplierRequestPartWithDetails).First(&supplierRequestPart)

	supplierRequestPart.Ending = true

	config.DB.Save(&supplierRequestPart)

	c.JSON(200, supplierRequestPart)

}
