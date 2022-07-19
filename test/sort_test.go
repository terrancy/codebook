package test

import (
    "awesome/base"
    "fmt"
    "testing"
)

//
// TestBubbleSort
// @Description: 冒泡排序
// @param t
//
func TestBubbleSort(t *testing.T) {
    nums := []int{3, 4, 2, 9, 1, 8}
    res := base.BubbleSort(nums)
    fmt.Println(res)
}

//
// TestSelectSort
// @Description: 选自排序
// @param t
//
func TestSelectSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.SelectSort(nums)
    fmt.Println(res)
}

//
// TestQuickSort
// @Description: 快排
// @param t
//
func TestQuickSortOld(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.QuickSortOld(nums)
    fmt.Println(res)
}

//
// TestQuickSort
// @Description: 快排
// @param t
//
func TestQuickSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    base.QuickSort(nums, 0, len(nums)-1)
    res := nums
    fmt.Println(res)
}

//
// TestQuickSort
// @Description: 快排
// @param t
//
func TestQuickSortII(t *testing.T) {
    nums := []int{5, 2, 1, 3, 4, 3, 3}
    base.QuickSortII(nums, 0, len(nums)-1)
    res := nums
    fmt.Println(res)
}

//
// TestInsertSort
// @Description: 插入排序
// @param t
//
func TestInsertSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.InsertSort(nums)
    fmt.Println(res)
}

//
// TestHeapSort
// @Description: 堆排序
// @param t
//
func TestHeapSort(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.HeapSortDESC(nums)
    fmt.Println(res)
}

//
// TestHeapSort
// @Description: 堆排序
// @param t
//
func TestFindLastKth(t *testing.T) {
    nums := []int{5, 2, 3, 1, 4}
    res := base.FindLastKth(nums, 3)
    fmt.Println(res)
}
