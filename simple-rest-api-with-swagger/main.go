package main

import (
	"fmt"
	_ "github.com/buraksecer/go-rest-api-samples/docs"
	"github.com/buraksecer/go-rest-api-samples/startups"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Go Project
// @version 1.0
// @description Go Project
// @termsOfService http://swagger.io/terms/

// @contact.name Burak SEÇER
// @contact.url http://www.buraksecer.com
// @contact.email burakscr@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
//go:generate swag init
func main() {
	//(http://localhost:8080/swagger/index.html)
	//username:burak
	//password:1
	fmt.Println("Api is starting...")

	router := gin.Default()
	router.RouterGroup.Handlers = router.RouterGroup.Handlers[0:0]
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	controllerArray := startups.Initialize()

	for _, key := range *controllerArray {
		// We register all controller and service
		key.Register(router)
	}
	// This configurations for swagger.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	router.Run()
}
