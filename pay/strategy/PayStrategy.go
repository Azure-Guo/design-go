package strategy

import "design-go/pay/pojo"

type PayStrategy interface {
	Pay(body *pojo.PayBody) bool
}

type ZfbStrategy struct{}
type WxStrategy struct{}

func (z *ZfbStrategy) Pay(body *pojo.PayBody) bool {
	// 支付宝策略
	return false
}

func (w *WxStrategy) Pay(body *pojo.PayBody) bool {
	// 微信策略
	return true
}
