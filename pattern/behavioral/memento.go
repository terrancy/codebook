package behavioral

// 备忘录模式（Memento）：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，
// 以便以后可将对象恢复到原先保存的状态。
//
// 典型场景：编辑器撤销、游戏存档。

// Memento 备忘录：仅存储状态，不暴露内部结构（由 Originator 私有构造）。
type Memento struct {
	state string
}

// Originator 发起者：创建备忘录并从备忘录恢复。
type Originator struct {
	state string
}

func (o *Originator) SetState(s string) { o.state = s }
func (o *Originator) State() string     { return o.state }

// Save 把当前状态快照存入备忘录。
func (o *Originator) Save() Memento { return Memento{state: o.state} }

// Restore 从备忘录恢复状态。
func (o *Originator) Restore(m Memento) { o.state = m.state }

// Caretaker 管理者：负责保存备忘录（不知道其内部内容）。
type Caretaker struct {
	history []Memento
}

func (c *Caretaker) Push(m Memento)   { c.history = append(c.history, m) }
func (c *Caretaker) Pop() Memento {
	if len(c.history) == 0 {
		return Memento{}
	}
	m := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]
	return m
}
