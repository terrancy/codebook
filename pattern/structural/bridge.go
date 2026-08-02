package structural

// 桥接模式（Bridge）：将抽象部分与其实现部分分离，使它们都可以独立地变化。
//
// 与“适配器”不同：桥接是在设计之初就把抽象和实现解耦，各自演化；适配器是事后补救不兼容接口。

// Implementor 实现维度（可变）：发送消息的渠道。
type Implementor interface {
	Send(msg string) string
}

// 具体实现 A：邮件渠道
type EmailSender struct{}

func (EmailSender) Send(msg string) string { return "email:" + msg }

// 具体实现 B：短信渠道
type SMSSender struct{}

func (SMSSender) Send(msg string) string { return "sms:" + msg }

// Abstraction 抽象维度（可变）：消息类型，持有实现。
type Abstraction struct {
	impl Implementor
}

func NewAbstraction(impl Implementor) *Abstraction { return &Abstraction{impl: impl} }

// Operation 抽象侧的行为，委派给实现侧。
func (a *Abstraction) Operation(msg string) string {
	return a.impl.Send(msg)
}

// RefinedAbstraction 可扩展的抽象（如紧急消息加前缀）。
type UrgentAbstraction struct{ Abstraction }

func (u UrgentAbstraction) Operation(msg string) string {
	return "[URGENT] " + u.impl.Send(msg)
}
