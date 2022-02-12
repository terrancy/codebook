package strings

import (
    "awesome"
)

//
//  Permutation
//  @Description: NC121 字符串的排列
//  @param str
//  @return []string
//
func Permutation(str string) []string {
    n := len(str)
    if n == 0 {
        return nil
    }
    // 排序
    bytes := awesome.StrSort(str)
    visited := make([]bool, n)
    res := make([]string, 0)
    backtrack(bytes, visited, &res, []byte(""))
    return res
}

//
//  backtrack
//  @Description: 回溯法遍历
//  @param bytes
//  @param visited
//  @param res
//  @param tmp
//
func backtrack(bytes []byte, visited []bool, res *[]string, tmp []byte) {
    // 结束条件
    if len(tmp) == len(bytes) {
        *res = append(*res, string(tmp))
    }
    
    for i := 0; i < len(bytes); i++ {
        if visited[i] {
            continue
        }
        // 选择 visited[i-1]，true: 父子节点，false：兄弟节点
        if i > 0 && bytes[i] == bytes[i-1] && !visited[i-1] {
            continue
        }
        // 加入
        tmp = append(tmp, bytes[i])
        visited[i] = true
        // 递归
        backtrack(bytes, visited, res, tmp)
        // 回退
        visited[i] = false
        tmp = tmp[:len(tmp)-1]
    }
}
