package dp

//
//  Fib
//  @Description: 剑指 Offer 10- I. 斐波那契数列
//  @param n
//  @return int
//
func Fib(n int) int {
    if n < 2 {
        return n
    }
    res := 0
    for a, b, i := 0, 1, 2; i <= n; i++ {
        res = a + b
        a, b = b, res
    }
    return res
}
