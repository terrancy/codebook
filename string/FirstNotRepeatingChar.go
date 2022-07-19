package strings

//
// FirstNotRepeatingChar
// @Description: NC31 第一个只出现一次的字符
// @param str
// @return int
//
func FirstNotRepeatingChar(str string) int {
    n := len(str)
    if n < 2 {
        return n
    }
    data := [256]int{}
    for i := 0; i < n; i++ {
        data[str[i]]++
    }
    for i := 0; i < n; i++ {
        if data[str[i]] == 1 {
            return i
        }
    }
    return -1
}
