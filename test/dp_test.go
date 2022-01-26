package dp

import (
    "awesome/dp"
    "fmt"
    "testing"
)

func TestLongCommonSubstring(t *testing.T) {
    str1 := "1AB2345CD"
    str2 := "12345EF"
    //res := longCommonSubstring(str1, str2)
    //fmt.Print(res)
    ans := dp.longCommonSubstring(str1, str2)
    fmt.Println(ans)
}

func TestThrowEggs(t *testing.T) {
    ans := dp.ThrowEggs(10, 1)
}
