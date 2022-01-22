package test

import (
    "awesome/base"
    "fmt"
    "testing"
)

//
//  TestBubbleSort
//  @Description: 冒泡排序
//  @param t
//
func TestBubbleSort(t *testing.T) {
    nums := []int{3, 4, 2, 9, 1, 8,}
    res := base.BubbleSort(nums)
    fmt.Println(res)
}

//
//  TestSelectSort
//  @Description: 选自排序
//  @param t
//
func TestSelectSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.SelectSort(nums)
    fmt.Println(res)
}

//
//  TestQuickSort
//  @Description: 快排
//  @param t
//
func TestQuickSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.QuickSort(nums)
    fmt.Println(res)
}

//
//  TestInsertSort
//  @Description: 插入排序
//  @param t
//
func TestInsertSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.InsertSort(nums)
    fmt.Println(res)
}

//
//  TestHeapSort
//  @Description: 堆排序
//  @param t
//
func TestHeapSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.HeapSort(nums)
    fmt.Println(res)
}
