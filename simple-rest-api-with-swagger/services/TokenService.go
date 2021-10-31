package services

import (
	"fmt"
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/models"
	"github.com/buraksecer/go-rest-api-samples/simple-rest-api-with-swagger/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenService interface {
	CreateToken(c *gin.Context, request *models.CreateTokenRequest)
	GetToken(c *gin.Context)
}

type TokenServiceImpl struct {
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}

// CreateToken godoc
// @Summary Create new token
// @Description Create new token
// @ID create-token
// @Accept json
// @Produce json
// @Param create token body models.CreateTokenRequest true "Create new token"
// @Success 200 {object} responses.JSONSuccessResult{data=models.CreateTokenResponse,code=int,message=string}
// @Success 400 {object} responses.JSONSuccessResult{code=int,message=string}
// @Success 500 {object} responses.JSONInternalErrorResult{code=int,message=string}
// @Router /create-token [post]
func (t *TokenServiceImpl) CreateToken(c *gin.Context, request *models.CreateTokenRequest) {
	fmt.Println("Create Token Request, username:" + request.Username + " Password: " + request.Password)
	// We created fail case response
	response := &models.CreateTokenResponse{
		Success: false,
		Token:   "Can not created token! Username password is incorrect!",
	}

	if "buraksecer" == request.Username && "1" == request.Password {
		response.Success = true
		response.Token = uuid.New().String()
		// We validated username and password and we returned success case response.
		responses.SuccessResponse(c, response)
		return
	}
	// We did not validate username and password and we returned fail case response.
	responses.FailResponse(c, 400, "Something Wrong!")
}

// GetToken godoc
// @Summary Get token
// @Description Get token
// @ID get-token
// @Accept json
// @Produce json
// @Success 200 {object} responses.JSONSuccessResult{data=models.CreateTokenResponse,code=int,message=string}
// @Success 400 {object} responses.JSONSuccessResult{code=int,message=string}
// @Success 500 {object} responses.JSONInternalErrorResult{code=int,message=string}
// @Router /get-token [get]
func (t *TokenServiceImpl) GetToken(c *gin.Context)  {
	fmt.Println("Get token request")
	responses.SuccessResponse(c,uuid.New().String())
}
