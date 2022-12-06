package statemachine

import "fmt"

type FSM struct {
	// 持有状态集合
	states map[OrderState]IFSMState
	// 当前状态
	currentState IFSMState
	// 默认状态
	defaultState IFSMState
	// 外部输入数据
	inputData OrderState
	// 是否初始化
	init bool
}

// 初始化FSM
func (this *FSM) Init() {
	this.Reset()
}

// 重置
func (this *FSM) Reset() {
	this.init = false
}

func (this *FSM) AddState(state OrderState, fmsState IFSMState) {
	if this.states == nil {
		this.states = make(map[OrderState]IFSMState)
	}
	this.states[state] = fmsState
}

func (this *FSM) SetDefaultState(fmsState IFSMState) {
	this.defaultState = fmsState
}

// 设置输入数据
func (this *FSM) SetInputData(inputData OrderState) {
	this.inputData = inputData
	this.TransitionState()
}

func (this *FSM) GetCurrentState() IFSMState {
	return this.currentState
}

// 状态转换
func (this *FSM) TransitionState() {
	if this.currentState != nil && this.inputData != this.currentState.NowState()+1 {
		fmt.Printf("越权")
		return
	}
	nextState := this.defaultState
	for _, state := range this.states {
		if this.inputData == state.NowState() {
			nextState = state
			break
		}
	}
	if ok := nextState.CheckTransition(this.inputData); ok {
		if this.currentState != nil {
			this.currentState.Exit()
		}
		this.currentState = nextState
		this.init = true
		this.currentState.Enter()
	}
}

func Init() *FSM {
	createState := NewCreateState(ORDER_CREATE)
	payState := NewPayState(ORDER_PAY)
	sendState := NewSendState(ORDER_SEND)
	receiveState := NewReceiveState(ORDER_RECEIVE)
	fsm := new(FSM)
	fsm.AddState(ORDER_CREATE, createState)
	fsm.AddState(ORDER_PAY, payState)
	fsm.AddState(ORDER_SEND, sendState)
	fsm.AddState(ORDER_RECEIVE, receiveState)
	fsm.SetDefaultState(createState)
	fsm.Init()
	return fsm
}
