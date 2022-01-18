package main

import (
    "testing"
)

func testLongCommonSubstring(t *testing.T) {
    str1 := "1AB2345CD"
    str2 := "12345EF"
    //res := longCommonSubstring(str1, str2)
    //fmt.Print(res)
    ans := longCommonSubstring(str1, str2)
    if ans != "123" {
        t.Errorf("测试不通过:%s", ans)
    }
}
