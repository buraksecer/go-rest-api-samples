package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type JSONBadRequestResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong Parameters"`
	Data    interface{} `json:"data"`
}

type JSONInternalErrorResult struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"Internal Error"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JSONSuccessResult{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Success",
	})
}

func FailResponse(c *gin.Context, resposeCode int, message string) {
	switch resposeCode {
	case http.StatusInternalServerError:
		c.JSON(resposeCode, JSONInternalErrorResult{
			Code:    resposeCode,
			Data:    nil,
			Message: message,
		})
	default:
		c.JSON(resposeCode, JSONBadRequestResult{
			Code:    resposeCode,
			Data:    nil,
			Message: message,
		})
	}
}
