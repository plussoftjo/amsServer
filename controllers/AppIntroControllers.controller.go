// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexAppIntroCountry ..
func IndexAppIntroCountry(c *gin.Context) {
	var countries []models.Countries

	config.DB.Find(&countries)

	c.JSON(200, gin.H{
		"countries": countries,
	})
}

// IndexAppIntroServices ..
func IndexAppIntroServices(c *gin.Context) {
	var services []models.Services

	config.DB.Find(&services)

	c.JSON(200, gin.H{
		"services": services,
	})
}

// IndexAppIntroSpecialtyData ..
func IndexAppIntroSpecialtyData(c *gin.Context) {
	var mainFactories []models.MainFactories
	config.DB.Find(&mainFactories)

	var fuels []models.CarFuels
	config.DB.Find(&fuels)

	c.JSON(200, gin.H{
		"mainFactories": mainFactories,
		"fuels":         fuels,
	})

}

// IndexAppIntroCarsWithMainFactoryID ..
func IndexAppIntroCarsWithMainFactoryID(c *gin.Context) {
	ID := c.Param("id")

	var cars []models.CarsMake
	config.DB.Where("main_factory_id = ?", ID).Find(&cars)

	c.JSON(200, gin.H{
		"cars": cars,
	})
}

// StoreSupplierSpecialty ..
func StoreSupplierSpecialty(c *gin.Context) {
	var supplierSpecialty models.SupplierSpecialty

	if err := c.ShouldBindJSON(&supplierSpecialty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&supplierSpecialty).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"supplierSpecialty": supplierSpecialty,
	})
}

// UpdateSupplierSpecialty ..
func UpdateSupplierSpecialty(c *gin.Context) {
	var data models.SupplierSpecialty

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Save(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"supplierSpecialty": data,
	})
}

// StoreSupplierPersonalDetails ..
func StoreSupplierPersonalDetails(c *gin.Context) {
	var supplierPersonalDetails models.SupplierPersonalDetails
	if err := c.ShouldBindJSON(&supplierPersonalDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&supplierPersonalDetails).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.Where("id = ?", supplierPersonalDetails.UserID).First(&user)

	user.EndingRegister = true

	config.DB.Save(&user)

	c.JSON(200, gin.H{
		"supplierPersonalDetails": supplierPersonalDetails,
	})
}

// UpdateSupplierPersonalDetails ..
func UpdateSupplierPersonalDetails(c *gin.Context) {
	var supplierPersonalDetails models.SupplierPersonalDetails
	var data models.SupplierPersonalDetails

	if err := c.ShouldBindJSON(&supplierPersonalDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", supplierPersonalDetails.ID).Find(&data)

	data.AreaID = supplierPersonalDetails.AreaID
	data.CityID = supplierPersonalDetails.CityID
	data.Image = supplierPersonalDetails.Image
	data.Description = supplierPersonalDetails.Description
	data.Location = supplierPersonalDetails.Location
	data.SupplierName = supplierPersonalDetails.SupplierName
	data.Facebook = supplierPersonalDetails.Facebook

	if err := config.DB.Save(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"supplierPersonalDetails": data,
	})
}

// StoreUpdateUserWithServicesIDType ..
type StoreUpdateUserWithServicesIDType struct {
	UserID     uint `json:"userID"`
	ServicesID uint `json:"servicesID"`
	UserType   uint `json:"userType"`
}

// StoreUpdateUserWithServicesID ..
func StoreUpdateUserWithServicesID(c *gin.Context) {
	var data StoreUpdateUserWithServicesIDType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch User
	var user models.User
	config.DB.Where("id = ?", data.UserID).First(&user)

	user.UserType = data.UserType
	user.ServicesID = data.ServicesID

	config.DB.Save(&user)

	c.JSON(200, user)

}

// StoreClientsCar ..
func StoreClientsCar(c *gin.Context) {
	var data models.ClientsCar
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.Where("id = ?", data.UserID).First(&user)

	user.UserType = 2

	config.DB.Save(&user)

	var clientCar models.ClientsCar
	config.DB.Where("id = ?", data.ID).Scopes(models.ClientsCarWithDetails).First(&clientCar)

	c.JSON(200, gin.H{
		"clientCar": clientCar,
	})
}
