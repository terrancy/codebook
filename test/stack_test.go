package test

import (
    "awesome/queue"
    "awesome/stack"
    "fmt"
    "testing"
)

//
// TestCalculator
// @Description: 表达式计算器
// @param t
//
func TestCalculator(t *testing.T) {
    expression := "(2*(3-4))*5"
    res := stack.Calculator(expression)
    fmt.Println(res)
}

//
// TestIsPopOrder
// @Description: JZ31 栈的压入、弹出序列
// @param t
//
func TestIsPopOrder(t *testing.T) {
    source := []int{1, 2, 3, 4, 5}
    target := []int{4, 5, 3, 2, 1}
    res := stack.IsPopOrder(source, target)
    fmt.Println(res)
}

// ============ 队列 ============

//
// TestMaxInWindows
// @Description: NC82 滑动窗口的最大值 => 单调队列
// @param t
//
func TestMaxInWindows(t *testing.T) {
    num := []int{2, 3, 4, 2, 6, 2, 5, 1}
    ans := queue.MaxInWindows(num, 3)
    fmt.Println(ans)
}
