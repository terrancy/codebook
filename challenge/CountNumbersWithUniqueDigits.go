package challenge

//
//  CountNumbersWithUniqueDigits
//  @Description: 357. 统计各位数字都不同的数字个数
//  @param n
//  @return int
//
func CountNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    ans := 10
    for i, last := 2, 9; i <= n; i++ {
        cur := last * (10 - i + 1)
        ans += cur
        last = cur
    }
    return ans
}
