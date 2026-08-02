package creational

import "testing"

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
