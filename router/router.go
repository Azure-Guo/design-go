package router

import (
	"design-go/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 创建订单
	router.GET("/create", controller.Create)
	// Pay 支付模块,策略+策略枚举+门面+工厂
	router.POST("/pay", controller.Pay)
	// 发货
	router.GET("/send", controller.Send)
	// 收货
	router.GET("receive", controller.Receive)
	// Suggest 投放模块,责任链,实现实时修改删除责任链的位置
	router.GET("/suggest", controller.Suggest)
	return router
}
