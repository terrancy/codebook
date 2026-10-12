package trees

// BSTIterator
// @Title: LC173.二叉搜索树迭代器
// @Description: 实现一个 BST 迭代器，按中序遍历顺序依次返回节点值
//               hasNext() 判断是否还有下一个节点
//               next() 返回下一个节点值
// @Link: https://leetcode.cn/problems/binary-search-tree-iterator/
//
//	Example:
//		root := [7, 3, 15, null, null, 9, 20]
//		iter := BSTIteratorConstructor(root)
//		iter.Next()  // 3
//		iter.Next()  // 7
//		iter.HasNext() // true
//		iter.Next()  // 9
//		iter.HasNext() // true
//		iter.Next()  // 15
//		iter.HasNext() // true
//		iter.Next()  // 20
//		iter.HasNext() // false

type BSTIterator struct {
	stack []*TreeNode
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{stack: make([]*TreeNode, 0)}
	iter.pushLeft(root)
	return iter
}

func (it *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

func (it *BSTIterator) Next() int {
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	if node.Right != nil {
		it.pushLeft(node.Right)
	}
	return node.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}