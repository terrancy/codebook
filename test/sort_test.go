package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/base"
)

// —— 冒泡排序 ——

var bubbleSortCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{3, 4, 2, 9, 1, 8}, []int{1, 2, 3, 4, 8, 9}},
	{"test_case_2", []int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
	{"test_case_3", []int{}, []int{}},
	{"test_case_4", []int{1}, []int{1}},
	{"test_case_5", []int{3, 2, 1}, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range bubbleSortCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.BubbleSort(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— 选择排序 ——

var selectSortCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, []int{}},
	{"test_case_3", []int{1}, []int{1}},
	{"test_case_4", []int{1, 2, 3}, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
}

func TestSelectSort(t *testing.T) {
	for _, tt := range selectSortCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.SelectSort(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— 插入排序 ——

var insertSortCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, []int{}},
	{"test_case_3", []int{1}, []int{1}},
	{"test_case_4", []int{1, 2, 3}, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
}

func TestInsertSort(t *testing.T) {
	for _, tt := range insertSortCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.InsertSort(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— 朴素快排 ——

var quickSortOldCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, nil},
	{"test_case_3", []int{1}, []int{1}},
	{"test_case_4", []int{1, 2, 3}, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
}

func TestQuickSortOld(t *testing.T) {
	for _, tt := range quickSortOldCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.QuickSortOld(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— 快排优化 ——

var quickSortCases = []struct {
	name     string
	nums     []int
	left     int
	right    int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, 0, 4, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, 0, -1, []int{}},
	{"test_case_3", []int{1}, 0, 0, []int{1}},
	{"test_case_4", []int{1, 2, 3}, 0, 2, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, 0, 2, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, 0, 3, []int{1, 1, 2, 2}},
}

func TestQuickSort(t *testing.T) {
	for _, tt := range quickSortCases {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if len(nums) > 0 {
				base.QuickSort(nums, tt.left, tt.right)
			}
			assert.Equal(t, tt.expected, nums)
		})
	}
}

// —— 三路快排 ——

var quickSortIICases = []struct {
	name     string
	nums     []int
	left     int
	right    int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, 0, 4, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, 0, -1, []int{}},
	{"test_case_3", []int{1}, 0, 0, []int{1}},
	{"test_case_4", []int{1, 2, 3}, 0, 2, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, 0, 2, []int{1, 2, 3}},
	{"test_case_6", []int{5, 2, 1, 3, 4, 3, 3}, 0, 6, []int{1, 2, 3, 3, 3, 4, 5}},
}

func TestQuickSortII(t *testing.T) {
	for _, tt := range quickSortIICases {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if len(nums) > 0 {
				base.QuickSortII(nums, tt.left, tt.right)
			}
			assert.Equal(t, tt.expected, nums)
		})
	}
}

// —— 堆排序（升序） ——

var heapSortASCCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5}},
	{"test_case_2", []int{}, []int{}},
	{"test_case_3", []int{1}, []int{1}},
	{"test_case_4", []int{1, 2, 3}, []int{1, 2, 3}},
	{"test_case_5", []int{3, 2, 1}, []int{1, 2, 3}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{1, 1, 2, 2}},
}

func TestHeapSortASC(t *testing.T) {
	for _, tt := range heapSortASCCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.HeapSortASC(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— 堆排序（降序） ——

var heapSortDESCCases = []struct {
	name     string
	nums     []int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, []int{5, 4, 3, 2, 1}},
	{"test_case_2", []int{}, []int{}},
	{"test_case_3", []int{1}, []int{1}},
	{"test_case_4", []int{1, 2, 3}, []int{3, 2, 1}},
	{"test_case_5", []int{3, 2, 1}, []int{3, 2, 1}},
	{"test_case_6", []int{2, 1, 2, 1}, []int{2, 2, 1, 1}},
}

func TestHeapSortDESC(t *testing.T) {
	for _, tt := range heapSortDESCCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.HeapSortDESC(tt.nums)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// —— Top K 问题 ——

var findLastKthCases = []struct {
	name     string
	nums     []int
	k        int
	expected []int
}{
	{"test_case_1", []int{5, 2, 3, 1, 4}, 3, []int{3, 4, 5}},
	{"test_case_2", []int{5, 2, 3, 1, 4}, 1, []int{5}},
	{"test_case_3", []int{5, 2, 3, 1, 4}, 5, []int{1, 2, 3, 4, 5}},
	{"test_case_4", []int{5, 2, 3, 1, 4}, 0, []int{}},
}

func TestFindLastKth(t *testing.T) {
	for _, tt := range findLastKthCases {
		t.Run(tt.name, func(t *testing.T) {
			res := base.FindLastKth(tt.nums, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}
