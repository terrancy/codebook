package test

import (
    "awesome/challenge"
    "fmt"
    "testing"
)

//
//  TestPlatesBetweenCandles
//  @Description: 2055. 蜡烛之间的盘子
//  @param t
//
func TestPlatesBetweenCandles(t *testing.T) {
    s := "**|**|***|"
    queries := [][]int{{2, 5}, {5, 9}}
    res := challenge.PlatesBetweenCandles(s, queries)
    fmt.Println(res)
}

//
//  TestBestRotation
//  @Description: 798. 得分最高的最小轮调798. 得分最高的最小轮调
//  @param t
//
func TestBestRotation(t *testing.T) {
    num := []int{2, 3, 1, 4, 0}
    res := challenge.BestRotation(num)
    fmt.Println(res)
}
