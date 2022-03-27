package challenge

//
//  MissingRolls
//  @Description: 2028. 找出缺失的观测数据
//  @param rolls
//  @param mean
//  @param n
//  @return []int
//
func MissingRolls(rolls []int, mean int, n int) []int {
    s := mean * (len(rolls) + n)
    for _, roll := range rolls {
        s -= roll
    }
    if s < n || s > 6*n {
        return nil
    }
    ans, d := make([]int, n), s/n
    for i := 0; i < s%n; i++ {
        ans[i] = d + 1
    }
    for i := s % n; i < n; i++ {
        ans[i] = d
    }
    return ans
}
