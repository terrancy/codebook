package strings

//
// getLongestPalindrome
// @Description: NC17 最长回文子串
// @Description: 对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度
// @Solution: 中心扩散法
// @param str
// @return string
//
func LongestPalindrome(str string) int {
    n := len(str)
    maxLen := 1
    
    // 奇数回文
    for i := 0; i < n; i++ {
        for j, k := i-1, i+1; j >= 0 && k < n && str[j] == str[k]; j, k = j-1, k+1 {
            if maxLen < k-j+1 {
                maxLen = k - j + 1
            }
        }
    }
    
    // 偶数回文
    for i := 0; i < n; i++ {
        for j, k := i, i+1; j >= 0 && k < n && str[j] == str[k]; j, k = j-1, k+1 {
            if maxLen < k-j+1 {
                maxLen = k - j + 1
            }
        }
    }
    
    return maxLen
}
