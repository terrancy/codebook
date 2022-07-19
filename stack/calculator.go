package stack

// 计算器系列相关

//
// Calculator
// @Description: 请写一个整数计算器，支持加减乘三种运算和括号
// @param s
// @return int
//
func Calculator(s string) int {
    // 数据初始化
    stack := make([]int, 0)
    data := []byte(s)
    number := 0
    var operator byte = '+'
    res := 0
    
    n := len(data)
    for i := 0; i < n; i++ {
        // 判断 0 - 9
        if data[i] >= '0' && data[i] <= '9' {
            number = 10*number + int(data[i]-'0')
        }
        // 判断 '('
        if data[i] == '(' {
            start, end := i+1, i+1
            count := 1
            for count > 0 {
                if data[end] == '(' {
                    count++
                }
                if data[end] == ')' {
                    count--
                }
                end++
            }
            i = end - 1
            number = Calculator(s[start : end-1])
        }
        
        // 计算 '+'/'-'/'*'
        if data[i] < '0' || data[i] > '9' || i == len(data)-1 {
            switch operator {
            case '+':
                stack = append(stack, number)
            case '-':
                stack = append(stack, -number)
            case '*':
                stack[len(stack)-1] = stack[len(stack)-1] * number
            }
            number = 0
            operator = data[i]
        }
    }
    // 累加操作
    for i := len(stack) - 1; i >= 0; i-- {
        res += stack[i]
    }
    return res
}
