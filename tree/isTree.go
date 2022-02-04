package trees

import (
    "math"
)

//
//  IsSymmetric
//  @Description: JZ28 对称的二叉树
//  @param root
//  @return bool
//
func IsSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    queue := []*TreeNode{root}
    const INF = math.MinInt32
    dummy := &TreeNode{Val: INF}
    for len(queue) > 0 {
        data := make([]int, 0)
        for size := len(queue); size > 0; size-- {
            node := queue[0]
            queue = queue[1:]
            data = append(data, node.Val)
            if node.Val == INF {
                continue
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            } else {
                queue = append(queue, dummy)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            } else {
                queue = append(queue, dummy)
            }
        }
        if checkIsSymmetric(data) == false {
            return false
        }
    }
    return true
}

//
//  checkIsSymmetric
//  @Description: 判断数组是否是对称的
//  @param data
//  @return bool
//
func checkIsSymmetric(data []int) bool {
    n := len(data)
    for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
        if data[l] != data[r] {
            return false
        }
    }
    return true
}

//
//  IsSymmetricII
//  @Description: 先根遍历
//  @param root
//  @return bool
//
func IsSymmetricII(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetricPreOrder(root.Left, root.Right)
    //return isSymmetricPreOrder(root, root)
}

//
//  isSymmetricPreOrder
//  @Description: 先根遍历
//  @param root1
//  @param root2
//  @return bool
//
func isSymmetricPreOrder(root1, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil || root2 == nil {
        return false
    }
    if root1.Val != root2.Val {
        return false
    }
    return isSymmetricPreOrder(root1.Left, root2.Right) && isSymmetricPreOrder(root1.Right, root2.Left)
}
