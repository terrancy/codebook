package dp

//
// longCommonSubstring
// @Description:
// @param str1
// @param str2
// @return string
//
func LongCommonSubstring(str1 string, str2 string) string {
    n := len(str1)
    m := len(str2)
    dp := make([]int, m+1)
    maxLen := 0
    maxIndex := 0
    for i := 0; i < n; i++ {
        for j := m - 1; j >= 0; j-- {
            if str1[i] != str2[j] {
                dp[j+1] = 0
                continue
            }
            dp[j+1] = dp[j] + 1
            if dp[j+1] > maxLen {
                maxLen = dp[j+1]
                maxIndex = i
            }
        }
    }
    return str1[maxIndex-maxLen+1 : maxIndex+1]
}
