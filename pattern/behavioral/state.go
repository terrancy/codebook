package behavioral

// 状态模式（State）：允许一个对象在其内部状态改变时改变它的行为，对象看起来似乎修改了它的类。
//
// 与策略的区别：状态是“自身状态驱动的自动切换”，策略是“外部显式指定”。
// 典型场景：订单状态机、播放器状态。

// State 状态接口。
type State interface {
	Handle(ctx *Context) string
}

// Context 上下文：持有当前状态，并把请求委派给它。
type Context struct {
	state State
}

func NewContext(s State) *Context { return &Context{state: s} }
func (c *Context) SetState(s State) { c.state = s }
func (c *Context) Request() string { return c.state.Handle(c) }

// 具体状态 A：待支付
type PendingState struct{}

func (PendingState) Handle(ctx *Context) string {
	ctx.SetState(PaidState{}) // 自动流转到已支付
	return "pending -> paid"
}

// 具体状态 B：已支付
type PaidState struct{}

func (PaidState) Handle(ctx *Context) string {
	ctx.SetState(DoneState{})
	return "paid -> done"
}

// 具体状态 C：已完成（终态）
type DoneState struct{}

func (DoneState) Handle(_ *Context) string { return "already done" }
