package strategyContext

import (
	"design-go/pay/pojo"
	"design-go/pay/strategy"
)

type PayContext struct {
	PayStrategy strategy.PayStrategy
}

func (c *PayContext) Execute(body *pojo.PayBody) bool {
	return c.PayStrategy.Pay(body)
}
