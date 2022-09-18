package router

import (
	"design-go/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/pay", controller.Pay)
	return router
}
