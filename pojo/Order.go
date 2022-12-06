package pojo

import "design-go/order/statemachine"

type Order struct {
	OrderId int `form:"id"`
	State   statemachine.OrderState
}
