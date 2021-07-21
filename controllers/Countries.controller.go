// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Country -------------//

// StoreCountry ..
func StoreCountry(c *gin.Context) {
	// Register var and bind json
	var country models.Countries
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, country)
}

// IndexCountries ..
func IndexCountries(c *gin.Context) {
	// Register countries in var
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, countries)
}

// DestroyCountry ..
func DestroyCountry(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Countries{}, ID)
	// Return Countries list
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, countries)
}

// UpdateCountry ..
func UpdateCountry(c *gin.Context) {
	var country models.Countries
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&country).Update(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var countries []models.Countries
	config.DB.Find(&countries)
	c.JSON(200, gin.H{
		"country":   country,
		"countries": countries,
	})
}

// ------------- End Country ------------ //

// ------------- Cites -------------//

// StoreCity ..
func StoreCity(c *gin.Context) {
	// Register var and bind json
	var city models.Cites
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&city).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("Country").Find(&city)
	c.JSON(200, city)
}

// IndexCites ..
func IndexCites(c *gin.Context) {
	// Register countries in var
	var cites []models.Cites
	config.DB.Preload("Country").Find(&cites)
	c.JSON(200, cites)
}

// DestroyCity ..
func DestroyCity(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Cites{}, ID)
	// Return Countries list
	var cites []models.Cites
	config.DB.Preload("Country").Find(&cites)
	c.JSON(200, cites)
}

// UpdateCity ..
func UpdateCity(c *gin.Context) {
	var city models.Cites
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&city).Update(&city).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cites []models.Cites
	config.DB.Preload("Country").Find(&cites)
	c.JSON(200, gin.H{
		"city":  city,
		"cites": cites,
	})
}

// IndexCitesWithCountryID ..
func IndexCitesWithCountryID(c *gin.Context) {
	var cites []models.Cites
	ID := c.Param("id")

	config.DB.Where("country_id = ?", ID).Find(&cites)

	c.JSON(200, cites)
}

// ------------- End Cites ------------ //

// ------------- Areas -------------//

// StoreArea ..
func StoreArea(c *gin.Context) {
	// Register var and bind json
	var area models.Areas
	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&area).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Preload("Country").Preload("City").Find(&area)
	c.JSON(200, area)
}

// IndexAreas ..
func IndexAreas(c *gin.Context) {
	// Register countries in var
	var areas []models.Areas
	config.DB.Preload("Country").Preload("City").Find(&areas)
	c.JSON(200, areas)
}

// DestroyArea ..
func DestroyArea(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.Areas{}, ID)
	// Return Countries list
	var areas []models.Areas
	config.DB.Preload("Country").Preload("City").Find(&areas)
	c.JSON(200, areas)
}

// UpdateArea ..
func UpdateArea(c *gin.Context) {
	var area models.Areas
	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&area).Update(&area).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var areas []models.Areas
	config.DB.Preload("Country").Preload("City").Find(&areas)
	c.JSON(200, gin.H{
		"area":  area,
		"areas": areas,
	})
}

// IndexAreasWithCityID ..
func IndexAreasWithCityID(c *gin.Context) {
	var areas []models.Areas
	ID := c.Param("id")
	config.DB.Preload("Country").Preload("City").Where("city_id = ?", ID).Find(&areas)

	c.JSON(200, areas)

}

// ------------- End Areas ------------ //
