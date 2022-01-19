package test

import (
    strings "awesome/string"
    "fmt"
    "testing"
)

//
//  TestKMP
//  @Description: NC149 kmp算法
//  @param t
//
func TestKMP(t *testing.T) {
    patten := "abab"
    target := "abacabab"
    cnt := strings.KMP(target, patten)
    fmt.Println(cnt)
}
