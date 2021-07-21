// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Services -------------//

// StoreService ..
func StoreService(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, service)
}

// IndexServices ..
func IndexServices(c *gin.Context) {
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, services)
}

// DestroyService ..
func DestroyService(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Services{}, ID)
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, services)
}

// UpdateService ..
func UpdateService(c *gin.Context) {
	var service models.Services
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&service).Update(&service).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var services []models.Services
	config.DB.Find(&services)
	c.JSON(200, gin.H{
		"service":  service,
		"services": services,
	})
}

// ------------- End Services ------------ //

// ------------- Sub Services -------------//

// StoreSubService ..
func StoreSubService(c *gin.Context) {
	var subService models.SubServices
	if err := c.ShouldBindJSON(&subService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&subService).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", subService.ID).Preload("Services").Find(&subService)
	c.JSON(200, subService)
}

// IndexSubServices ..
func IndexSubServices(c *gin.Context) {
	var subServices []models.SubServices
	config.DB.Preload("Services").Find(&subServices)
	c.JSON(200, subServices)
}

// DestroySubService ..
func DestroySubService(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.SubServices{}, ID)
	var subServices []models.SubServices
	config.DB.Preload("Services").Find(&subServices)
	c.JSON(200, subServices)
}

// UpdateSubService ..
func UpdateSubService(c *gin.Context) {
	var subService models.SubServices
	if err := c.ShouldBindJSON(&subService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&subService).Update(&subService).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var subServices []models.SubServices
	config.DB.Preload("Services").Find(&subServices)
	c.JSON(200, gin.H{
		"subService":  subService,
		"subServices": subServices,
	})
}

// IndexSubServicesWithServiceID ..
func IndexSubServicesWithServiceID(c *gin.Context) {
	ID := c.Param("id")
	var subServices []models.SubServices
	config.DB.Where("services_id = ?", ID).Find(&subServices)

	c.JSON(200, subServices)
}

// ------------- End Sub Services ------------ //

// ------------- Services Options -------------//

// StoreServicesOptions ..
func StoreServicesOptions(c *gin.Context) {
	var servicesOptions models.ServicesOptions
	if err := c.ShouldBindJSON(&servicesOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&servicesOptions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", servicesOptions.ID).Preload("Services").Preload("SubServices").Find(&servicesOptions)
	c.JSON(200, servicesOptions)
}

// IndexServicesOptions ..
func IndexServicesOptions(c *gin.Context) {
	var servicesOptions []models.ServicesOptions
	config.DB.Preload("Services").Preload("SubServices").Find(&servicesOptions)
	c.JSON(200, servicesOptions)
}

// DestroyServiceOptions ..
func DestroyServiceOptions(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.ServicesOptions{}, ID)
	var servicesOptions []models.ServicesOptions
	config.DB.Preload("Services").Preload("SubServices").Find(&servicesOptions)
	c.JSON(200, servicesOptions)
}

// UpdateServicesOptions ..
func UpdateServicesOptions(c *gin.Context) {
	var servicesOption models.ServicesOptions
	if err := c.ShouldBindJSON(&servicesOption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&servicesOption).Update(&servicesOption).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var servicesOptions []models.ServicesOptions
	config.DB.Preload("Services").Preload("SubServices").Find(&servicesOptions)
	c.JSON(200, gin.H{
		"servicesOption":  servicesOption,
		"servicesOptions": servicesOptions,
	})
}

// ------------- End Sub Services ------------ //
