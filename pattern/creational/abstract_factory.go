package creational

// 抽象工厂模式（Abstract Factory）：创建一系列相关或相互依赖的产品族，而不指定具体类。
//
// 与工厂方法的区别：工厂方法只产一个产品；抽象工厂产“一族”配套产品。

// 抽象产品：按钮与文本框（同一 GUI 族里配套使用，风格必须一致）。
type Button interface{ Render() string }
type TextBox interface{ Render() string }

// 具体产品：Windows 族
type WinButton struct{}

func (WinButton) Render() string { return "WinButton" }

type WinTextBox struct{}

func (WinTextBox) Render() string { return "WinTextBox" }

// 具体产品：Mac 族
type MacButton struct{}

func (MacButton) Render() string { return "MacButton" }

type MacTextBox struct{}

func (MacTextBox) Render() string { return "MacTextBox" }

// GUIFactory 抽象工厂：负责产出一整套风格一致的产品。
type GUIFactory interface {
	CreateButton() Button
	CreateTextBox() TextBox
}

// WinFactory 生产 Windows 风格整套组件
type WinFactory struct{}

func (WinFactory) CreateButton() Button   { return WinButton{} }
func (WinFactory) CreateTextBox() TextBox { return WinTextBox{} }

// MacFactory 生产 Mac 风格整套组件
type MacFactory struct{}

func (MacFactory) CreateButton() Button   { return MacButton{} }
func (MacFactory) CreateTextBox() TextBox { return MacTextBox{} }

// NewGUIFactory 根据操作系统返回对应工厂（简单工厂辅助函数）。
func NewGUIFactory(os string) GUIFactory {
	if os == "mac" {
		return MacFactory{}
	}
	return WinFactory{}
}
