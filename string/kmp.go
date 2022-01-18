package main

//
//  kmp 算法
//  @Description: 给你一个文本串 T ，一个非空模板串 S ，问 S 在 T 中出现了多少次
//  @param target abababab
//  @param patten ababab
//  @return int
//
func kmp(target string, patten string) int {
    n := len(patten)
    m := len(target)
    cnt := 0
    
    // 生成next数组
    next := make([]int, n)
    next[0] = -1
    for i, j := -1, 0; j < n-1; {
        if i == -1 || patten[i] == patten[j] {
            i++
            j++
            next[j] = i
        } else {
            i = next[i]
        }
    }
    
    // 移位匹配
    for i, j := 0, 0; j < m; {
        if i == -1 || patten[i] == target[j] {
            i++
            j++
        } else {
            i = next[i]
        }
        
        // 这段容易出错!!
        if i == n {
            j--
            i = next[i-1]
            cnt++
        }
    }
    
    return cnt
}
