package test

import (
    "awesome/challenge"
    "fmt"
    "testing"
)

//
//  TestConsecutiveNumbersSum
//  @Description: 829. 连续整数求和
//  @param t
//
func TestConsecutiveNumbersSum(t *testing.T) {
    res := challenge.ConsecutiveNumbersSum(5)
    fmt.Println(res)
}

//
//  TestNumUniqueEmails
//  @Description: 929. 独特的电子邮件地址
//  @param t
//
func TestNumUniqueEmails(t *testing.T) {
    emails := []string{
        "test.email+alex@leetcode.com",
        "test.e.mail+bob.cathy@leetcode.com",
        "testemail+david@lee.tcode.com",
    }
    res := challenge.NumUniqueEmail(emails)
    fmt.Println(res)
}

//
//  TestIsBoomerang
//  @Description: 1037. 有效的回旋镖
//  @param t
//
func TestIsBoomerang(t *testing.T) {
    points := [][]int{{1, 1}, {2, 2}, {3, 2}}
    result := challenge.IsBoomerang(points)
    fmt.Println(result)
}
