package dp

import (
    "awesome"
    "math"
)

// 背包问题相关
// @Link: https://leetcode-cn.com/problems/coin-change-2/solution/gong-shui-san-xie-xiang-jie-wan-quan-bei-6hxv/

//
//  coinChange
//  @Description: LC322. 零钱兑换
//  @Description: 给定不同的硬币，以及指定的零钱，求最少兑换次数
//  @Solution: dp[i]: 给定i的钱能兑换的最少次数
//  @param coins
//  @param amount
//  @return int
//
func CoinChange(coins []int, amount int) int {
    // 边界
    max := math.MaxInt32
    dp := make([]int, amount+1)
    for i := 0; i <= amount; i++ {
        dp[i] = max
    }
    dp[0] = 0
    // 转移方程
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            // 不兑换 or 兑换+1
            dp[i] = awesome.MinInt(dp[i], dp[i-coin]+1)
        }
    }
    // 最值
    if max == dp[amount] {
        return -1
    }
    return dp[amount]
}
