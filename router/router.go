package router

import (
	"design-go/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// Pay 支付模块,策略+策略枚举+门面+工厂
	router.POST("/pay", controller.Pay)
	// Suggest 投放模块,责任链,实现实时修改删除责任链的位置
	router.GET("/suggest", controller.Suggest)
	return router
}
