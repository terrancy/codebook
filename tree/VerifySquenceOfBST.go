package main

//
//  VerifySequenceOfBST
//  @Description: JZ33 二叉搜索树的后序遍历序列
//  @param sequence
//  @return bool
//
func VerifySequenceOfBST(sequence []int) bool {
    if sequence == nil {
        return false
    }
    return verifyIfBST(sequence)
}

//
//  verifyIfBST
//  @Description: 递归判断是否满足二叉搜索树
//  @Solution:
//  @param sequence
//  @return bool
//
func verifyIfBST(sequence []int) bool {
    n := len(sequence)
    if n < 2 {
        return true
    }
    mid := sequence[n-1]
    left := make([]int, 0)
    right := make([]int, 0)
    flag := -1
    for i := 0; i < n-1; i++ {
        if sequence[i] < mid {
            if flag == -1 {
                left = append(left, sequence[i])
            } else {
                return false
            }
        } else {
            if flag == -1 {
                flag = 0
            }
            right = append(right, sequence[i])
        }
    }
    
    return verifyIfBST(left) && verifyIfBST(right)
}
