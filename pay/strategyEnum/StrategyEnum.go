package strategyEnum

import (
	"design-go/pay/strategy"
	"reflect"
)

type StrategyEnum int

const (
	ZFB = iota
	WX
)

func (s StrategyEnum) GetType() reflect.Type {
	return [...]reflect.Type{
		reflect.TypeOf(strategy.ZfbStrategy{}),
		reflect.TypeOf(strategy.WxStrategy{}),
	}[s]
}
