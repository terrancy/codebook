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

//
//  TestEditDistance
//  @Description: NC196 编辑距离(一)
//  @param t
//
func TestEditDistance(t *testing.T) {
    str1, str2 := "nowcoder", "new"
    ans := dp.EditDistance(str1, str2)
    fmt.Println(ans)
}

//
//  TestEditCost
//  @Description:
//  @param t
//
func TestEditCost(t *testing.T) {
    str1, str2 := "abc", "adc"
    insertCost, deleteCost, replaceCost := 5, 3, 2
    ans := dp.EditCost(str1, str2, insertCost, deleteCost, replaceCost)
    fmt.Println(ans)
}
