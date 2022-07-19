package strings

import "math"

//
// strToInt
// @Description: 字符转数字
// @param str
// @return int
//
func strToInt(str string) int {
    n := len(str)
    i := 0
    // 删除前导空格
    for ; i < n; i++ {
        if str[i] != ' ' {
            break
        }
    }
    if i == n {
        return 0
    }
    // 正负号
    sign := 1
    if str[i] == '-' {
        sign = -1
        i++
    } else if str[i] == '+' {
        sign = 1
        i++
    }
    
    // 获取数值
    str = str[i:]
    chars := []byte(str)
    ans := 0
    for _, char := range chars {
        if char >= '0' && char <= '9' {
            ans = ans*10 + int(char-'0')*sign
            // 是否溢出
            if sign == 1 && ans > math.MaxInt32 {
                ans = math.MaxInt32
            }
            if sign == -1 && ans < math.MinInt32 {
                ans = math.MinInt32
            }
        } else {
            break
        }
    }
    return ans
}
