package structural

import "testing"

// Adapter 适配器模式：将不兼容接口转换为目标接口
func TestAdapter(t *testing.T) {
	var target Target = NewAdapter(Adaptee{})
	if got := target.Request(); got != "adapter:adaptee specific" {
		t.Fatalf("adapter got %q", got)
	}
}

// Decorator 装饰器模式：动态为对象添加职责
func TestDecorator(t *testing.T) {
	base := SimpleCoffee{}
	c := NewSugar(NewMilk(base))
	if c.Cost() != 14 { // 10 + 3 + 1
		t.Fatalf("expected cost 14, got %d", c.Cost())
	}
	if c.Desc() != "coffee+milk+sugar" {
		t.Fatalf("unexpected desc: %s", c.Desc())
	}
}

// Proxy 代理模式：为对象提供代理以控制访问
func TestProxy(t *testing.T) {
	p := NewProxy(false)
	if got := p.Do(); got != "proxy: access denied" {
		t.Fatalf("denied case got %q", got)
	}
	p2 := NewProxy(true)
	if got := p2.Do(); got != "proxy:before -> real work done" {
		t.Fatalf("granted case got %q", got)
	}
}

// Bridge 桥接模式：分离抽象与实现，使其独立变化
func TestBridge(t *testing.T) {
	a := NewAbstraction(EmailSender{})
	if got := a.Operation("hi"); got != "email:hi" {
		t.Fatalf("email bridge got %q", got)
	}
	u := UrgentAbstraction{*NewAbstraction(SMSSender{})}
	if got := u.Operation("hi"); got != "[URGENT] sms:hi" {
		t.Fatalf("urgent bridge got %q", got)
	}
}

// Composite 组合模式：将对象组织为树形结构，统一处理单个与组合对象
func TestComposite(t *testing.T) {
	root := NewDirectory("root")
	root.Add(NewFile("a.txt", 10))
	sub := NewDirectory("sub")
	sub.Add(NewFile("b.txt", 5))
	root.Add(sub)
	if root.Size() != 15 {
		t.Fatalf("expected size 15, got %d", root.Size())
	}
	if root.Name() != "root" || sub.Name() != "sub" {
		t.Fatal("directory name wrong")
	}
}

// Facade 外观模式：为子系统提供统一的简化接口
func TestFacade(t *testing.T) {
	f := NewOrderFacade()
	if got := f.PlaceOrder("sku1", 100); got != "dispatched:sku1" {
		t.Fatalf("normal order got %q", got)
	}
	if got := f.PlaceOrder("", 100); got != "out of stock" {
		t.Fatalf("empty sku got %q", got)
	}
	if got := f.PlaceOrder("sku2", 0); got != "payment failed" {
		t.Fatalf("zero amount got %q", got)
	}
}

// Flyweight 享元模式：共享对象以减少内存占用
func TestFlyweight(t *testing.T) {
	f := NewGlyphFactory()
	g1 := f.Get('A')
	g2 := f.Get('A')
	if g1 != g2 {
		t.Fatal("flyweight should return same instance for same key")
	}
	if got := g1.Render("mono"); got != "A@mono" {
		t.Fatalf("glyph render got %q", got)
	}
}