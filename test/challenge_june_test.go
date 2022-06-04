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
