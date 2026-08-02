package creational

// 原型模式（Prototype）：用原型实例指定创建对象的种类，并通过拷贝这些原型创建新的对象。
//
// Go 没有语言级的 clone，通常给类型加一个 Clone 方法返回深拷贝。
// 适合：创建成本高的对象（如含大量预计算字段）想被快速复制时。

// Prototype 原型接口：可自我克隆。
type Prototype interface {
	Clone() Prototype
}

// Resume 具体原型：简历，含可变字段。
type Resume struct {
	Name    string
	Age     int
	Skills  []string // 切片需深拷贝，否则多个副本共享底层数组
}

// Clone 返回深拷贝，避免副本间相互影响。
func (r *Resume) Clone() Prototype {
	skills := make([]string, len(r.Skills))
	copy(skills, r.Skills)
	return &Resume{Name: r.Name, Age: r.Age, Skills: skills}
}
