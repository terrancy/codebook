package challenge

//
// ConsecutiveNumbersSum
// @Description: 829. 连续整数求和
// @param num
// @return int
//
func ConsecutiveNumbersSum(n int) int {
    ans := 0
    n = n * 2
    for i := 1; i*i < n; i++ {
        if n%i != 0 {
            continue
        }
        if (n/i-(i-1))%2 == 0 {
            ans++
        }
    }
    return ans
}
