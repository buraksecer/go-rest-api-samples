package controllers

import (
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/models"
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/services"
	"github.com/gin-gonic/gin"
)

// We sould keep service in this struct
type TokenController struct {
	service services.TokenService
}

func NewTokenController(service services.TokenService) Controller {
	return &TokenController{service: service}
}

func (t *TokenController) Register(engine *gin.Engine) {
	// We should register method with request type in Gin Engine
	engine.GET("/createToken", t.CreateToken)
}

func (t *TokenController) CreateToken(ctx *gin.Context) {
	var request models.CreateTokenRequest
	// We check to see if there is createTokenRequest in Gin-Gonics.
	err := ctx.Bind(&request)

	// if we did not find request type Gin-Gonics returns error.
	if err != nil {
		ctx.JSON(400, models.CreateTokenResponse{
			Success: false,
			Token:   "Token not created! Request is invalid!",
		})
		return
	}
	// We call create token method from service.
	response := t.service.CreateToken(&request)

	// If service' response failed, we returned fail response with HttpStatus code is 400
	if !response.Success {
		ctx.JSON(400, response)
		return
	}
	// We returned success response with HttpStatus code is 200
	ctx.JSON(200, response)
}
