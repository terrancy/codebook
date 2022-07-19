package dp

import (
    "awesome"
)

//
// CuttingRopes
// @Description: JZ14 剪绳子 - 动态规划
// @param n
// @return int
//
func CuttingRopesDP(number int) int {
    // DP[i] 表示长度为i的绳子的最大乘积
    dp := make([]int, number+1)
    dp[0] = 0
    dp[1] = 0
    dp[2] = 1
    for i := 3; i <= number; i++ {
        for j := 2; j < i; j++ {
            dp[i] = awesome.MaxInt(dp[i], j*awesome.MaxInt(dp[i-j], i-j))
        }
    }
    return dp[number]
}

//
// CuttingRosesGreedy
// @Description: Z14 剪绳子 - 贪心算法
// @param number
// @return int
//
func CuttingRosesGreedy(number int) int {
    // 重点在于公式推导: 将绳子等分为长度为3的若干份时乘积最大
    if number == 2 {
        return 1
    }
    if number == 3 {
        return 2
    }
    n := int64(number / 3)
    m := number % 3
    var res int64
    if m == 0 {
        res = cuttingRopePow(3, n)
    } else if m == 1 {
        res = cuttingRopePow(3, n-1) << 2 % 1000000007
    } else {
        res = cuttingRopePow(3, n) << 1 % 1000000007
    }
    return int(res)
}

//
// cuttingRopePow
// @Description: 乘方操作时需要模操作
// @param x
// @param n
// @return int64
//
func cuttingRopePow(x, n int64) int64 {
    var res int64 = 1
    for n > 0 {
        if n&1 == 1 {
            res = res * x % 1000000007
        }
        x = x * x % 1000000007
        n >>= 1
    }
    return res
}
