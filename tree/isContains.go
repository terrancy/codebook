package trees

// IsContains
// @Title: LC572.另一棵树的子树
// @Description: 判断两棵树是否一致 递归,先根遍历
// @Link: https://leetcode.cn/problems/subtree-of-another-tree/
// @param root1
// @param root2
// @return bool
func IsContains(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return false
	}
	return IsSame(root1, root2) || IsContains(root1.Left, root2) || IsContains(root1.Right, root2)
}

// IsSame
// @Title: LC100.相同的树
// @Description: 判断两棵树是否一致
// @Link: https://leetcode.cn/problems/same-tree/
// @param root1
// @param root2
// @return bool
func IsSame(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && IsSame(root1.Right, root2.Right) && IsSame(root1.Left, root2.Left)
}