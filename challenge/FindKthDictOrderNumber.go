package challenge

//
// FindKthDictOrderNumber
// @Description: 440. 字典序的第K小数字
// @param n
// @param k
// @return int
//
func FindKthDictOrderNumber(n int, k int) int {
    cur := 1
    k--
    for k > 0 {
        cnt := dfsFindKthDictOrderNumber(cur, cur, n)
        if cnt <= k {
            k -= cnt
            cur++
        } else {
            k--
            cur *= 10
        }
    }
    return cur
}

func dfsFindKthDictOrderNumber(l, r, n int) int {
    if l > n {
        return 0
    }
    if r > n {
        r = n
    }
    return r - l + 1 + dfsFindKthDictOrderNumber(l*10, r*10+9, n)
}
