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
//  TestMaxSubArrayII
//  @Description: JZ85 连续子数组的最大和(二)
//  @param t
//
func TestMaxSubArrayII(t *testing.T) {
    data := []int{1, -2, 3, 10, -4, 7, 2, -5}
    res := dp.MaxSubArrayII(data)
    fmt.Println(res)
}

//
//  TestMaxProfit
//  @Description: JZ63 买卖股票的最好时机(一)
//  @param t
//
func TestMaxProfit(t *testing.T) {
    data := []int{8, 9, 2, 5, 4, 7, 1}
    res := dp.MaxProfit(data)
    fmt.Println(res)
}

//
//  TestMaxProfit
//  @Description: JZ63 买卖股票的最好时机(一)
//  @param t
//
func TestMaxProfitII(t *testing.T) {
    data := []int{8, 5, 1, 7, 10, 12}
    res := dp.MaxProfitII(data)
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
//  @Description: LC322. 零钱兑换
//  @param t
//
func TestCoinChange(t *testing.T) {
    coins := []int{1, 2, 5}
    amount := 11
    res := dp.CoinChange(coins, amount)
    fmt.Println(res)
}

//
//  TestLongestPalindromeII
//  @Description: NC17 最长回文子串
//  @param t
//
func TestLongestPalindromeII(t *testing.T) {
    str := "ababc"
    res := dp.LongestPalindrome(str)
    fmt.Println(res)
}

//
//  TestLongestPalindromeSubsequence
//  @Description: 最长回文子序列
//  @param t
//
func TestLongestPalindromeSubsequence(t *testing.T) {
    str := "bbbab"
    res := dp.LongestPalindromeSubsequence(str)
    fmt.Println(res)
}

//
//  TestLongestValidParentheses
//  @Description: 32. 最长有效括号
//  @param t
//
func TestLongestValidParentheses(t *testing.T) {
    str := ")()())"
    res := dp.LongestValidParentheses(str)
    fmt.Println(res)
}
