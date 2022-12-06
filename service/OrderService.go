package service

import (
	"design-go/order/statemachine"
	"design-go/pay/facade"
	"design-go/pay/pojo"
	body "design-go/pojo"
)

var global *statemachine.FSM
var list []*body.Order

func getFSM() *statemachine.FSM {
	if global != nil {
		return global
	}
	global = statemachine.Init()
	return global
}

func Create(id int) *body.Order {
	fsm := getFSM()
	fsm.SetInputData(statemachine.ORDER_CREATE)
	if list == nil {
		list = make([]*body.Order, 0)
		list = append(list, &body.Order{
			OrderId: id,
			State:   fsm.GetCurrentState().NowState(),
		})
	}
	return list[len(list)-1]
}

func Pay(body pojo.PayBody) *body.Order {
	fsm := getFSM()
	success := facade.Pay(body)
	if success {
		saveToDb(body)
		fsm.SetInputData(statemachine.ORDER_PAY)
	}
	list[len(list)-1].State = fsm.GetCurrentState().NowState()
	return list[len(list)-1]
}
func saveToDb(body pojo.PayBody) {
}

func Send() *body.Order {
	fsm := getFSM()
	fsm.SetInputData(statemachine.ORDER_SEND)
	list[len(list)-1].State = fsm.GetCurrentState().NowState()
	return list[len(list)-1]
}

func Receive() *body.Order {
	fsm := getFSM()
	fsm.SetInputData(statemachine.ORDER_RECEIVE)
	list[len(list)-1].State = fsm.GetCurrentState().NowState()
	return list[len(list)-1]
}
