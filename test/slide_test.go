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
