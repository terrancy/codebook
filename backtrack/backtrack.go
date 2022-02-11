package backtrack

import (
    "awesome/array"
    strings "awesome/string"
)

//
//  BTPermuteUnique
//  @Description: 字符串全排列
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
