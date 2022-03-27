package strings

//
//  CheckBracketValid
//  @Description: 20. 有效的括号
//  @param str
//  @return bool
//
func CheckBracketValid(str string) bool {
    n := len(str)
    if n%2 == 1 {
        return false
    }
    dict := map[byte]byte{'{': '}', '[': ']', '(': ')'}
    stack := make([]byte, 0)
    for i := 0; i < n; i++ {
        if _, ok := dict[str[i]]; ok {
            // 左半边
            stack = append(stack, dict[str[i]])
        } else {
            // 右半边
            if len(stack) == 0 || str[i] != stack[len(stack)-1] {
                return false
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    if len(stack) == 0 {
        return true
    }
    return false
}
