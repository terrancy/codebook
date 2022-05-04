package challenge

//
//  FindTheWinner
//  @Description: 1823. 找出游戏的获胜者 约瑟夫环
//  @Description: 这题数组[1,2,...,n],从1开始
//  @param n
//  @param m
//  @return int
//
func FindTheWinner(n, m int) int {
    res := 0
    for i := 2; i <= n; i++ {
        res = (res + m) % i
    }
    return res + 1
}
