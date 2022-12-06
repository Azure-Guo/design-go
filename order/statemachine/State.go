package statemachine

import "fmt"

type OrderState int

const (
	ORDER_CREATE  OrderState = iota
	ORDER_PAY     OrderState = iota
	ORDER_SEND    OrderState = iota
	ORDER_RECEIVE OrderState = iota
)

// 接口
type IFSMState interface {
	Enter()
	Exit()
	CheckTransition(state OrderState) bool
	NowState() OrderState
}

// 创建订单
type CreateState struct {
	state OrderState
}

func NewCreateState(state OrderState) *CreateState {
	return &CreateState{state: state}
}
func (this *CreateState) Enter() {
	fmt.Println("Entry CreateState")
}
func (this *CreateState) Exit() {
	fmt.Println("Exit CreateState")
}
func (this *CreateState) NowState() OrderState {
	return this.state
}
func (this *CreateState) CheckTransition(state OrderState) bool {
	return this.state == state
}

// 支付
type PayState struct {
	state OrderState
}

func NewPayState(state OrderState) *PayState {
	return &PayState{state: state}
}
func (this *PayState) Enter() {
	fmt.Println("Entry PayState")
}
func (this *PayState) Exit() {
	fmt.Println("Exit PayState")
}
func (this *PayState) NowState() OrderState {
	return this.state
}
func (this *PayState) CheckTransition(state OrderState) bool {
	return this.state == state
}

// 发货
type SendState struct {
	state OrderState
}

func NewSendState(state OrderState) *SendState {
	return &SendState{state: state}
}
func (this *SendState) Enter() {
	fmt.Println("Entry SendState")
}
func (this *SendState) Exit() {
	fmt.Println("Exit SendState")
}
func (this *SendState) NowState() OrderState {
	return this.state
}
func (this *SendState) CheckTransition(state OrderState) bool {
	return this.state == state
}

// 收货
type ReceiveState struct {
	state OrderState
}

func NewReceiveState(state OrderState) *ReceiveState {
	return &ReceiveState{state: state}
}
func (this *ReceiveState) Enter() {
	fmt.Println("Entry ReceiveState")
}
func (this *ReceiveState) Exit() {
	fmt.Println("Exit ReceiveState")
}
func (this *ReceiveState) NowState() OrderState {
	return this.state
}
func (this *ReceiveState) CheckTransition(state OrderState) bool {
	return this.state == state
}
