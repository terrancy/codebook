package behavioral

// 观察者模式（Observer）：定义对象间的一对多依赖，当一个对象状态改变时，所有依赖它的对象都得到通知并自动更新。
//
// 对应「发布-订阅」。Go 里用 slice 持有观察者，遍历通知即可。

// Observer 观察者接口。
type Observer interface {
	Update(msg string)
}

// Subject 被观察者接口。
type Observable interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(msg string)
}

// NewsPublisher 具体被观察者：新闻发布器。
type NewsPublisher struct {
	observers []Observer
}

func (p *NewsPublisher) Attach(o Observer) { p.observers = append(p.observers, o) }

func (p *NewsPublisher) Detach(o Observer) {
	for i, ob := range p.observers {
		if ob == o {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *NewsPublisher) Notify(msg string) {
	for _, o := range p.observers {
		o.Update(msg)
	}
}

// Reader 具体观察者：读者。
type Reader struct {
	Name string
	Got  string
}

func (r *Reader) Update(msg string) { r.Got = msg }
