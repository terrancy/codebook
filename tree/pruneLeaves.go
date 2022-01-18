package trees

//
//  pruneLeaves
//  @Description: NC169 修剪叶子
//  @link: NC169 https://www.nowcoder.com/practice/39b559fb84864bde93eccd6e87312650?tpId=117&&tqId=39314&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking
//  @param root
//  @return *TreeNode
//
func pruneLeaves(root *TreeNode) *TreeNode {
	// 空节点 或者 叶子节点
	if root == nil || (root.Left == nil && root.Right == nil) {
		return nil
	}
	// 左孩子是叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return nil
	}
	// 右孩子是叶子节点
	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil {
		return nil
	}

	root.Left = pruneLeaves(root.Left)
	root.Right = pruneLeaves(root.Right)
	return root
}

//
//  pruneTree
//  @Description: 814. 二叉树剪枝
//  @Link: https://leetcode-cn.com/problems/binary-tree-pruning/
//  @param root
//  @return *TreeNode
//
func pruneTree(root *TreeNode) *TreeNode {
	return root
}
