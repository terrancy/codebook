package trees

import "awesome"

//
// MaxPathSum
// @Description: NC6 二叉树中的最大路径和
// @param root
// @return int
//
var maxPathSumRes = -1001

func MaxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    postorderMaxPathSum(root)
    return maxPathSumRes
}

func postorderMaxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := postorderMaxPathSum(root.Left)
    right := postorderMaxPathSum(root.Right)
    maxPathSumRes = awesome.MaxInt(maxPathSumRes, left+right+root.Val)
    return awesome.MaxInt(awesome.MaxInt(left, right)+root.Val, 0)
}
