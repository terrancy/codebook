package test

import (
    "awesome/dp"
    "fmt"
    "testing"
)

//
//  TestLongCommonSubstring
//  @Description: 最长公共子串
//  @param t
//
func TestLongCommonSubstring(t *testing.T) {
    str1 := "1AB2345CD"
    str2 := "12345EF"
    //res := longCommonSubstring(str1, str2)
    //fmt.Print(res)
    ans := dp.LongCommonSubstring(str1, str2)
    fmt.Println(ans)
}

//
//  TestThrowEggs
//  @Description: NC87 丢棋子问题
//  @Description:
//  @param t
//
func TestThrowEggs(t *testing.T) {
    ans := dp.ThrowEggsII(105, 2)
    fmt.Println(ans)
}
