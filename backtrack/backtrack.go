package backtrack

import (
    "awesome/array"
    strings "awesome/string"
)

//
//  SubsetsII
//  @Description: NC27 集合的所有子集(一)
//  @Solution: 1、数组长度 2、遍历的起始值设置 3、回溯与记录
//  @param nums
//  @return [][]int
//
func Subsets(nums []int) [][]int {
    return array.SubsetsII(nums)
}

//
//  BTPermuteUnique
//  @Description: NC121 字符串的排列
//  @param str
//  @return []string
//
func BTPermutation(str string) []string {
    return strings.Permutation(str)
}

//
//  BTPermuteUnique
//  @Description: NC42 有重复项数字的全排列
//  @param num
//  @return [][]int
//
func BTPermuteUnique(num []int) [][]int {
    return PermuteUnique(num)
}

//
//  CombinationSum2
//  @Description: NC46 加起来和为目标值的组合(二)
//  @param num
//  @param target
//  @return [][]int
//
func CombinationSum2(num []int, target int) [][]int {
    return array.CombinationSum2(num, target)
}

//
//  GenerateParenthesis
//  @Description: NC26 括号生成
//  @param n
//  @return []string
//
func GenerateParenthesis(n int) []string {
    return strings.GenerateParenthesis(n)
}
