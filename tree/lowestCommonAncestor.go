package trees

//
//  lowestCommonAncestor
//  @Description: NC191 二叉搜索树的最近公共祖先
//  @param root *TreeNode
//  @param p int
//  @param q int
//  @return int
//
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	node := dspLowestCommonAncestor(root, p, q)
	return node.Val
}

//
//  dspLowestCommonAncestor
//  @Description: 递归查找最近公共祖先
//  @param root
//  @param p
//  @param q
//  @return *TreeNode
//
func dspLowestCommonAncestor(root *TreeNode, p int, q int) *TreeNode {
	if root == nil || root.Val == p || root.Val == q {
		return root
	}

	left := dspLowestCommonAncestor(root.Left, p, q)
	right := dspLowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
