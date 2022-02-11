package array

import (
    "sort"
)

//
//  combinationSum2 NC46 加起来和为目标值的组合(二)
//  @Description: 给出一组候选数 c 和一个目标数 t ，找出候选数中起来和等于 t 的所有组合
//  @param num
//  @param target
//  @return [][]int
//
func CombinationSum2(num []int, target int) [][]int {
    sort.Ints(num)
    temp := make([]int, 0)
    res := make([][]int, 0)
    backtrace(num, target, &temp, &res, 0)
    return res
}

//
//  backtrace
//  @Description: 回溯
//  @param num
//  @param target
//  @param temp
//  @param res
//  @param start
//
func backtrace(num []int, target int, temp *[]int, res *[][]int, start int) {
    if target == 0 {
        t := make([]int, len(*temp))
        copy(t, *temp)
        *res = append(*res, t)
        return
    }
    n := len(num)
    if start == n {
        return
    }
    
    for i := start; i < n; i++ {
        if num[i] > target {
            return
        }
        // 剪枝
        if i > start && num[i-1] == num[i] {
            continue
        }
        *temp = append(*temp, num[i])
        backtrace(num, target-num[i], temp, res, i+1)
        // 回溯,不选num[i]的情况
        *temp = (*temp)[:len(*temp)-1]
    }
    return
}