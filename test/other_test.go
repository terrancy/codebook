package test

import (
    "awesome/others"
    "fmt"
    "testing"
)

//
//  TestFoundOnceNumber
//  @Description: NC156 数组中只出现一次的数（其它数出现k次）
//  @param t
//
func TestFoundOnceNumber(t *testing.T) {
    data := []int{5, -4, 1, 1, 5, 1, 5}
    res := others.FoundOnceNumber(data, 3)
    fmt.Println(res)
}

//
//  TestFindNumsAppearOnce
//  @Description: NC75 数组中只出现一次的两个数字
//  @param t
//
func TestFindNumsAppearOnce(t *testing.T) {
    data := []int{1, 2, 3, 3, 2, 9}
    res := others.FindNumsAppearOnce(data)
    fmt.Println(res)
}

//
//  TestScaleConvert
//  @Description: 进制转换
//  @param t
//
func TestScaleConvert(t *testing.T) {
    res := others.ScaleConvert(-4, 3)
    fmt.Println(res)
}

//
//  TestTopKStrings
//  @Description: NC97 字符串出现次数的TopK问题
//  @param t
//
func TestTopKStrings(t *testing.T) {
    data := []string{"123", "123", "231", "32"}
    res := others.TopKStrings(data, 2)
    fmt.Println(res)
}

//
//  TestSqrt
//  @Description: NC32 求平方根
//  @param t
//
func TestSqrt(t *testing.T) {
    res := others.Sqrt(5)
    fmt.Println(res)
}
