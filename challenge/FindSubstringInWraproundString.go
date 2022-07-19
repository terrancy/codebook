package challenge

import "awesome"

//
// FindSubstringInWrapRoundString
// @Description: 467. 环绕字符串中唯一的子字符串
// @param p
// @return int
//
func FindSubstringInWrapRoundString(p string) (ans int) {
    dp := [26]int{}
    k := 0
    for i, ch := range p {
        if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
            k++
        } else {
            k = 1
        }
        dp[ch-'a'] = awesome.MaxInt(dp[ch-'a'], k)
    }
    for _, v := range dp {
        ans += v
    }
    return
}
