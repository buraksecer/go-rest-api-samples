package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(engine *gin.Engine)
}
