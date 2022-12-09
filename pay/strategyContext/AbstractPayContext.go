package strategyContext

import "design-go/pay/pojo"

type AbstractPayContext interface {
	Execute(body *pojo.PayBody) bool
}
