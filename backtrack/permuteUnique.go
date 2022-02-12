package backtrack

import (
    "sort"
)

//
//  permuteUnique
//  @Description: NC42 有重复项数字的全排列
//  @Description: 给出一组可能包含重复项的数字，返回该组数字的所有排列。结果以字典序升序排列
//  @param num
//  @return [][]int
//
func PermuteUnique(num []int) [][]int {
    sort.Ints(num)
    use := make([]bool, len(num))
    res := make([][]int, 0)
    backtrackPermuteUnique(num, use, &res, []int{})
    return res
}

//
//  backtrackPermuteUnique
//  @Description: 全排列回溯框架
//  @param num
//  @param used
//  @param res
//  @param tmp
//
func backtrackPermuteUnique(num []int, used []bool, res *[][]int, tmp []int) {
    if len(num) == len(tmp) {
        //*res = append(*res, tmp)
        // !!使用切片的copy方法,不然容易出错
        target := make([]int, len(num))
        copy(target, tmp)
        *res = append(*res, target)
    }
    
    for i := 0; i < len(num); i++ {
        // 过滤父子重复
        if used[i] {
            continue
        }
        
        // 过滤兄弟重复
        if i > 0 && !used[i-1] && num[i-1] == num[i] {
            continue
        }
        
        // 选择与撤销
        tmp = append(tmp, num[i])
        used[i] = true
        backtrackPermuteUnique(num, used, res, tmp)
        used[i] = false
        tmp = tmp[:len(tmp)-1]
    }
}
