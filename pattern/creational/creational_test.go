package creational

import "testing"

// Singleton 单例模式：确保一个类只有一个实例，并提供全局访问点
func TestSingleton(t *testing.T) {
	a := GetInstance()
	b := GetInstance()
	if a != b {
		t.Fatalf("expected same instance, got %p and %p", a, b)
	}
	if a.Name() != "unique" {
		t.Fatalf("unexpected name: %s", a.Name())
	}
}

// Factory Method 工厂方法模式：定义创建对象的接口，由子类决定实例化哪个类
func TestFactoryMethod(t *testing.T) {
	var c Creator = ConsoleCreator{}
	if got := c.Log("hi"); got != "[console] hi" {
		t.Fatalf("console creator got %q", got)
	}
	c = FileCreator{}
	if got := c.Log("hi"); got != "[file] hi" {
		t.Fatalf("file creator got %q", got)
	}
}

// Abstract Factory 抽象工厂模式：创建产品族的接口，无需指定具体类
func TestAbstractFactory(t *testing.T) {
	wf := NewGUIFactory("windows")
	if wf.CreateButton().Render() != "WinButton" || wf.CreateTextBox().Render() != "WinTextBox" {
		t.Fatal("windows factory mismatch")
	}
	mf := NewGUIFactory("mac")
	if mf.CreateButton().Render() != "MacButton" || mf.CreateTextBox().Render() != "MacTextBox" {
		t.Fatal("mac factory mismatch")
	}
}

// Builder 建造者模式：将复杂对象的构建与其表示分离
func TestBuilder(t *testing.T) {
	pc := NewComputerBuilder().
		SetCPU("i9").
		SetRAM("32G").
		SetStorage("1TB").
		SetGPU("RTX4090").
		Build()
	if pc.CPU != "i9" || pc.RAM != "32G" || pc.Storage != "1TB" || pc.GPU != "RTX4090" {
		t.Fatalf("builder result wrong: %+v", pc)
	}
}

// Prototype 原型模式：通过克隆现有对象创建新对象
func TestPrototype(t *testing.T) {
	orig := &Resume{Name: "tom", Age: 30, Skills: []string{"go", "java"}}
	clone := orig.Clone().(*Resume)
	clone.Skills[0] = "rust" // 修改副本不应影响原对象
	if orig.Skills[0] != "go" {
		t.Fatalf("prototype deep copy failed, orig.Skills=%v", orig.Skills)
	}
	if clone.Skills[0] != "rust" || clone.Name != "tom" {
		t.Fatalf("clone wrong: %+v", clone)
	}
}