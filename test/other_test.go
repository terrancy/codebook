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
func TestSQRT(t *testing.T) {
    res := others.SQRT(5)
    fmt.Println(res)
}

//
//  TestNumIslands
//  @Description: NC109 岛屿数量
//  @param t
//
func TestNumIslands(t *testing.T) {
    grid := [][]byte{
        {'1', '1', '0', '0', '0'},
        {'0', '1', '0', '1', '1'},
        {'0', '0', '0', '1', '1'},
        {'0', '0', '0', '0', '0'},
        {'0', '0', '1', '1', '1'},
    }
    cnt := others.NumIslands(grid)
    fmt.Println(cnt)
}

//
//  TestIsPokerContinuous
//  @Description: NC63 扑克牌顺子
//  @param t
//
func TestIsPokerContinuous(t *testing.T) {
    data := []int{6, 0, 2, 0, 4}
    res := others.IsPokerContinuous(data)
    fmt.Println(res)
}

//
//  TestIsPalindrome
//  @Description: NC56 回文数字
//  @param t
//
func TestIsPalindrome(t *testing.T) {
    num := 121
    res := others.IsPalindromeII(num)
    fmt.Println(res)
}
