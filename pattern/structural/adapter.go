package structural

// 适配器模式（Adapter）：将一个类的接口转换成客户希望的另一个接口，使原本因接口不兼容而不能一起工作的类可以协同。
//
// 常见场景：三方库接口与自有接口不匹配，加一层适配而非改两端。

// Target 客户期望的接口。
type Target interface {
	Request() string
}

// Adaptee 已存在的、接口不兼容的类。
type Adaptee struct{}

func (Adaptee) SpecificRequest() string { return "adaptee specific" }

// Adapter 把 Adaptee 适配成 Target。
type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(a Adaptee) Adapter { return Adapter{adaptee: a} }

func (a Adapter) Request() string {
	// 转换/转发调用
	return "adapter:" + a.adaptee.SpecificRequest()
}
