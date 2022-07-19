package trees

//
// DiameterOfBinaryTree
// @Description: 543. 二叉树的直径
// @param root
// @return int
//
func DiameterOfBinaryTree(root *TreeNode) int {
    res := 0
    dfsDiameterOfBinaryTree(root, &res)
    return res
}

func dfsDiameterOfBinaryTree(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    left := 0
    if root.Left != nil {
        left = dfsDiameterOfBinaryTree(root.Left, res) + 1
    }
    right := 0
    if root.Right != nil {
        right = dfsDiameterOfBinaryTree(root.Right, res) + 1
    }
    *res = maxIntDiameterOfBinaryTree(*res, left+right)
    return maxIntDiameterOfBinaryTree(left, right)
}

func maxIntDiameterOfBinaryTree(a, b int) int {
    if a > b {
        return a
    }
    return b
}
