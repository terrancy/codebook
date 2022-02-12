package array

import "sort"

//
//  subsets
//  @Description:
//  @param nums
//  @return [][]int
//
func Subsets(nums []int) [][]int {
    n := len(nums)
    if n == 0 {
        return [][]int{[]int{}}
    }
    item := nums[n-1]
    res := Subsets(nums[0 : n-1])
    
    ln := len(res)
    for i := 0; i < ln; i++ {
        tmp := make([]int, len(res[i]))
        copy(tmp, res[i])
        tmp = append(tmp, item)
        res = append(res, tmp)
    }
    
    return res
}

//
//  SubsetsII
//  @Description: NC27 集合的所有子集(一)
//  @Solution: 1、数组长度 2、遍历的起始值设置 3、回溯与记录
//  @param nums
//  @return [][]int
//
func SubsetsII(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    res := make([][]int, 0)
    for i := 0; i <= n; i++ {
        BSSubsets(nums, []int{}, &res, i, 0)
    }
    return res
}

//
//  BSSubsets
//  @Description:
//  @param nums
//  @param stack
//  @param res
//  @param k
//  @param start
//
func BSSubsets(nums []int, stack []int, res *[][]int, k int, start int) {
    // 结束条件
    if k == len(stack) {
        track := make([]int, len(stack))
        copy(track, stack)
        *res = append(*res, track)
    }
    
    // 遍历、选择、递归、撤销
    for i := start; i < len(nums); i++ {
        stack = append(stack, nums[i])
        BSSubsets(nums, stack, res, k, i+1)
        stack = stack[:len(stack)-1]
    }
}
