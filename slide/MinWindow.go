package slide

//
// MinWindow
// @Description: NC28 最小覆盖子串
// @Description: 给出两个字符串 s 和 t，要求在 s 中找出最短的包含 t 中所有字符的连续子串
// @param str
// @param target
// @return string
//
func MinWindow(str string, target string) string {
    // 初始化 窗口
    window := make(map[byte]int)
    need := make(map[byte]int)
    for idx, _ := range target {
        need[target[idx]]++
    }
    start := 0
    const maxLen = 10001
    minLen := maxLen
    for left, right, cnt := 0, 0, 0; right < len(str); right++ {
        // 右指针的判断,是否存在,是否满足条件
        val := str[right]
        if _, ok := need[val]; !ok {
            continue
        }
        window[val]++
        if window[val] == need[val] {
            cnt++
        }
        // 左指针的判断,寻找最大长度,是否满足条件
        for ; cnt == len(need); left++ {
            if minLen > right-left {
                start = left
                minLen = right - left + 1
            }
            val := str[left]
            if _, ok := window[val]; !ok {
                continue
            }
            if window[val] == need[val] {
                cnt--
            }
            window[val]--
        }
    }
    if minLen == maxLen {
        return ""
    }
    return str[start : start+minLen]
}
