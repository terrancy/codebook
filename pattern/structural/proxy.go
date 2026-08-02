package structural

// 代理模式（Proxy）：为其他对象提供一种代理以控制对这个对象的访问。
//
// 典型用途：延迟加载、访问控制、日志/缓存、远程调用。下面的例子是“保护代理 + 日志代理”的简化版。

// Subject 真实对象与代理共同实现的接口。
type Subject interface {
	Do() string
}

// RealSubject 真实主题，真正干活的。
type RealSubject struct{}

func (RealSubject) Do() string { return "real work done" }

// Proxy 代理：在调用真实对象前后可插入控制逻辑。
type Proxy struct {
	real      *RealSubject
	hasAccess bool
}

func NewProxy(hasAccess bool) *Proxy {
	return &Proxy{hasAccess: hasAccess}
}

func (p *Proxy) Do() string {
	if !p.hasAccess {
		return "proxy: access denied"
	}
	if p.real == nil {
		p.real = &RealSubject{} // 懒加载真实对象
	}
	return "proxy:before -> " + p.real.Do()
}
