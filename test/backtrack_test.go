package test

import (
    "awesome/backtrack"
    "fmt"
    "testing"
)

//
//  TestMovingCount
//  @Description: JZ13 机器人的运动范围
//  @param t
//
func TestMovingCount(t *testing.T) {
    res := backtrack.MovingCount(2, 3, 1)
    fmt.Println(res)
}

//
//  TestHasMatrixPath
//  @Description: JZ12 矩阵中的路径
//  @param t
//
func TestHasMatrixPath(t *testing.T) {
    data := [][]byte{{'a', 'b'}, {'c', 'd'}}
    word := "abcd"
    res := backtrack.HasMatrixPath(data, word)
    fmt.Println(res)
}
