package dp

import (
    "math"
)

//
//  ThrowEggs
//  @Description: NC87 丢棋子问题 / LC 887鸡蛋掉落
//  @param n
//  @param k
//  @return int
//
func ThrowEggs(n int, k int) int {
    if n < 2 || k == 1 {
        return n
    }
    // 楼层有序可用二分查找,最少需要多少棋子
    bs := log2n(n) + 1
    if k >= bs {
        return bs
    }
    //dp[k]用k颗棋子最多恰好在第几层摔碎
    dp := make([]int, k+1)
    dp[1] = 1
    for i := 2; ; i++ {
        for j := k; j >= 1; j-- {
            dp[j] = dp[j-1] + dp[j] + 1
            if dp[j] >= n {
                return i
            }
        }
    }
}

func log2n(n int) int {
    return int(math.Log2(float64(n)))
}

//
//  ThrowEggsII
//  @Description:
//  @param n
//  @param k
//  @return int
//
func ThrowEggsII(n int, k int) int {
    dp := make([]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = i
    }
    for i := 2; i <= k; i++ {
        dp2 := make([]int, n+1)
        x := 1
        for j := 1; j < n; j++ {
            for x < j && maxInt(dp[x-1], dp2[j-x]) > maxInt(dp[x], dp2[j-x-1]) {
                x++
            }
            dp2[j] = maxInt(dp[x-1], dp2[j-x]) + 1
        }
        dp = dp2
    }
    return dp[n]
}
