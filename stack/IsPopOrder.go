package stack
//
//  IsPopOrder
//  @Description: JZ31 栈的压入、弹出序列
//  @param pushV
//  @param popV
//  @return bool
//
func IsPopOrder(pushV []int, popV []int) bool {
    stack := make([]int, 0)
    for i, j := 0, 0; i < len(pushV); {
        if pushV[i] != popV[j] {
            stack = append(stack, pushV[i])
            i++
        } else {
            i++
            j++
            for len(stack) != 0 && stack[len(stack)-1] == popV[j] {
                stack = stack[:len(stack)-1]
                j++
            }
        }
    }
    if len(stack) == 0 {
        return true
    }
    return false
}
