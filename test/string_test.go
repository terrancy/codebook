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
func TestLongestPalindrome(t *testing.T) {
    str := "ccbcabaabba"
    res := strings.LongestPalindrome(str)
    fmt.Println(res)
}

//
//  TestLongRepeatedSubstring
//  @Description: NC142 最长重复子串
//  @param t
//
func TestLongRepeatedSubstring(t *testing.T) {
    str := "abcab"
    res := strings.LongRepeatedSubstring(str)
    fmt.Println(res)
}

//
//  TestFirstNotRepeatingChar
//  @Description: NC31 第一个只出现一次的字符
//  @param t
//
func TestFirstNotRepeatingChar(t *testing.T) {
    str := "google"
    res := strings.FirstNotRepeatingChar(str)
    fmt.Println(res)
}

//
//  TestCheckIpAddress
//  @Description: NC113 验证IP地址
//  @param t
//
func TestCheckIpAddress(t *testing.T) {
    str := "127.0.0.1"
    res := strings.CheckIpAddress(str)
    fmt.Println(res)
}

//
//  TestReverseSentence
//  @Description: JZ73 翻转单词序列
//  @param t
//
func TestReverseSentence(t *testing.T) {
    str := "nowcoder. a am I"
    res := strings.ReverseSentence(str)
    fmt.Println(res)
}

//
//  TestGenerateParenthesis
//  @Description: NC26 括号生成
//  @param t
//
func TestGenerateParenthesis(t *testing.T) {
    n := 3
    res := strings.GenerateParenthesis(n)
    fmt.Println(res)
}

//
//  TestCheckBracketValid
//  @Description: 20. 有效的括号
//  @param t
//
func TestCheckBracketValid(t *testing.T) {
    str := "()[]{}"
    res := strings.CheckBracketValid(str)
    fmt.Println(res)
}
