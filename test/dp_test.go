package test

import (
    "awesome/dp"
    "fmt"
    "testing"
)

//
//  TestFib
//  @Description: Fibonacci数列
//  @param t
//
func TestFib(t *testing.T) {
    res := dp.Fib(10)
    fmt.Println(res)
}

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
    res := dp.LongCommonSubstring(str1, str2)
    fmt.Println(res)
}

//
//  TestThrowEggs
//  @Description: NC87 丢棋子问题
//  @Description:
//  @param t
//
func TestThrowEggs(t *testing.T) {
    res := dp.ThrowEggsII(105, 2)
    fmt.Println(res)
}

//
//  TestCuttingRopes
//  @Description: JZ14 剪绳子 - 动态规划
//  @param t
//
func TestCuttingRopesDP(t *testing.T) {
    res := dp.CuttingRopesDP(100)
    fmt.Println(res)
}

//
//  TestCuttingRopesGreedy
//  @Description: Z14 剪绳子 - 贪心算法
//  @param t
//
func TestCuttingRopesGreedy(t *testing.T) {
    res := dp.CuttingRosesGreedy(120)
    fmt.Println(res)
}

//
//  TestEditDistance
//  @Description: NC196 编辑距离(一)
//  @param t
//
func TestEditDistance(t *testing.T) {
    str1, str2 := "nowcoder", "new"
    res := dp.EditDistance(str1, str2)
    fmt.Println(res)
}

//
//  TestEditCost
//  @Description: NC35 编辑距离(二)
//  @param t
//
func TestEditCost(t *testing.T) {
    str1, str2 := "abc", "adc"
    insertCost, deleteCost, replaceCost := 5, 3, 2
    res := dp.EditCost(str1, str2, insertCost, deleteCost, replaceCost)
    fmt.Println(res)
}

//
//  TestCoinChange
//  @Description:
//  @param t
//
func TestCoinChange(t *testing.T) {
    coins := []int{1, 2, 5}
    amount := 11
    res := dp.CoinChange(coins, amount)
    fmt.Println(res)
}
