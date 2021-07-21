// Package controllers ...
package controllers

import (
	"net/http"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IndexServicesForWorkshoplist ..
func IndexServicesForWorkshoplist(c *gin.Context) {
	var subServices []models.SubServices
	config.DB.Where("section = ?", 2).Find(&subServices)

	var users []models.User
	config.DB.Where("user_type = 1").Order("RAND()").Limit(10).Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
		return db.Preload("City").Preload("Area")
	}).Find(&users)

	c.JSON(200, gin.H{
		"subServices": subServices,
		"users":       users,
	})
}

// IndexWorkshopList ..
func IndexWorkshopList(c *gin.Context) {
	type WorkShopListType struct {
		CityID        uint `json:"cityID"`
		AreaID        uint `json:"areaID"`
		SubServicesID uint `json:"subServicesID"`
	}
	var data WorkShopListType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var subService models.SubServices
	config.DB.Where("id = ?", data.SubServicesID).First(&subService)

	var users []models.User
	config.DB.Where("services_id = ?", subService.ServicesID).Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
		return db.Preload("City").Preload("Area")
	}).Find(&users)

	if data.CityID == 0 && data.AreaID == 0 {
		c.JSON(200, gin.H{
			"users": users,
		})
		return
	}

	if data.CityID != 0 && data.AreaID == 0 {
		var fullUsersList []models.User
		for _, user := range users {
			var supplierPersonalDetails models.SupplierPersonalDetails
			config.DB.Where("user_id = ?", user.ID).Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
				return db.Preload("City").Preload("Area")
			}).First(&supplierPersonalDetails)
			if supplierPersonalDetails.CityID == data.CityID {
				fullUsersList = append(fullUsersList, user)
			}
		}
		c.JSON(200, gin.H{
			"users": fullUsersList,
		})
		return

	}
	if data.CityID != 0 && data.AreaID != 0 {
		var fullUsersList []models.User
		for _, user := range users {
			var supplierPersonalDetails models.SupplierPersonalDetails
			config.DB.Where("user_id = ?", user.ID).Preload("SupplierPersonalDetails", func(db *gorm.DB) *gorm.DB {
				return db.Preload("City").Preload("Area")
			}).First(&supplierPersonalDetails)
			if supplierPersonalDetails.CityID == data.CityID && supplierPersonalDetails.AreaID == data.AreaID {
				fullUsersList = append(fullUsersList, user)
			}
		}
		c.JSON(200, gin.H{
			"users": fullUsersList,
		})
		return
	}

}
