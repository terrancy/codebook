package main

//
//  isPail
//  @Description: NC141 判断是否为回文字符串
//  @param str
//  @return bool
//
func isPail(str string) bool {
    n := len(str)
    if n == 1 {
        return true
    }
    
    for l, r := 0, n-1; l < r; {
        if str[l] != str[r] {
            return false
        }
        l++
        r--
    }
    return true
}
