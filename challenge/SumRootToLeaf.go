package challenge

import trees "awesome/tree"

//
//  SumRootToLeaf
//  @Description: 1022. 从根到叶的二进制数之和
//  @param root
//  @return int
//
func SumRootToLeaf(root *trees.TreeNode) int {
    return DFSSumRootToLeaf(root, 0)
}

func DFSSumRootToLeaf(root *trees.TreeNode, data int) int {
    if root == nil {
        return 0
    }
    data = data*2 + root.Val
    if root.Left == nil && root.Right == nil {
        return data
    }
    return DFSSumRootToLeaf(root.Left, data) + DFSSumRootToLeaf(root.Right, data)
}
