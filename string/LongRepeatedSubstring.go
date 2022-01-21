package strings

//
//  LongRepeatedSubstring
//  @Description: NC142 最长重复子串
//  @param str
//  @return int
//
func LongRepeatedSubstring(str string) int {
    n := len(str)
    if n < 2 {
        return n
    }
    
    res := 0
    for i := int(n / 2); i > 0; i-- {
        for j := 0; j < n-i; j++ {
            if str[j] == str[i+j] {
                res++
            } else {
                res = 0
            }
            if res == i {
                return res * 2
            }
        }
    }
    return 0
}
