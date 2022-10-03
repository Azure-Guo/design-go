package controller

import (
	"design-go/pay/pojo"
	"design-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pay(c *gin.Context) {
	var body pojo.PayBody
	c.Bind(&body)
	pay := service.Pay(body)
	c.JSON(http.StatusOK, gin.H{"data": pay})
}
func Suggest(c *gin.Context) {
	var userName string
	c.Bind(&userName)
	suggestList := service.Suggest(userName)
	c.JSON(http.StatusOK, gin.H{"data": suggestList})
}
