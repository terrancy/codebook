package test

import (
    "awesome/array"
    "fmt"
    "testing"
)

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
//  @Description: NC54 数组中相加和为0的三元组
//  @param t
//
func TestTreeSum(t *testing.T) {
    data := []int{-10, 0, 10, 20, -10, -40}
    res := array.ThreeSumII(data)
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
