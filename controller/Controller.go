package controller

import (
	"design-go/pay/pojo"
	body "design-go/pojo"
	"design-go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var b body.Order
	c.Bind(&b)
	fmt.Println(b)
	create := service.Create(b.OrderId)
	c.JSON(http.StatusOK, gin.H{
		"data": create,
	})
}

func Pay(c *gin.Context) {
	var body pojo.PayBody
	c.Bind(&body)
	pay := service.Pay(body)
	c.JSON(http.StatusOK, gin.H{"data": pay})
}

func Send(c *gin.Context) {
	send := service.Send()
	c.JSON(http.StatusOK, gin.H{"data": send})
}

func Receive(c *gin.Context) {
	receive := service.Receive()
	c.JSON(http.StatusOK, gin.H{"data": receive})
}

func Suggest(c *gin.Context) {
	var userName string
	c.Bind(&userName)
	suggestList := service.Suggest(userName)
	c.JSON(http.StatusOK, gin.H{"data": suggestList})
}
