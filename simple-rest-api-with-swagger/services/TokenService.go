package services

import (
	"fmt"
	"github.com/buraksecer/go-rest-api-samples/models"
	"github.com/google/uuid"
)

type TokenService interface {
	CreateToken(request *models.CreateTokenRequest) *models.CreateTokenResponse
}

type TokenServiceImpl struct {
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}

func (t *TokenServiceImpl) CreateToken(request *models.CreateTokenRequest) *models.CreateTokenResponse {
	fmt.Println("Create Token Request, username:" + request.Username + " Password: " + request.Password)
	// We created fail case response
	response := &models.CreateTokenResponse{
		Success: false,
		Token:   "Can not created token! Username password is incorrect!",
	}

	if "caglar" == request.Username && "12345" == request.Password {
		response.Success = true
		response.Token = uuid.New().String()
		// We validated username and password and we returned success case response.
		return response
	}
	// We did not validate username and password and we returned fail case response.
	return response
}
