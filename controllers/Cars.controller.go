// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------- Main Factory ----------- //

// StoreMainFactory..
func StoreMainFactory(c *gin.Context) {
	// Register var and bind json
	var mainFactory models.MainFactories
	if err := c.ShouldBindJSON(&mainFactory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&mainFactory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, mainFactory)
}

// IndexMainFactory ..
func IndexMainFactory(c *gin.Context) {
	// Register countries in var
	var mainFactories []models.MainFactories
	config.DB.Find(&mainFactories)
	c.JSON(200, mainFactories)
}

// DestroyMainFactory ..
func DestroyMainFactory(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.MainFactories{}, ID)
	// Return Countries list
	var mainFactories []models.MainFactories
	config.DB.Find(&mainFactories)
	c.JSON(200, mainFactories)
}

// UpdateMainFactory ..
func UpdateMainFactory(c *gin.Context) {
	var mainFactory models.MainFactories
	if err := c.ShouldBindJSON(&mainFactory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&mainFactory).Update(&mainFactory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var mainFactories []models.MainFactories
	config.DB.Find(&mainFactories)
	c.JSON(200, gin.H{
		"mainFactory":   mainFactory,
		"mainFactories": mainFactories,
	})
}

// ------------- End Main Factory ----------- //

// ------------- Cars Make ----------- //

// StoreCarMake..
func StoreCarMake(c *gin.Context) {
	// Register var and bind json
	var carMake models.CarsMake
	if err := c.ShouldBindJSON(&carMake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&carMake).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("MainFactory").Find(&carMake)
	c.JSON(200, carMake)
}

// IndexCarsMake ..
func IndexCarsMake(c *gin.Context) {
	var carsMake []models.CarsMake
	config.DB.Preload("MainFactory").Find(&carsMake)
	c.JSON(200, carsMake)
}

// DestroyCarMake ..
func DestroyCarMake(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.CarsMake{}, ID)
	// Return Countries list
	var carsMake []models.CarsMake
	config.DB.Preload("MainFactory").Find(&carsMake)
	c.JSON(200, carsMake)
}

// UpdateCarsMake ..
func UpdateCarsMake(c *gin.Context) {
	var carMake models.CarsMake
	if err := c.ShouldBindJSON(&carMake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&carMake).Update(&carMake).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var carsMake []models.CarsMake
	config.DB.Preload("MainFactory").Find(&carsMake)
	c.JSON(200, gin.H{
		"carMake":  carMake,
		"carsMake": carsMake,
	})
}

// IndexCarsMakeWithMainFactory ..
func IndexCarsMakeWithMainFactory(c *gin.Context) {
	var carsMake []models.CarsMake
	ID := c.Param("id")

	config.DB.Where("main_factory_id = ?", ID).Find(&carsMake)

	c.JSON(200, carsMake)
}

// ------------- End Cars Make ----------- //

// ------------- Cars Models ----------- //

// StoreCarModel..
func StoreCarModel(c *gin.Context) {
	// Register var and bind json
	var carModel models.CarsModels
	if err := c.ShouldBindJSON(&carModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&carModel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Preload("MainFactory").Preload("CarsMake").Find(&carModel)
	c.JSON(200, carModel)
}

// IndexCarsModels ..
func IndexCarsModels(c *gin.Context) {
	var carsModels []models.CarsModels
	config.DB.Preload("MainFactory").Preload("CarsMake").Find(&carsModels)
	c.JSON(200, carsModels)
}

// DestroyCarModel ..
func DestroyCarModel(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.CarsModels{}, ID)
	// Return Countries list
	var carsModels []models.CarsModels
	config.DB.Preload("MainFactory").Preload("CarsMake").Find(&carsModels)
	c.JSON(200, carsModels)
}

// UpdateCarsModels ..
func UpdateCarsModels(c *gin.Context) {
	var carModel models.CarsModels
	if err := c.ShouldBindJSON(&carModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&carModel).Update(&carModel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var carsModels []models.CarsModels
	config.DB.Preload("MainFactory").Preload("CarsMake").Find(&carsModels)
	c.JSON(200, gin.H{
		"carModel":   carModel,
		"carsModels": carsModels,
	})
}

// IndexCarsModelsWithCarMake ..
func IndexCarsModelsWithCarMake(c *gin.Context) {
	var carsModels []models.CarsModels
	ID := c.Param("id")

	config.DB.Where("cars_make_id = ?", ID).Find(&carsModels)

	c.JSON(200, carsModels)
}

// ------------- CarParts ----------- //

// StoreCarPart..
func StoreCarPart(c *gin.Context) {
	var carParts models.CarParts
	if err := c.ShouldBindJSON(&carParts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&carParts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, carParts)
}

// IndexCarParts ..
func IndexCarParts(c *gin.Context) {
	var carParts []models.CarParts
	config.DB.Find(&carParts)
	c.JSON(200, carParts)
}

// DestroyCarPart ..
func DestroyCarPart(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.CarParts{}, ID)
	// Return Countries list
	var carParts []models.CarParts
	config.DB.Find(&carParts)
	c.JSON(200, carParts)
}

// UpdateCarPart ..
func UpdateCarPart(c *gin.Context) {
	var carPart models.CarParts
	if err := c.ShouldBindJSON(&carPart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&carPart).Update(&carPart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var carParts []models.CarParts
	config.DB.Find(&carParts)
	c.JSON(200, gin.H{
		"carPart":  carPart,
		"carParts": carParts,
	})
}

// ------------- End Car Parts ----------- //

// ------------- Car Fuels ----------- //

// StoreCarFuel..
func StoreCarFuel(c *gin.Context) {
	var carFuel models.CarFuels
	if err := c.ShouldBindJSON(&carFuel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&carFuel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, carFuel)
}

// IndexCarFuels ..
func IndexCarFuels(c *gin.Context) {
	var carFuels []models.CarFuels
	config.DB.Find(&carFuels)
	c.JSON(200, carFuels)
}

// DestroyCarFuels ..
func DestroyCarFuels(c *gin.Context) {
	ID := c.Param("id")
	config.DB.Delete(&models.CarFuels{}, ID)
	// Return Countries list
	var carFuels []models.CarFuels
	config.DB.Find(&carFuels)
	c.JSON(200, carFuels)
}

// UpdateCarFuels ..
func UpdateCarFuels(c *gin.Context) {
	var carFuel models.CarFuels
	if err := c.ShouldBindJSON(&carFuel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&carFuel).Update(&carFuel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var carFuels []models.CarFuels
	config.DB.Find(&carFuels)
	c.JSON(200, gin.H{
		"carFuel":  carFuel,
		"carFuels": carFuels,
	})
}

// ------------- End Main Factory ----------- //
