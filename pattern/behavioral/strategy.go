package behavioral

// 策略模式（Strategy）：定义一系列算法，把它们封装起来，并且使它们可以互相替换。
// 策略让算法独立于使用它的客户而变化（配合 OCP）。
//
// 对应七大原则中的开闭原则：新增算法只需新增策略类型，不改调用方。

// Strategy 出行策略接口。
type Strategy interface {
	Route(from, to string) string
}

type WalkStrategy struct{}

func (WalkStrategy) Route(from, to string) string { return "walk: " + from + " -> " + to }

type DriveStrategy struct{}

func (DriveStrategy) Route(from, to string) string { return "drive: " + from + " -> " + to }

// Navigator 上下文：持有一个策略，运行时可替换。
type Navigator struct {
	strategy Strategy
}

func (n *Navigator) SetStrategy(s Strategy) { n.strategy = s }

func (n *Navigator) Go(from, to string) string {
	if n.strategy == nil {
		return "no strategy set"
	}
	return n.strategy.Route(from, to)
}
