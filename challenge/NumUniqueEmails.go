package challenge

import "strings"

//
// NumUniqueEmail
// @Description: 929. 独特的电子邮件地址
// @param emails
// @return int
//
func NumUniqueEmail(emails []string) int {
    dict := map[string]struct{}{}
    temp := ""
    idx := 0
    for _, email := range emails {
        // 获取@的前部分
        idx = strings.Index(email, "@")
        // 保留第一个+之前的字符串
        temp = strings.SplitN(email[:idx], "+", 2)[0]
        // 过滤.
        temp = strings.ReplaceAll(temp, ".", "")
        // 去重
        dict[temp+email[idx:]] = struct{}{}
    }
    // 计算
    return len(dict)
}
