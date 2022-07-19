package strings

import (
    "strings"
)

//
// ReverseSentence
// @Description: JZ73 翻转单词序列
// @param str
// @return string
//
func ReverseSentence(str string) string {
    if str == "" {
        return ""
    }
    data := strings.Split(str, " ")
    for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
    return strings.Join(data, " ")
}
