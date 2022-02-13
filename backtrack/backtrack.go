package backtrack

import (
    "awesome/array"
    strings "awesome/string"
)

//
//  SubsetsII
//  @Description: NC27 集合的所有子集(一)
//  @Solution: 1、数组长度 2、遍历的起始值设置 3、回溯与记录
//  @param nums
//  @return [][]int
//
func Subsets(nums []int) [][]int {
    return array.SubsetsII(nums)
}

//
//  BTPermuteUnique
//  @Description: NC121 字符串的排列
//  @param str
//  @return []string
//
func BTPermutation(str string) []string {
    return strings.Permutation(str)
}

//
//  BTPermuteUnique
//  @Description: NC42 有重复项数字的全排列
//  @param num
//  @return [][]int
//
func BTPermuteUnique(num []int) [][]int {
    return PermuteUnique(num)
}

//
//  CombinationSum2
//  @Description: NC46 加起来和为目标值的组合(二)
//  @param num
//  @param target
//  @return [][]int
//
func CombinationSum2(num []int, target int) [][]int {
    return array.CombinationSum2(num, target)
}

//
//  GenerateParenthesis
//  @Description: NC26 括号生成
//  @param n
//  @return []string
//
func GenerateParenthesis(n int) []string {
    return strings.GenerateParenthesis(n)
}

//////////////////////////////////////////////////////
// 图的遍历DFS
//////////////////////////////////////////////////////

//
//  MovingCount
//  @Description: JZ13 机器人的运动范围
//  @param m
//  @param n
//  @param k
//  @return int
//
func MovingCount(m int, n int, k int) int {
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }
    return DSFMovingCount(m, n, k, visited, 0, 0)
}

func DSFMovingCount(m int, n int, k int, visited [][]bool, i int, j int) int {
    if i >= m || j >= n || getSum(i, j) > k || visited[i][j] {
        return 0
    }
    visited[i][j] = true
    return 1 + DSFMovingCount(m, n, k, visited, i+1, j) + DSFMovingCount(m, n, k, visited, i, j+1)
}
func getSum(a, b int) int {
    res := 0
    for a > 0 {
        res += a % 10
        a /= 10
    }
    for b > 0 {
        res += b % 10
        b /= 10
    }
    return res
}

//
//  HasMatrixPath
//  @Description: JZ12 矩阵中的路径
//  @param board
//  @param word
//  @return bool
//

var hasMatrixPathDirections = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var other byte = '#'

func HasMatrixPath(board [][]byte, word string) bool {
    // 路径可以从矩阵的任意格子开始
    n, m := len(board), len(board[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if DFSHasMatrixPath(board, word, i, j, 0) {
                return true
            }
        }
    }
    return false
}

//
//  DFSHasMatrixPath
//  @Description: 先根遍历
//  @param board
//  @param word
//  @param i
//  @param j
//  @param k
//  @return bool
//
func DFSHasMatrixPath(board [][]byte, word string, i, j, k int) bool {
    if i >= len(board) || i < 0 {
        return false
    }
    if j >= len(board[0]) || j < 0 {
        return false
    }
    if word[k] != board[i][j] {
        return false
    }
    if k == len(word)-1 {
        return true
    }
    for _, direction := range hasMatrixPathDirections {
        board[i][j], other = other, board[i][j]
        if DFSHasMatrixPath(board, word, i+direction[0], j+direction[1], k+1) {
            return true
        }
        board[i][j], other = other, board[i][j]
    }
    return false
}
