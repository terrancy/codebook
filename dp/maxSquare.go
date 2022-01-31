package dp

import "awesome"

func solve(matrix [][]byte) int {
    n := len(matrix)
    if n == 0 {
        return 0
    }
    m := len(matrix[0])
    dp := make([][]int, n)
    max := 0
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' {
                dp[i][j] = 0
            } else if i == 0 || j == 0 {
                dp[i][j] = intVal(matrix[i][j])
            } else {
                dp[i][j] = awesome.MinInt(awesome.MinInt(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
            }
            max = awesome.MaxInt(max, dp[i][j])
        }
    }
    return max * max
}

func intVal(a byte) int {
    if a == '1' {
        return 1
    }
    return 0
}
