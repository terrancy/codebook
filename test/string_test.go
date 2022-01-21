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

//
//  TestPermutation
//  @Description: 字符串的全排列
//  @param t
//
func TestPermutation(t *testing.T) {
    str := "aab"
    res := strings.Permutation(str)
    fmt.Println(res)
}

//
//  TestGetLongestPalindrome
//  @Description: 最长回文子串
//  @param t
//
func TestGetLongestPalindrome(t *testing.T) {
    str := "ccbcabaabba"
    res := strings.GetLongestPalindrome(str)
    fmt.Println(res)
}
