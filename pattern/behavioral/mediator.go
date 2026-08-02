package behavioral

// 中介者模式（Mediator）：用一个中介对象来封装一系列对象之间的交互，使得各对象不需要显式地相互引用，
// 从而降低耦合，并且可以独立地改变它们之间的交互。
//
// 典型场景：聊天室、机场调度、UI 组件联动。

// Colleague 同事接口。
type Colleague interface {
	Send(msg string)
	Receive(msg string)
	Name() string
}

// Mediator 中介者接口。
type Mediator interface {
	Register(c Colleague)
	Relay(from, msg string)
}

// ChatRoom 具体中介者：群聊房间。
type ChatRoom struct {
	members map[string]Colleague
}

func NewChatRoom() *ChatRoom { return &ChatRoom{members: map[string]Colleague{}} }

func (r *ChatRoom) Register(c Colleague) { r.members[c.Name()] = c }

// Relay 把消息转发给除发送者外的所有成员，解耦收发双方。
func (r *ChatRoom) Relay(from, msg string) {
	for name, c := range r.members {
		if name != from {
			c.Receive(msg)
		}
	}
}

// User 具体同事。
type User struct {
	name    string
	room    *ChatRoom
	inbox   []string
}

func NewUser(name string, room *ChatRoom) *User {
	u := &User{name: name, room: room}
	room.Register(u)
	return u
}

func (u *User) Name() string                          { return u.name }
func (u *User) Send(msg string)                       { u.room.Relay(u.name, msg) }
func (u *User) Receive(msg string)                    { u.inbox = append(u.inbox, msg) }
func (u *User) Inbox() []string                       { return u.inbox }
