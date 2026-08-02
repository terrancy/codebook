package behavioral

// 访问者模式（Visitor）：表示一个作用于某对象结构中的各元素的操作，
// 它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
//
// 典型场景：AST 遍历、报表导出（对同一批对象做多种不同处理）。

// Element 元素接口：接受访问者。
type Element interface {
	Accept(v Visitor) string
}

// Visitor 访问者接口：为每种元素定义访问操作。
type Visitor interface {
	VisitFile(*FileElement) string
	VisitDir(*DirElement) string
}

// FileElement 具体元素：文件。
type FileElement struct{ Name string }

func (e *FileElement) Accept(v Visitor) string { return v.VisitFile(e) }

// DirElement 具体元素：目录。
type DirElement struct{ Name string }

func (e *DirElement) Accept(v Visitor) string { return v.VisitDir(e) }

// SizeVisitor 具体访问者：统计大小（文件=1，目录=5 仅为示例）。
type SizeVisitor struct{}

func (SizeVisitor) VisitFile(f *FileElement) string { return f.Name + ":1" }
func (SizeVisitor) VisitDir(d *DirElement) string   { return d.Name + ":5" }

// PathVisitor 具体访问者：打印路径（另一种不修改元素的新操作）。
type PathVisitor struct{}

func (PathVisitor) VisitFile(f *FileElement) string { return "/files/" + f.Name }
func (PathVisitor) VisitDir(d *DirElement) string   { return "/dirs/" + d.Name }
