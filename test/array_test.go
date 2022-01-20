package test

import (
    "awesome/array"
    "fmt"
    "testing"
)

//
//  TestMajorityElement
//  @Description: NC73 数组中出现次数超过一半的数字
//  @param t
//
func TestMajorityElement(t *testing.T) {
    data := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
    res := array.MajorityElement(data)
    fmt.Println(res)
}
