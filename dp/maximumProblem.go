package dp

import (
    "awesome"
    "awesome/array"
)

//
// MaxLengthRepeatLess
// @Description: NC41 最长无重复子数组
// @Solution: 哈希表 + 双指针, l, r := -1, 0
// @param nums
// @return int
//
func MaxLengthRepeatLess(nums []int) int {
    return array.MaxLengthRepeatLess(nums)
}

//
// LongestPalindrome
// @Description: NC17 最长回文子串
// @Description: 中心扩散法
// @param str
// @return string
//
func LongestPalindrome(str string) string {
    n := len(str)
    maxLen := 1
    start := 0
    
    // 奇数回文串
    for k := 0; k < n; k++ {
        for i, j := k-1, k+1; i >= 0 && j < n && str[i] == str[j]; i, j = i-1, j+1 {
            if j-i+1 > maxLen {
                maxLen = j - i + 1
                start = i
            }
        }
    }
    
    // 偶数回文串
    for k := 0; k < n; k++ {
        for i, j := k, k+1; i >= 0 && j < n && str[i] == str[j]; i, j = i-1, j+1 {
            if j-i+1 > maxLen {
                maxLen = j - i + 1
                start = i
            }
        }
    }
    
    return str[start : start+maxLen]
}

//
// LongestPalindromeSubsequence
// @Description: 最长回文子序列
// @param str
// @return int
//
func LongestPalindromeSubsequence(str string) int {
    n := len(str)
    if n == 1 {
        return 1
    }
    // dp[i][j]: 字符串i到j的回文串长度 => dp[0][n-1]
    dp := make([][]int, n)
    // 边界
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }
    // 状态转移 dp[i][j] = max(dp[i+1][j], dp[i][j-1])
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            // i,j往中间靠
            if str[i] == str[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = awesome.MaxInt(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    return dp[0][n-1]
}
