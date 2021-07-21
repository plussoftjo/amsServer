// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// IndexAutoMobileServiceServices ..
func IndexAutoMobileServiceServices(c *gin.Context) {
	var services []models.SubServices

	config.DB.Where("section = ?", 3).Preload("Services").Find(&services)

	c.JSON(200, services)
}

// IndexServiceOptionsWithSubServiceID ..
func IndexServiceOptionsWithSubServiceID(c *gin.Context) {
	ID := c.Param("id")
	var serviceOptions []models.ServicesOptions

	config.DB.Where("sub_services_id = ?", ID).Find(&serviceOptions)

	c.JSON(200, serviceOptions)
}

// StoreAutoMobileService ..
func StoreAutoMobileService(c *gin.Context) {
	var autoMobileService models.AutoMobileRequest

	if err := c.ShouldBindJSON(&autoMobileService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&autoMobileService).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", autoMobileService.ID).Scopes(models.AutoMobileRequestWithDetails).First(&autoMobileService)

	c.JSON(200, autoMobileService)
}

// IndexAutoMobileServiceForSupplier ..
func IndexAutoMobileServiceForSupplier(c *gin.Context) {
	ID := c.Param("id")

	var user models.User
	config.DB.Where("id = ?", ID).First(&user)

	var subServices []models.SubServices
	config.DB.Where("services_id = ?", user.ServicesID).Where("section = ?", 3).Find(&subServices)

	var autoMobileRequest []models.AutoMobileRequest
	for _, subService := range subServices {
		var amrList []models.AutoMobileRequest
		config.DB.Where("sub_service_id = ?", subService.ID).Where("end = ?", 0).Order("id desc").Limit(10).Scopes(models.AutoMobileRequestWithDetails).Find(&amrList)
		autoMobileRequest = append(autoMobileRequest, amrList...)
	}

	var fullListForSupplier []models.AutoMobileRequest
	for _, amr := range autoMobileRequest {
		var supplierAmsOffer models.SupplierAmrOffer
		err := config.DB.Where("user_id = ?", ID).Where("auto_mobile_request_id = ?", amr.ID).First(&supplierAmsOffer).Error
		if err != nil {
			fullListForSupplier = append(fullListForSupplier, amr)
		} else {
			amr.SupplierMakeOffer = true
			amr.SupplierOffer = supplierAmsOffer
			fullListForSupplier = append(fullListForSupplier, amr)
		}
	}

	c.JSON(200, gin.H{
		"autoMobileRequest": fullListForSupplier,
	})
}

// IndexAutoMobileServiceForUser ..
func IndexAutoMobileServiceForUser(c *gin.Context) {
	ID := c.Param("id")

	var autoMobileRequest []models.AutoMobileRequest
	config.DB.Where("user_id = ?", ID).Order("id desc").Where("end = ?", 0).Scopes(models.AutoMobileRequestWithDetails).Find(&autoMobileRequest)

	c.JSON(200, autoMobileRequest)
}

// SupplierTakeAMS ..
func SupplierTakeAMS(c *gin.Context) {
	type SupplierTakeAmsType struct {
		AmsID      uint `json:"amsID"`
		SupplierID uint `json:"supplierID"`
		OfferID    uint `json:"offerID"`
	}

	var data SupplierTakeAmsType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ams models.AutoMobileRequest
	config.DB.Where("id = ?", data.AmsID).Where("end = ?", 0).Scopes(models.AutoMobileRequestWithDetails).First(&ams)

	ams.SupplierUserID = data.SupplierID
	ams.TakenOfferID = data.OfferID

	config.DB.Save(&ams)

	c.JSON(200, ams)

}

// StoreSupplierAmrOffer ..
func StoreSupplierAmrOffer(c *gin.Context) {
	var data models.SupplierAmrOffer
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&data)

	c.JSON(200, data)

}

// Ending AMS ..
func EndingAms(c *gin.Context) {
	type EndingAmsType struct {
		AutoMobileRequestID uint            `json:"autoMobileRequestID"`
		UserRate            models.UserRate `json:"userRate"`
	}

	var data EndingAmsType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ams models.AutoMobileRequest
	config.DB.Where("id = ?", data.AutoMobileRequestID).Scopes(models.AutoMobileRequestWithDetails).First(&ams)

	ams.End = 1

	config.DB.Save(&ams)

	config.DB.Where("id = ?", data.AutoMobileRequestID).Scopes(models.AutoMobileRequestWithDetails).First(&ams)

	userRate := data.UserRate

	if err := config.DB.Create(&userRate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"userRate": userRate,
		"ams":      ams,
	})
}

// Ending AMS ..
func EndAmsFromClient(c *gin.Context) {
	amsID := c.Param("id")
	var ams models.AutoMobileRequest
	config.DB.Where("id = ?", amsID).Scopes(models.AutoMobileRequestWithDetails).First(&ams)

	ams.End = 1

	config.DB.Save(&ams)
	c.JSON(200, gin.H{
		"ams": ams,
	})
}
