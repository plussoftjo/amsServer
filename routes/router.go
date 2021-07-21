// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// gin.SetMode(gin.ReleaseMode)
	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})
	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.POST("/app/register", controllers.AppRegisterController)
	auth.POST("/app/login", controllers.AppLoginController)
	auth.GET("/app/auth", controllers.AuthAppUser)
	auth.POST("/app/changePassword", controllers.ChangePassword)
	auth.GET("/auth", controllers.Auth)
	auth.GET("/users/index", controllers.UsersListIndex)
	auth.GET("/users/delete/:id", controllers.DeleteUser)
	auth.POST("/update", controllers.UpdateUser)
	auth.POST("/app/update", controllers.AppUpdateUser)
	auth.POST("/checkHasPhone", controllers.CheckIfHasPhone)
	auth.POST("/resetPassword", controllers.ResetPassword)

	// --------- Basics ------- //
	basics := r.Group("/basics")

	// UploadImage => For All
	basics.POST("/upload_image/:imageType", controllers.UpdateImage)

	// --------- User Controller ----------------- //
	user := r.Group("/users")
	// ~~~ User Roles ~~~ //
	user.POST("/roles/store", controllers.StoreUserRoles)
	user.POST("/roles/update", controllers.UpdateUserRole)
	user.GET("/roles/index", controllers.IndexUserRoles)
	user.GET("/roles/delete/:id", controllers.DeleteUserRole)
	// --------------- Employ Controller ----------- //
	user.POST("/employee/store", controllers.StoreEmployee)
	user.GET("/employee/index", controllers.IndexEmployee)
	user.GET("/employee/delete/:id", controllers.DeleteEmployee)
	user.POST("/employee/update", controllers.UpdateEmployee)

	// ---------- Countries & Cites & Areas ------------ //
	countries := r.Group("/countries")
	countries.POST("/store", controllers.StoreCountry)
	countries.GET("/index", controllers.IndexCountries)
	countries.GET("/destroy/:id", controllers.DestroyCountry)
	countries.POST("/update", controllers.UpdateCountry)

	cites := r.Group("/cites")
	cites.POST("/store", controllers.StoreCity)
	cites.GET("/index", controllers.IndexCites)
	cites.GET("/destroy/:id", controllers.DestroyCity)
	cites.POST("/update", controllers.UpdateCity)
	cites.GET("/indexCitesWithCountryID/:id", controllers.IndexCitesWithCountryID)

	areas := r.Group("/areas")
	areas.POST("/store", controllers.StoreArea)
	areas.GET("/index", controllers.IndexAreas)
	areas.GET("/destroy/:id", controllers.DestroyArea)
	areas.POST("/update", controllers.UpdateArea)
	areas.GET("/indexAreasWithCityID/:id", controllers.IndexAreasWithCityID)

	// ---------- Cars ----------- //
	// Main Factories
	mainFactories := r.Group("/mainFactories")
	mainFactories.POST("/store", controllers.StoreMainFactory)
	mainFactories.GET("/index", controllers.IndexMainFactory)
	mainFactories.GET("/destroy/:id", controllers.DestroyMainFactory)
	mainFactories.POST("/update", controllers.UpdateMainFactory)
	// CarsMake .
	carsMake := r.Group("/carsMake")
	carsMake.POST("/store", controllers.StoreCarMake)
	carsMake.GET("/index", controllers.IndexCarsMake)
	carsMake.GET("/destroy/:id", controllers.DestroyCarMake)
	carsMake.POST("/update", controllers.UpdateCarsMake)
	carsMake.GET("/indexCarsMakeWithMainFactoryID/:id", controllers.IndexCarsMakeWithMainFactory)
	// CarModels
	carsModels := r.Group("/carsModels")
	carsModels.POST("/store", controllers.StoreCarModel)
	carsModels.GET("/index", controllers.IndexCarsModels)
	carsModels.GET("/destroy/:id", controllers.DestroyCarModel)
	carsModels.POST("/update", controllers.UpdateCarsModels)
	carsModels.GET("/indexCarsModelsWithCarMakeID/:id", controllers.IndexCarsModelsWithCarMake)
	// CarParts
	carParts := r.Group("/carParts")
	carParts.POST("/store", controllers.StoreCarPart)
	carParts.GET("/index", controllers.IndexCarParts)
	carParts.GET("/destroy/:id", controllers.DestroyCarPart)
	carParts.POST("/update", controllers.UpdateCarPart)
	// CarFuels
	carFuels := r.Group("/carFuels")
	carFuels.POST("/store", controllers.StoreCarFuel)
	carFuels.GET("/index", controllers.IndexCarFuels)
	carFuels.GET("/destroy/:id", controllers.DestroyCarFuels)
	carFuels.POST("/update", controllers.UpdateCarFuels)

	// -------- Services ---------- //
	services := r.Group("/services")
	services.POST("/store", controllers.StoreService)
	services.GET("/index", controllers.IndexServices)
	services.GET("/destroy/:id", controllers.DestroyService)
	services.POST("/update", controllers.UpdateService)

	// Sub Services
	subServices := r.Group("/subServices")
	subServices.POST("/store", controllers.StoreSubService)
	subServices.GET("/index", controllers.IndexSubServices)
	subServices.GET("/destroy/:id", controllers.DestroySubService)
	subServices.GET("/indexWithServiceID/:id", controllers.IndexSubServicesWithServiceID)
	subServices.POST("/update", controllers.UpdateSubService)

	// Services Options
	servicesOptions := r.Group("/servicesOptions")
	servicesOptions.POST("/store", controllers.StoreServicesOptions)
	servicesOptions.GET("/index", controllers.IndexServicesOptions)
	servicesOptions.GET("/destroy/:id", controllers.DestroyServiceOptions)
	servicesOptions.POST("/update", controllers.UpdateServicesOptions)

	// ----------------------- App Intro Controller -------------------------------- //
	appIntro := r.Group("/appIntro")
	appIntro.GET("/countries/index", controllers.IndexAppIntroCountry)
	appIntro.GET("/services/index", controllers.IndexAppIntroServices)
	appIntro.GET("/specialtyData/index", controllers.IndexAppIntroSpecialtyData)
	appIntro.GET("/cars/withMainFactoryID/:id", controllers.IndexAppIntroCarsWithMainFactoryID)
	appIntro.POST("/supplierSpecialty/store", controllers.StoreSupplierSpecialty)
	appIntro.POST("/supplierSpecialty/update", controllers.UpdateSupplierSpecialty)
	appIntro.POST("/supplierPersonalDetails/store", controllers.StoreSupplierPersonalDetails)
	appIntro.POST("/supplierPersonalDetails/update", controllers.UpdateSupplierPersonalDetails)
	appIntro.POST("/updateUserWithServiceID/store", controllers.StoreUpdateUserWithServicesID)
	// Clients App Intro
	appIntro.POST("/storeClientCar", controllers.StoreClientsCar)

	// --------------- Supplier Request --------------------- //
	supplierRequest := r.Group("/supplierRequest")
	supplierRequest.GET("/IndexMainFactoryForSupplierRequest/index", controllers.IndexMainFactoryForSupplierRequest)
	supplierRequest.POST("/store", controllers.StoreSupplierRequestPart)
	supplierRequest.GET("/test", controllers.TestIndexSupplierRequestPart)
	supplierRequest.GET("/indexForSupplier/:id", controllers.IndexSupplierRequestPartForSuppliers)
	supplierRequest.GET("/IndexSupplierMyRequestPart/:id", controllers.IndexSupplierMyRequestPart)
	supplierRequest.POST("/supplierOfferPrice/store", controllers.StoreSupplierOfferPrice)
	supplierRequest.POST("/UpdateTakenOfferPriceForRequestPart", controllers.UpdateTakenOfferPriceForRequestPart)
	supplierRequest.GET("/endSupplierRequestPart/:id", controllers.EndSupplierRequestPart)

	// --------------- Clients Part Request --------------------- //
	clientsPartRequest := r.Group("/clientsPartRequest")
	clientsPartRequest.POST("/store", controllers.StoreClientsPartRequest)
	clientsPartRequest.GET("/indexForSuppliers/:id", controllers.IndexClientsPartRequestForSupplier)
	clientsPartRequest.POST("/storeOfferPrice", controllers.StoreSupplierOfferPriceForClients)
	clientsPartRequest.POST("/takeOffer", controllers.ClientPartRequestTakeOfferPrice)
	clientsPartRequest.GET("/indexForClients/:id", controllers.IndexClientsPartRequest)
	clientsPartRequest.GET("/endRequest/:id", controllers.ClientsEndPartRequest)

	// --------------- Auto Mobile Services --------------------- //
	autoMobileServices := r.Group("/ams")
	autoMobileServices.GET("/indexServices", controllers.IndexAutoMobileServiceServices)
	autoMobileServices.GET("/indexServiceOptionsWithSubServiceID/:id", controllers.IndexServiceOptionsWithSubServiceID)
	autoMobileServices.POST("/store", controllers.StoreAutoMobileService)
	autoMobileServices.GET("/indexForSupplier/:id", controllers.IndexAutoMobileServiceForSupplier)
	autoMobileServices.GET("/indexForClients/:id", controllers.IndexAutoMobileServiceForUser)
	autoMobileServices.POST("/storeSupplierAmrOffer", controllers.StoreSupplierAmrOffer)
	autoMobileServices.POST("/approveFromSupplier", controllers.SupplierTakeAMS)
	autoMobileServices.POST("/endingAms", controllers.EndingAms)
	autoMobileServices.GET("/endAmsFromClient/:id", controllers.EndAmsFromClient)

	// --------------- WorkshopList --------------------- //
	workshopList := r.Group("/workshopList")
	workshopList.GET("/indexSubServices", controllers.IndexServicesForWorkshoplist)
	workshopList.POST("/index", controllers.IndexWorkshopList)

	r.Run(":8082")
}
