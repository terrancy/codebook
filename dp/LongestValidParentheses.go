package dp

import "awesome"

//
// LongestValidParentheses
// @Description: 32. 最长有效括号
// @param s
// @return int
//
func LongestValidParentheses(str string) int {
    maxAns := 0
    dp := make([]int, len(str))
    for i := 1; i < len(str); i++ {
        if str[i] == ')' {
            if str[i-1] == '(' {
                if i >= 2 {
                    dp[i] = dp[i-2] + 2
                } else {
                    dp[i] = 2
                }
            } else if i-dp[i-1] > 0 && str[i-dp[i-1]-1] == '(' {
                if i-dp[i-1] >= 2 {
                    dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
                } else {
                    dp[i] = dp[i-1] + 2
                }
            }
            maxAns = awesome.MaxInt(maxAns, dp[i])
        }
    }
    return maxAns
}
