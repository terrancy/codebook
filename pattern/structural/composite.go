package structural

// 组合模式（Composite）：将对象组合成树形结构以表示“部分-整体”的层次结构，
// 使得用户对单个对象和组合对象的使用具有一致性。
//
// 典型场景：文件系统、UI 组件树、组织架构。

// Component 组件接口：叶子与容器共用，统一对待。
type Component interface {
	Name() string
	Size() int
}

// File 叶子节点：文件。
type File struct {
	name string
	size int
}

func NewFile(name string, size int) File { return File{name: name, size: size} }
func (f File) Name() string              { return f.name }
func (f File) Size() int                 { return f.size }

// Directory 容器节点：目录，可包含子组件。
type Directory struct {
	name     string
	children []Component
}

func NewDirectory(name string) *Directory { return &Directory{name: name} }

func (d *Directory) Name() string { return d.name }

func (d *Directory) Add(c Component) { d.children = append(d.children, c) }

// Size 递归统计：目录大小 = 所有子组件大小之和。
func (d *Directory) Size() int {
	total := 0
	for _, c := range d.children {
		total += c.Size()
	}
	return total
}
