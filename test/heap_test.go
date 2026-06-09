package test

import (
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
