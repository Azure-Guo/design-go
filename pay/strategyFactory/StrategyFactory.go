package strategyFactory

import (
	"design-go/pay/strategy"
	"design-go/pay/strategyEnum"
	"reflect"
)

func GetPayStrategy(enum strategyEnum.StrategyEnum) strategy.PayStrategy {
	refVal := reflect.New(enum.GetType())
	payStrategy, ok := refVal.Interface().(strategy.PayStrategy)
	if !ok {
		return nil
	}
	return payStrategy
}
