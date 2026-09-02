package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/heap"
)

var findKthLargestCase = []struct {
	name     string
	nums     []int
	k        int
	expected int
}{
	{
		name:     "test1",
		nums:     []int{3, 2, 1, 5, 6, 4},
		k:        2,
		expected: 5,
	},
	{
		name:     "test2",
		nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		k:        4,
		expected: 4,
	},
}

// TestFindKthLargest
// LC215. 数组中的第K个最大元素
func TestFindKthLargest(t *testing.T) {
	for _, tt := range findKthLargestCase {
		t.Run(tt.name, func(t *testing.T) {
			res := heap.FindKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestFindKthLargestWithHeapsort(t *testing.T) {
	for _, tt := range findKthLargestCase {
		t.Run(tt.name, func(t *testing.T) {
			res := heap.FindKthLargestWithHeapsort(tt.nums, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var findKSmallestPairsCases = []struct {
	name     string
	nums1    []int
	nums2    []int
	k        int
	expected [][]int
}{
	{
		name:     "test_case_1",
		nums1:    []int{1, 7, 11},
		nums2:    []int{2, 4, 6},
		k:        3,
		expected: [][]int{{1, 2}, {1, 4}, {1, 6}},
	},
	{
		name:     "test_case_2",
		nums1:    []int{1, 1, 2},
		nums2:    []int{1, 2, 3},
		k:        2,
		expected: [][]int{{1, 1}, {1, 1}},
	},
	{
		name:     "test_case_3",
		nums1:    []int{1, 2},
		nums2:    []int{3},
		k:        3,
		expected: [][]int{{1, 3}, {2, 3}},
	},
	{
		name:     "test_case_4",
		nums1:    []int{},
		nums2:    []int{1, 2, 3},
		k:        2,
		expected: nil,
	},
}

// TestFindKSmallestPairs
// LC373. 查找和最小的 K 对数字
func TestFindKSmallestPairs(t *testing.T) {
	for _, tt := range findKSmallestPairsCases {
		t.Run(tt.name, func(t *testing.T) {
			res := heap.FindKSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if tt.expected == nil {
				assert.Nil(t, res)
				return
			}
			sort.Slice(res, func(i, j int) bool {
				if res[i][0]+res[i][1] != res[j][0]+res[j][1] {
					return res[i][0]+res[i][1] < res[j][0]+res[j][1]
				}
				return res[i][0] < res[j][0]
			})
			sort.Slice(tt.expected, func(i, j int) bool {
				if tt.expected[i][0]+tt.expected[i][1] != tt.expected[j][0]+tt.expected[j][1] {
					return tt.expected[i][0]+tt.expected[i][1] < tt.expected[j][0]+tt.expected[j][1]
				}
				return tt.expected[i][0] < tt.expected[j][0]
			})
			assert.Equal(t, tt.expected, res)
		})
	}
}