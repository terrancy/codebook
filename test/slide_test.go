package test

import (
    "awesome/slide"
    "fmt"
    "testing"
)

//
//  TestMinWindow
//  @Description: 最小覆盖子串
//  @param t
//
func TestMinWindow(t *testing.T) {
    str := "abcAbA"
    target := "AA"
    res := slide.MinWindow(str, target)
    fmt.Println(res)
}

//
//  TestMaxWater
//  @Description: NC128 接雨水问题
//  @param t
//
func TestMaxWater(t *testing.T) {
    data := []int{3, 1, 2, 5, 2, 4}
    res := slide.MaxWater(data)
    fmt.Println(res)
}
