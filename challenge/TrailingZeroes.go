package challenge

//
// TrailingZeroes
// @Description: 172. 阶乘后的零
// @param n
// @return int
//
func TrailingZeroes(n int) int {
    if n == 0 {
        return 0
    }
    return n/5 + TrailingZeroes(n/5)
}
