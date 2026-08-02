package creational

// 建造者模式（Builder）：将一个复杂对象的构建与其表示分离，使得同样的构建过程可以创建不同的表示。
//
// 适合：构造参数很多（telescoping constructor 反模式）、且部分可选的场景。

// Computer 最终产品：字段多，用 Builder 分步设置。
type Computer struct {
	CPU     string
	RAM     string
	Storage string
	GPU     string
}

// ComputerBuilder 建造者接口：每一步返回自身以便链式调用。
type ComputerBuilder interface {
	SetCPU(string) ComputerBuilder
	SetRAM(string) ComputerBuilder
	SetStorage(string) ComputerBuilder
	SetGPU(string) ComputerBuilder
	Build() Computer
}

type computerBuilder struct {
	c Computer
}

// NewComputerBuilder 返回新的建造者（Director 可用它来规范步骤）。
func NewComputerBuilder() ComputerBuilder {
	return &computerBuilder{}
}

func (b *computerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.c.CPU = cpu
	return b
}

func (b *computerBuilder) SetRAM(ram string) ComputerBuilder {
	b.c.RAM = ram
	return b
}

func (b *computerBuilder) SetStorage(s string) ComputerBuilder {
	b.c.Storage = s
	return b
}

func (b *computerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.c.GPU = gpu
	return b
}

func (b *computerBuilder) Build() Computer { return b.c }
