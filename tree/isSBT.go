package trees

import "math"

func IsSearchBinaryTree(root *TreeNode) bool {
    min := math.MinInt32
    return isSBT(root, &min)
}

//
//  isSBT
//  @Description: 判断是否是搜索二叉树
//  @Solution: 先根遍历是升序的,设置最小值,发现有更小的值就有问题
//  @param root
//  @return bool
//
func isSBT(root *TreeNode, min *int) bool {
    if root == nil {
        return true
    }
    flag := isSBT(root.Left, min)
    if flag == false {
        return false
    }
    if *min > root.Val {
        return false
    }
    *min = root.Val
    return isSBT(root.Right, min)
}

//
//  KthNode
//  @Description: 二叉搜索树第K小(大)的值
//  @param root
//  @param k
//  @return bool
//
var res, cnt = -1, 0

func KthNode(root *TreeNode, k int) int {
    // 第K大
    inorderKthMax(root, k)
    // 第K小
    //inorderKthMin(root, k)
    return res
}

//
//  inorderKthMax
//  @Description: 第K大
//  @param root
//  @param k
//
func inorderKthMax(root *TreeNode, k int) {
    if root == nil {
        return
    }
    inorderKthMax(root.Right, k)
    cnt++
    if cnt == k {
        res = root.Val
    }
    if cnt >= k {
        return
    }
    inorderKthMax(root.Left, k)
}

//
//  inorderKthMin
//  @Description: 第K小
//  @param root
//  @param k
//
func inorderKthMin(root *TreeNode, k int) {
    if root == nil {
        return
    }
    inorderKthMin(root.Left, k)
    cnt++
    if cnt == k {
        res = root.Val
    }
    if cnt >= k {
        return
    }
    inorderKthMin(root.Right, k)
}
