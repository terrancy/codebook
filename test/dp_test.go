package test

import (
    "awesome/dp"
    "fmt"
    "testing"
)

//
//  TestLongCommonSubstring
//  @Description: 最长公共子串
//  @param t
//
func TestLongCommonSubstring(t *testing.T) {
    str1 := "1AB2345CD"
    str2 := "12345EF"
    //res := longCommonSubstring(str1, str2)
    //fmt.Print(res)
    ans := dp.LongCommonSubstring(str1, str2)
    fmt.Println(ans)
}

//
//  TestThrowEggs
//  @Description: NC87 丢棋子问题
//  @Description:
//  @param t
//
func TestThrowEggs(t *testing.T) {
    ans := dp.ThrowEggsII(105, 2)
    fmt.Println(ans)
}

//
//  TestCuttingRopes
//  @Description: JZ14 剪绳子 - 动态规划
//  @param t
//
func TestCuttingRopesDP(t *testing.T) {
    ans := dp.CuttingRopesDP(100)
    fmt.Println(ans)
}
//
//  TestCuttingRopesGreedy
//  @Description: Z14 剪绳子 - 贪心算法
//  @param t
//
func TestCuttingRopesGreedy(t *testing.T) {
    ans := dp.CuttingRosesGreedy(120)
    fmt.Println(ans)
}
