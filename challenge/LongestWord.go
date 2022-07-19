package challenge

import "strings"

//
// LongestWord
// @Description:
// @param words
// @return string
//
func LongestWord(words []string) string {
    // 哈希表
    dict := make(map[string]int)
    for idx, str := range words {
        dict[str] = idx
    }
    res := ""
    for _, str := range words {
        n := len(str)
        m := len(res)
        // 判断长度比较
        if n < m {
            continue
        }
        // 长度一致时, 取字典序最小的单词
        if n == m && strings.Compare(res, str) < 0 {
            continue
        }
        // 长度更长的，判断是否又小的组成
        flag := true
        for idx, _ := range str {
            substr := str[0 : idx+1]
            if _, ok := dict[substr]; !ok {
                flag = false
                break
            }
        }
        if flag == true {
            res = str
        }
        
    }
    return res
}
