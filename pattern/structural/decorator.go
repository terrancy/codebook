package structural

// 装饰器模式（Decorator）：动态地给一个对象添加额外职责。就增加功能而言，装饰器比生成子类更灵活。
//
// Go 中通过“包裹同一接口”实现层层包裹，避免继承导致的类爆炸。

// Coffee 抽象组件。
type Coffee interface {
	Cost() int
	Desc() string
}

// SimpleCoffee 具体组件：基础咖啡。
type SimpleCoffee struct{}

func (SimpleCoffee) Cost() int     { return 10 }
func (SimpleCoffee) Desc() string  { return "coffee" }

// 装饰器基类：持有被装饰对象，并实现同一接口。
type coffeeDecorator struct {
	coffee Coffee
}

func (d coffeeDecorator) Cost() int    { return d.coffee.Cost() }
func (d coffeeDecorator) Desc() string { return d.coffee.Desc() }

// Milk 具体装饰器：加奶。
type Milk struct{ coffeeDecorator }

func NewMilk(c Coffee) Milk { return Milk{coffeeDecorator{c}} }

func (m Milk) Cost() int    { return m.coffee.Cost() + 3 }
func (m Milk) Desc() string { return m.coffee.Desc() + "+milk" }

// Sugar 具体装饰器：加糖。
type Sugar struct{ coffeeDecorator }

func NewSugar(c Coffee) Sugar { return Sugar{coffeeDecorator{c}} }

func (s Sugar) Cost() int    { return s.coffee.Cost() + 1 }
func (s Sugar) Desc() string { return s.coffee.Desc() + "+sugar" }
