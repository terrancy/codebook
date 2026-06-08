package trees

import "terrancy/awesome"

// MinDepth
// @Description: NC234 二叉树的最小深度
// @param root
// @return int
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		level++
		for size := len(queue); size > 0; size-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return level
			}
		}
	}
	return level
}

func MinDepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MinDepthDfs(root.Left)
	right := MinDepthDfs(root.Right)

	return awesome.MaxInt(left, right) + 1
}
