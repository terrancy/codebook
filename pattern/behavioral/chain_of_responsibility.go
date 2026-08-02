package behavioral

// 责任链模式（Chain of Responsibility）：使多个对象都有机会处理请求，从而避免请求的发送者与接收者耦合。
// 将这些对象连成一条链，并沿着链传递请求，直到有一个对象处理它为止。
//
// 典型场景：中间件、审批流、日志级别过滤。

// Handler 处理者接口。
type Handler interface {
	SetNext(h Handler) Handler
	Handle(req string) string
}

// baseHandler 提供链式串联的通用实现，具体处理者嵌入它即可。
type baseHandler struct {
	next Handler
}

func (b *baseHandler) SetNext(h Handler) Handler {
	b.next = h
	return h
}

// AuthHandler 认证处理者。
type AuthHandler struct{ baseHandler }

func (h *AuthHandler) Handle(req string) string {
	if req == "unauthorized" {
		return "AuthHandler: rejected (no auth)"
	}
	if h.next != nil {
		return h.next.Handle(req)
	}
	return "AuthHandler: passed"
}

// LogHandler 日志处理者。
type LogHandler struct{ baseHandler }

func (h *LogHandler) Handle(req string) string {
	_ = "LogHandler: logged request " + req // 实际中写日志
	if h.next != nil {
		return h.next.Handle(req)
	}
	return "LogHandler: done"
}
