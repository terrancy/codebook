package challenge

//
//  OneEditAway
//  @Description: 面试题 01.05. 一次编辑
//  @param first
//  @param second
//  @return bool
//
func OneEditAway(first, second string) bool {
    m, n := len(first), len(second)
    if m < n {
        first, second = second, first
        m, n = n, m
    }
    if m-n > 1 {
        return false
    }
    for idx, char := range second {
        if first[idx] != byte(char) {
            if m == n {
                return first[idx+1:] == second[idx+1:]
            }
            return first[idx+1:] == second[idx:]
        }
    }
    return true
}
