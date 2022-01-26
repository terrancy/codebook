package test

import (
    "awesome/stack"
    "fmt"
    "testing"
)

//
//  TestCalculator
//  @Description: 表达式计算器
//  @param t
//
func TestCalculator(t *testing.T) {
    expression := "(2*(3-4))*5"
    res := stack.Calculator(expression)
    fmt.Println(res)
}
