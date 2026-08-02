package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/link"
)

var reverseCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
}

func TestLinkReverse(t *testing.T) {
	for _, tt := range reverseCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.Reverse(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var halfReverseCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"odd_len", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 7, 6, 5}},
}

func TestHalfReverse(t *testing.T) {
	for _, tt := range halfReverseCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.HalfReverse(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var deleteDuplicatesCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 2, 2}, []int{1, 2}},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, tt := range deleteDuplicatesCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.DeleteDuplicates(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var deleteDuplicatesIICases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 2, 2}, []int{1}},
}

func TestDeleteDuplicatesII(t *testing.T) {
	for _, tt := range deleteDuplicatesIICases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.DeleteDuplicatesII(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var removeNthFromEndCases = []struct {
	name     string
	data     []int
	n        int
	expected []int
}{
	{"remove_last", []int{1, 2}, 2, []int{2}},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tt := range removeNthFromEndCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.RemoveNthFromEndII(head, tt.n)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var reverseBetweenCases = []struct {
	name     string
	data     []int
	m        int
	n        int
	expected []int
}{
	{"normal", []int{1, 2}, 1, 2, []int{2, 1}},
}

func TestReverseBetween(t *testing.T) {
	for _, tt := range reverseBetweenCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.ReverseBetweenII(head, tt.m, tt.n)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var reverseKGroupCases = []struct {
	name     string
	data     []int
	k        int
	expected []int
}{
	{"k_is_2", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{2, 1, 4, 3, 6, 5, 8, 7, 9}},
}

func TestReverseKGroup(t *testing.T) {
	for _, tt := range reverseKGroupCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.ReverseKGroupII(head, tt.k)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var insertionSortListCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{2, 4, 1}, []int{1, 2, 4}},
}

func TestInsertSortList(t *testing.T) {
	for _, tt := range insertionSortListCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.InsertionSortList(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var swapLinkedPairCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
}

func TestSwapLinkedPair(t *testing.T) {
	for _, tt := range swapLinkedPairCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.SwapLinkedPair(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var sortLinkedListCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 3, 2, 2, 3, 1}, []int{1, 1, 2, 2, 3, 3}},
}

func TestSortLinkedList(t *testing.T) {
	for _, tt := range sortLinkedListCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.SortLinkedList(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var rotateLinkedListCases = []struct {
	name     string
	data     []int
	k        int
	expected []int
}{
	{"large_k", []int{1, 2, 3}, 1000000000, []int{3, 1, 2}},
}

func TestRotateLinkedList(t *testing.T) {
	for _, tt := range rotateLinkedListCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.RotateLinkedList(head, tt.k)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var plusOneCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"carry", []int{9, 9, 9}, []int{1, 0, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range plusOneCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.PlusOne(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var oddEvenListCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{"normal", []int{1, 4, 6, 3, 7}, []int{1, 6, 7, 4, 3}},
}

func TestOddEvenList(t *testing.T) {
	for _, tt := range oddEvenListCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildListNode(tt.data)
			dummy := link.OddEvenListII(head)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}

var copyRandomListCases = []struct {
	name string
	data []int
}{
	{"normal", []int{1, 2, 3, 4, 5}},
}

func TestCopyRandomList(t *testing.T) {
	for _, tt := range copyRandomListCases {
		t.Run(tt.name, func(t *testing.T) {
			head := link.BuildRandomListNode(tt.data)
			cloned := link.CopyRandomList(head)
			assert.NotNil(t, cloned)
		})
	}
}

var mergeKListsCases = []struct {
	name     string
	data     [][]int
	expected []int
}{
	{"normal", [][]int{{1, 2}, {1, 4, 5}, {6}}, []int{1, 1, 2, 4, 5, 6}},
}

func TestMergeKLists(t *testing.T) {
	for _, tt := range mergeKListsCases {
		t.Run(tt.name, func(t *testing.T) {
			list := make([]*link.ListNode, 0)
			for _, item := range tt.data {
				list = append(list, link.BuildListNode(item))
			}
			dummy := link.MergeKLists(list)
			assert.Equal(t, tt.expected, link.Serialize(dummy))
		})
	}
}
