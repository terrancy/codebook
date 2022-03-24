package test

import (
    "awesome/array"
    "fmt"
    "testing"
)

//
//  TestTwoSum
//  @Description:NC61 两数之和
//  @param t
//
func TestTwoSum(t *testing.T) {
    data := []int{20, 70, 110, 150}
    res := array.TwoSum(data, 90)
    fmt.Println(res)
}

//
//  TestTreeSum
//  @Description: NC54 数组中相加和为0的三元组 /  三数之和
//  @param t
//
func TestTreeSum(t *testing.T) {
    data := []int{-10, 0, 10, 20, -10, -40}
    res := array.ThreeSumII(data)
    fmt.Println(res)
}

//
//  TestMaxLengthRepeatLess
//  @Description: NC41 最长无重复子数组
//  @param t
//
func TestMaxLengthRepeatLess(t *testing.T) {
    data := []int{2, 2, 3, 4, 3}
    res := array.MaxLengthRepeatLess(data)
    fmt.Println(res)
}

//
//  TestFindDataTarget
//  @Description: NC29 二维数组中的查找
//  @Solution: 1. 左下角或者右上角 2. if + continue
//  @param t
//
func TestFindDataTarget(t *testing.T) {
    data := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
    target := 3
    res := array.FindTarget(target, data)
    fmt.Println(res)
}

//
//  TestMajorityElement
//  @Description: NC73 数组中出现次数超过一半的数字
//  @param t
//
func TestMajorityElement(t *testing.T) {
    data := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
    res := array.MajorityElement(data)
    fmt.Println(res)
}

//
//  TestReOrderArray
//  @Description: NC77 调整数组顺序使奇数位于偶数前面(一)
//  @param t
//
func TestReOrderArray(t *testing.T) {
    data := []int{1, 3, 5, 6, 7}
    res := array.ReOrderArrayII(data)
    fmt.Println(res)
}

//
//  TestReOrderArrayRepeated
//  @Description: 调整数组顺序使奇数位于偶数前面(二)
//  @param t
//
func TestReOrderArrayRepeated(t *testing.T) {
    data := []int{11, 2, 3, 4, 5, 4, 0}
    res := array.ReOrderArrayRepeatedII(data)
    fmt.Println(res)
}

//
//  TestFindMedianInTwoSortedArrayBS
//  @Description: NC36 在两个长度相等的排序数组中找到上中位数
//  @param t
//
func TestFindMedianInTwoSortedArrayBS(t *testing.T) {
    data1 := []int{1, 2, 3, 4}
    data2 := []int{3, 4, 5, 6}
    res := array.FindMedianInTwoSortedArrayBS(data1, data2)
    fmt.Println(res)
}

//
//  TestInversePairs
//  @Description: JZ51 数组中的逆序对
//  @param t
//
func TestInversePairs(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7, 0}
    res := array.InversePairsQuickSort(data)
    fmt.Println(res)
}

//
//  TestFindRepeatNumber
//  @Description: JZ3 数组中重复的数字
//  @param t
//
func TestFindRepeatNumber(t *testing.T) {
    data := []int{2, 3, 1, 0, 2, 5, 3}
    res := array.FindRepeatNumber(data)
    fmt.Println(res)
}

//
//  TestPrintMinNumber
//  @Description: ERR!!
//  @param t
//
func TestPrintMinNumber(t *testing.T) {
    data := []int{3, 32, 321}
    res := array.PrintMinNumber(data)
    fmt.Println(res)
}

//
//  TestPrintNumbers
//  @Description: JZ17 打印从1到最大的n位数
//  @param t
//
func TestPrintNumbers(t *testing.T) {
    n := 2
    res := array.PrintNumbers(n)
    fmt.Println(res)
}

//
//  TestSubsetsII
//  @Description: NC27 集合的所有子集(一)
//  @Solution: 1、数组长度 2、遍历的起始值设置 3、回溯与记录
//  @param t
//
func TestSubsetsII(t *testing.T) {
    data := []int{1, 2, 3}
    res := array.SubsetsII(data)
    fmt.Println(res)
}

//
//  TestCombinationSum2
//  @Description: NC46 加起来和为目标值的组合(二)
//  @param t
//
func TestCombinationSum2(t *testing.T) {
    data := []int{100, 10, 20, 70, 60, 10, 50}
    target := 80
    res := array.CombinationSum2(data, target)
    fmt.Println(res)
}

//
//  TestSubArrayRanges
//  @Description: 2104. 子数组范围和
//  @param t
//
func TestSubArrayRanges(t *testing.T) {
    data := []int{4, -2, -3, 4, 1}
    res := array.SubArrayRanges(data)
    fmt.Println(res)
}

//
//  TestFirstMissingPositive
//  @Description: 41. 缺失的第一个正数
//  @param t
//
func TestFirstMissingPositive(t *testing.T) {
    data := []int{3, 4, -1, 1}
    res := array.FirstMissingPositive(data)
    fmt.Println(res)
}
