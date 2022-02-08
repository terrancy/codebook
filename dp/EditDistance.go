package dp

import "awesome"

//
//  EditDistance
//  @Description: NC196 编辑距离(一)
//  @param word1
//  @param word2
//  @return int
//
func EditDistance(str1 string, str2 string) int {
    n, m := len(str1), len(str2)
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
    }
    // 边界
    for i := 0; i <= n; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = awesome.MinInt(awesome.MinInt(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[n][m]
}

//
//  EditCost
//  @Description: NC35 编辑距离(二)
//  @param str1
//  @param str2
//  @param insertCost 插入操作的成本
//  @param deleteCost 删除操作的成本
//  @param replaceCost 替换操作的成本
//  @return int
//
func EditCost(str1 string, str2 string, insertCost int, deleteCost int, replaceCost int) int {
    n, m := len(str1), len(str2)
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
    }
    // 边界
    for i := 0; i <= n; i++ {
        dp[i][0] = i * deleteCost
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j * insertCost
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = awesome.MinInt(awesome.MinInt(dp[i][j-1]+insertCost, dp[i-1][j]+deleteCost), dp[i-1][j-1]+replaceCost)
            }
        }
    }
    return dp[n][m]
}
