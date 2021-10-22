package startups

import (
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/controllers"
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/services"
)

func Initialize() *[]controllers.Controller {
	// Array for controllers with Controller Interface
	var list []controllers.Controller

	// Token Service
	// This method returns TokenServiceImpl
	tokenService := services.NewTokenService()

	// We put Token Service in Token controller
	// This method returns TokenControllerImpl
	trackingController := controllers.NewTokenController(tokenService)

	// We append list our controllers
	list = append(list, trackingController)
	return &list
}
