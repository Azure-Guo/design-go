package decorator

import (
	"design-go/pay/pojo"
	"design-go/pay/strategyContext"
	"fmt"
)

type Decorator interface {
	ExtraFunction(body *pojo.PayBody)
}

type AbstractDecorator struct {
	AbstractDecorator strategyContext.PayContext
}

// 老活
func (this *AbstractDecorator) Execute(body *pojo.PayBody) bool {
	return this.AbstractDecorator.Execute(body)
}

// 新活
func (this *AbstractDecorator) ExtraFunction(body *pojo.PayBody) {

}

type AddDecorator struct {
	AbstractDecorator
}

func (this *AddDecorator) ExtraFunction(body *pojo.PayBody) {
	product := body.Product
	// 从db里获取product详细信息
	// 从配置中心里获取产品的更新策略
	// 根据策略跟新用户平台币，或者红包
	fmt.Println("更新平台币成功", product)
	fmt.Println("发送红包成功", product)
}
func (this *AddDecorator) Execute(body *pojo.PayBody) bool {
	execute := this.AbstractDecorator.Execute(body)
	if execute {
		this.ExtraFunction(body)
	} else {
		// 重试机制
	}
	return execute
}
