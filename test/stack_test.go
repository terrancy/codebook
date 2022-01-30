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

//
//  TestIsPopOrder
//  @Description: JZ31 栈的压入、弹出序列
//  @param t
//
func TestIsPopOrder(t *testing.T) {
    source := []int{1, 2, 3, 4, 5}
    target := []int{4, 5, 3, 2, 1}
    res := stack.IsPopOrder(source, target)
    fmt.Println(res)
}
