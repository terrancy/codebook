package main

import "math"

func isSearchBinaryTree(root *TreeNode) bool {
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
