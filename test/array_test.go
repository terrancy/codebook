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