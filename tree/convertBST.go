package trees

//
//  convertBST
//  @Description: NC215 将二叉搜索树改为累加树
//  @Description: 给定一个二叉搜索树，树上的节点各不相同，请你将其修改为累加树，使每个节点的值变成原树中更大节点之和
//  @Link: https://www.nowcoder.com/practice/19bff16ca0d64d6da38ed24fd2903ce9?tpId=117&&tqId=39419&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking
//  @param root
//  @return *TreeNode
//
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dspConvertBST(root, &sum)
	return root
}

func dspConvertBST(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return root
	}
	dspConvertBST(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	dspConvertBST(root.Left, sum)

	return root
}
