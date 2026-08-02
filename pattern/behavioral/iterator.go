package behavioral

// 迭代器模式（Iterator）：提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示。
//
// Go 里语言内建就有 range + 切片，但有时需要自定义遍历顺序（逆序、树遍历等）。

// Iterator 迭代器接口。
type Iterator interface {
	HasNext() bool
	Next() string
}

// NameCollection 聚合对象：持有数据。
type NameCollection struct {
	items []string
}

func NewNameCollection(items []string) *NameCollection {
	return &NameCollection{items: items}
}

// Iterator 返回一个逆序迭代器（自定义遍历顺序的示例）。
func (c *NameCollection) Iterator() Iterator {
	return &reverseIterator{items: c.items, index: len(c.items) - 1}
}

type reverseIterator struct {
	items []string
	index int
}

func (it *reverseIterator) HasNext() bool { return it.index >= 0 }
func (it *reverseIterator) Next() string {
	v := it.items[it.index]
	it.index--
	return v
}
