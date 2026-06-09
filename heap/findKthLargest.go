package heap

import (
	"container/heap"
)

// FindKthLargest
// title: LC215. 数组中的第K个最大元素
// description: 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素
// link: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// params:
//   - nums: 整数数组 [1,2,3,4,5,6]
//   - k: 第 k 个最大的元素 [2]
//
// return: 第 k 个最大的元素 [5]
func FindKthLargest(nums []int, k int) int {
	h := NewIHeap()
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

func FindKthLargestWithHeapsort(nums []int, k int) int {
	Heapsort(nums)
	m := len(nums)
	return nums[m-k]
}
