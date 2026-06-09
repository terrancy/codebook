package heap

import "container/heap"

type IHeap []int

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() any {
	last := len(*h) - 1
	x := (*h)[last]
	*h = (*h)[:last]
	return x
}

func NewIHeap() heap.Interface {
	return &IHeap{}
}
