package behavioral

// 命令模式（Command）：将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；
// 支持请求的排队、记录日志、以及可撤销的操作。
//
// 典型场景：GUI 按钮、事务、撤销/重做。

// Command 命令接口：统一执行入口。
type Command interface {
	Execute() string
	Undo() string
}

// Light 接收者：真正执行动作的对象。
type Light struct{}

func (Light) On() string  { return "light on" }
func (Light) Off() string { return "light off" }

// TurnOnCommand 具体命令：开灯。
type TurnOnCommand struct {
	light Light
}

func (c TurnOnCommand) Execute() string { return c.light.On() }
func (c TurnOnCommand) Undo() string    { return c.light.Off() }

// TurnOffCommand 具体命令：关灯。
type TurnOffCommand struct {
	light Light
}

func (c TurnOffCommand) Execute() string { return c.light.Off() }
func (c TurnOffCommand) Undo() string    { return c.light.On() }

// Remote 调用者：持有命令，触发执行/撤销。
type Remote struct {
	cmd Command
}

func (r *Remote) SetCommand(c Command) { r.cmd = c }
func (r *Remote) Press() string {
	if r.cmd == nil {
		return "no command"
	}
	return r.cmd.Execute()
}
func (r *Remote) PressUndo() string {
	if r.cmd == nil {
		return "no command"
	}
	return r.cmd.Undo()
}
