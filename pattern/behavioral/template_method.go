package behavioral

// 模板方法模式（Template Method）：定义一个操作中的算法骨架，将一些步骤延迟到子类中实现。
// 子类可以不改变算法结构即可重定义该算法的某些特定步骤。
//
// Go 没有继承，用“嵌入 + 钩子接口”模拟：基类结构体持有“可变的步骤实现”。

// BrewSteps 可变步骤（由具体子类提供）。
type BrewSteps interface {
	AddMaterial() string
	Brew() string
}

// Beverage 模板：固定流程，调用可变步骤。
type Beverage struct {
	steps BrewSteps
}

func (b *Beverage) SetSteps(s BrewSteps) { b.steps = s }

// Make 算法骨架（模板方法）：固定流程不可被重写。
func (b *Beverage) Make() string {
	if b.steps == nil {
		return "no steps"
	}
	return "boil water -> " + b.steps.AddMaterial() + " -> " + b.steps.Brew()
}

// Tea 具体实现：茶的步骤。
type Tea struct{}

func (Tea) AddMaterial() string { return "add tea leaf" }
func (Tea) Brew() string        { return "steep 3 min" }

// CoffeeRecipe 具体实现：咖啡的步骤。
type CoffeeRecipe struct{}

func (CoffeeRecipe) AddMaterial() string { return "add coffee powder" }
func (CoffeeRecipe) Brew() string        { return "drip brew" }
