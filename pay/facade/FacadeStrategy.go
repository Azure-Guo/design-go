package facade

import (
	"design-go/pay/decorator"
	"design-go/pay/pojo"
	"design-go/pay/strategyContext"
	"design-go/pay/strategyEnum"
	"design-go/pay/strategyFactory"
)

func Pay(body pojo.PayBody) bool {
	enum := strategyEnum.StrategyEnum(body.Type)
	payStrategy := strategyFactory.GetPayStrategy(enum)
	context := strategyContext.PayContext{PayStrategy: payStrategy}
	addDecorator := decorator.AddDecorator{decorator.AbstractDecorator{
		AbstractDecorator: context,
	}}
	execute := addDecorator.Execute(&body)
	return execute
}
