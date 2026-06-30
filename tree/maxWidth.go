package trees

import (
	"terrancy/awesome"
)

// WidthOfBinaryTree
// @Title: LC662.二叉树最大宽度
// @Description: 给你一棵二叉树的根节点root，返回树的最大宽度
// @Description: 树的最大宽度是所有层中最大的宽度，每一层的宽度被定义为该层最左和最右的非空节点之间的长度
// @Link: https://leetcode.cn/problems/maximum-width-of-binary-tree/
// @param root
// @return int
func WidthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	// 初始化
	root.Val = 1
	queue = queuePush(queue, root)
	node := &TreeNode{}
	maxWidth := 1
	for size := 0; len(queue) > 0; {
		size = len(queue)
		// 比较大小
		maxWidth = awesome.MaxInt(maxWidth, queue[len(queue)-1].Val-queue[0].Val+1)
		for size > 0 {
			node, queue = queueShift(queue)
			if node.Left != nil {
				// 设置左
				node.Left.Val = 2 * node.Val
				queue = queuePush(queue, node.Left)
			}
			if node.Right != nil {
				// 设置右
				node.Right.Val = 2*node.Val + 1
				queue = queuePush(queue, node.Right)
			}
			size--
		}
	}

	return maxWidth
}