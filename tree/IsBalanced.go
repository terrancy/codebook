package trees

//
//  IsBalanced
//  @Description: NC62 判断是不是平衡二叉树
//  @param root
//  @return bool
//
func IsBalanced(root *TreeNode) bool {
    res := postorderIsBalanced(root)
    if res == -1 {
        return false
    }
    return true
}

func postorderIsBalanced(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := postorderIsBalanced(root.Left)
    if left == -1 {
        return -1
    }
    right := postorderIsBalanced(root.Right)
    if right == -1 {
        return -1
    }
    if absMinusIntIsBalanced(left, right) > 1 {
        return -1
    }
    return maxIntIsBalanced(left, right) + 1
}

func maxIntIsBalanced(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func absMinusIntIsBalanced(a, b int) int {
    if b > a {
        a, b = b, a
    }
    return a - b
}
