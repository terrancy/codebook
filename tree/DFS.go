package trees

//
// ThreeOrders
// @Description: NC45 实现二叉树先序，中序和后序遍历
// @param root
// @return [][]int
//
func ThreeOrders(root *TreeNode) [][]int {
    res := make([][]int, 3)
    for i := 0; i < 3; i++ {
        res[i] = make([]int, 0)
    }
    dfs(root, res)
    return res
}

// 递归
func dfs(root *TreeNode, res [][]int) {
    if root == nil {
        return
    }
    res[0] = append(res[0], root.Val)
    dfs(root.Left, res)
    res[1] = append(res[1], root.Val)
    dfs(root.Right, res)
    res[2] = append(res[2], root.Val)
}

// 非递归/迭代

//
// inorder
// @Description: 中根遍历 - 迭代
// @Solution: 根不空时循环左入栈,栈不空时根出栈且右入栈
// @param root
// @param res
//
func inorder(root *TreeNode, res [][]int) {
    if root == nil {
        return
    }
    // 根入栈
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) > 0 {
        for root != nil {
            // 循环左入栈
            stack = append(stack, root)
            root = root.Left
        }
        
        if len(stack) > 0 {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[0] = append(res[0], node.Val)
            // 右入栈
            root = node.Right
        }
    }
}

//
// preorder
// @Description: 先根遍历 - 迭代
// @Solution: 根入栈,出栈,右左入栈
// @param root
// @param res
//
func preorder(root *TreeNode, res [][]int) {
    if root == nil {
        return
    }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res[1] = append(res[1], node.Val)
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
}

//
// postorder
// @Description: 后根遍历 - 迭代
// @Solution: 根入栈,出栈,左右入栈,翻转
// @param root
// @param res
//
func postorder(root *TreeNode, res [][]int) {
    if root == nil {
        return
    }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res[1] = append(res[1], node.Val)
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    // 翻转操作
    if len(res[1]) > 0 {
        n := len(res[1])
        for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
            res[1][l], res[1][r] = res[1][r], res[1][l]
        }
    }
}
