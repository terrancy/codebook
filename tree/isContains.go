package trees

//
//  isContains
//  @Description: 判断两棵树是否一致 递归,先根遍历
//  @param root1
//  @param root2
//  @return bool
//
func IsContains(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return false
	}
	return IsSame(root1, root2) || IsContains(root1.Left, root2) || IsContains(root1.Right, root2)
}

//
//  isSame
//  @Description: 判断两棵树是否一致
//  @param root1
//  @param root2
//  @return bool
//
func IsSame(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && IsSame(root1.Right, root2.Right) && IsSame(root1.Left, root2.Left)
}
