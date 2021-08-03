// Package controllers ...
package controllers

import (
	"server/config"
	"server/models"

	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

// StoreClientsPartRequest ..
func StoreClientsPartRequest(c *gin.Context) {
	var clientsPartRequest models.ClientsPartRequest
	if err := c.ShouldBindJSON(&clientsPartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&clientsPartRequest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var withDetails models.ClientsPartRequest
	config.DB.Where("id = ?", clientsPartRequest.ID).Scopes(models.ClientsPartRequestWithDetails).First(&withDetails)

	c.JSON(200, withDetails)
}

// IndexClientsPartRequest ..
func IndexClientsPartRequest(c *gin.Context) {
	ID := c.Param("id")
	var clientsPartRequest []models.ClientsPartRequest
	config.DB.Where("user_id = ?", ID).Scopes(models.ClientsPartRequestWithDetails).Order("id desc").Find(&clientsPartRequest)

	c.JSON(200, clientsPartRequest)
}

// IndexClientsPartRequestForSupplier ..
func IndexClientsPartRequestForSupplier(c *gin.Context) {
	UserID := c.Param("id")

	// GET SUPPLIER SPECIALTY WITH USERID
	var supplierSpecialty models.SupplierSpecialty
	config.DB.Where("user_id = ?", UserID).First(&supplierSpecialty)

	// GET PERSONAL DETAILS FOR THE SUPPLIER WITH USER_ID
	var supplierPersonalDetails models.SupplierPersonalDetails
	config.DB.Where("user_id = ?", UserID).First(&supplierPersonalDetails)

	// THERE DETAILS IS VERY IMPORTANT => :CARMAKE :CITYID :AREAID
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
		carsMakeIDList = append(carsMakeIDList, strconv.FormatInt(carMakeID, 10))
		// fmt.Println(carMakeID)
	}

	// NOW MAKE IT STRING
	carMakeIDsString := strings.Join(carsMakeIDList, ",")
	// AND REGISTER THE CITY AND AREA ID
	cityID := supplierPersonalDetails.CityID
	areaID := supplierPersonalDetails.AreaID

	var clientsPartRequest []models.ClientsPartRequest
	// NOW MAKE QUERY
	err := config.DB.
		Where("city_id = ? OR city_id = 0", cityID).
		Where("area_id = ? OR area_id = 0", areaID).
		Where("ending = ?", false).
		Where("car_make_id IN " + "(" + carMakeIDsString + ")").
		Scopes(models.ClientsPartRequestWithDetails).
		Order("id desc").
		Limit(5).
		Find(&clientsPartRequest).Error
	if err != nil {
		c.JSON(500, err)
		return
	}

	// Register The full supplier request part with user
	var fullClientsPartRequest []models.ClientsPartRequest

	for _, cpr := range clientsPartRequest {

		// Check if have offer price
		var supplierOfferPrice models.SupplierOfferPriceForClients
		err := config.DB.Where("user_id = ?", UserID).Where("clients_part_request_id = ?", cpr.ID).Find(&supplierOfferPrice).Error
		if err != nil {
			cpr.SupplierTakeOffer = false
		} else {
			cpr.SupplierTakeOffer = true
			cpr.SupplierOfferPrice = supplierOfferPrice

			supplierOfferPrice.Seen = true
			config.DB.Save(&supplierOfferPrice)
		}
		fullClientsPartRequest = append(fullClientsPartRequest, cpr)

	}

	c.JSON(200, gin.H{
		"clientsRequestPartList": fullClientsPartRequest,
	})
}

// StoreSupplierOfferPriceForClients ..
func StoreSupplierOfferPriceForClients(c *gin.Context) {
	var SupplierOfferPriceForClients models.SupplierOfferPriceForClients
	if err := c.ShouldBindJSON(&SupplierOfferPriceForClients); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// StoreInDB
	if err := config.DB.Create(&SupplierOfferPriceForClients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, SupplierOfferPriceForClients)
}

// ClientPartRequestTakeOfferPrice ..
func ClientPartRequestTakeOfferPrice(c *gin.Context) {
	type ClientPartRequestTakeOfferPriceType struct {
		ClientPartRequestID  uint   `json:"clientPartRequestID"`
		SupplierOfferPriceID uint   `json:"supplierOfferPriceID"`
		BookingType          int64  `json:"bookingType"`
		BookingDate          string `json:"bookingDate"`
		BookingTime          string `json:"bookingTime"`
	}

	var data ClientPartRequestTakeOfferPriceType
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var clientPartRequest models.ClientsPartRequest
	config.DB.Where("id = ?", data.ClientPartRequestID).Scopes(models.ClientsPartRequestWithDetails).First(&clientPartRequest)

	clientPartRequest.TakenOfferID = data.SupplierOfferPriceID
	clientPartRequest.BookingType = data.BookingType
	clientPartRequest.BookingDate = data.BookingDate
	clientPartRequest.BookingTime = data.BookingTime

	config.DB.Save(&clientPartRequest)

	var newClientPartRequest models.ClientsPartRequest
	config.DB.Where("id = ?", data.ClientPartRequestID).Scopes(models.ClientsPartRequestWithDetails).First(&newClientPartRequest)

	c.JSON(200, gin.H{
		"message":           "Success taken",
		"code":              200,
		"clientPartRequest": newClientPartRequest,
	})

}

// ClientsEndPartRequest ..
func ClientsEndPartRequest(c *gin.Context) {
	ID := c.Param("id")

	var clientPartRequest models.ClientsPartRequest
	config.DB.Where("id = ?", ID).First(&clientPartRequest)

	clientPartRequest.Ending = true
	config.DB.Save(&clientPartRequest)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
