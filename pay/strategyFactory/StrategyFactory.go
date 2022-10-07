package strategyFactory

import (
	"design-go/pay/strategy"
	"design-go/pay/strategyEnum"
	"reflect"
	"sync"
)

var PayStrategyMap sync.Map = sync.Map{}

func GetPayStrategy(enum strategyEnum.StrategyEnum) strategy.PayStrategy {
	payStrategy, ex := PayStrategyMap.Load(enum.GetType())
	if !ex {
		refVal := reflect.New(enum.GetType())
		payStrategy, _ = refVal.Interface().(strategy.PayStrategy)
		PayStrategyMap.Store(enum.GetType(), payStrategy)
	}
	return payStrategy.(strategy.PayStrategy)
}
