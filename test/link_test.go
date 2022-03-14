package test

import (
    "awesome/link"
    "testing"
)

//
//  TestLinkReverse
//  @Description: 链表翻转
//  @param t
//
func TestLinkReverse(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    head := link.BuildListNode(data)
    dummy := link.Reverse(head)
    link.ShowData(dummy)
}

//
//  TestHalfReverse
//  @Description:
//
func TestHalfReverse(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7}
    head := link.BuildListNode(data)
    dummy := link.HalfReverse(head)
    link.ShowData(dummy)
}

//
//  TestDeleteDuplicates
//  @Description: NC25 删除有序链表中重复的元素-I
//  @param t
//
func TestDeleteDuplicates(t *testing.T) {
    data := []int{1, 2, 2}
    head := link.BuildListNode(data)
    dummy := link.DeleteDuplicates(head)
    link.ShowData(dummy)
}

//
//  TestDeleteDuplicatesII
//  @Description: NC24 删除有序链表中重复的元素-II
//  @param t
//
func TestDeleteDuplicatesII(t *testing.T) {
    data := []int{1, 2, 2}
    head := link.BuildListNode(data)
    dummy := link.DeleteDuplicatesII(head)
    link.ShowData(dummy)
}

//
//  TestRemoveNthFromEnd
//  @Description: NC53 删除链表的倒数第n个节点
//  @param t
//
func TestRemoveNthFromEnd(t *testing.T) {
    data := []int{1, 2}
    head := link.BuildListNode(data)
    dummy := link.RemoveNthFromEndII(head, 2)
    link.ShowData(dummy)
}

//
//  TestReverseBetween
//  @Description: NC21 链表内指定区间反转
//  @param t
//
func TestReverseBetween(t *testing.T) {
    data := []int{1, 2}
    head := link.BuildListNode(data)
    dummy := link.ReverseBetweenII(head, 1, 2)
    link.ShowData(dummy)
}

func TestReverseKGroup(t *testing.T) {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    head := link.BuildListNode(data)
    dummy := link.ReverseKGroupII(head, 2)
    link.ShowData(dummy)
}

//
//  TestInsertSortList
//  @Description: NC244 对链表进行插入排序
//  @param t
//
func TestInsertSortList(t *testing.T) {
    data := []int{2, 4, 1}
    head := link.BuildListNode(data)
    dummy := link.InsertionSortList(head)
    link.ShowData(dummy)
}

//
//  TestSwapLinkedPair
//  @Description: NC186 两两交换链表的节点
//  @param t
//
func TestSwapLinkedPair(t *testing.T) {
    data := []int{1, 2, 3, 4}
    head := link.BuildListNode(data)
    dummy := link.SwapLinkedPair(head)
    link.ShowData(dummy)
}

//
//  sortLinkedList
//  @Description: NC207 排序奇升偶降链表
//
func TestSortLinkedList(t *testing.T) {
    data := []int{1, 3, 2, 2, 3, 1}
    head := link.BuildListNode(data)
    dummy := link.SortLinkedList(head)
    link.ShowData(dummy)
}

//
//  TestRotateLinkedList
//  @Description: NC211 旋转链表
//  @param t
//
func TestRotateLinkedList(t *testing.T) {
    data := []int{1, 2, 3}
    head := link.BuildListNode(data)
    dummy := link.RotateLinkedList(head, 1000000000)
    link.ShowData(dummy)
}

//
//  TestPlusOne
//  @Description: 链表加1
//  @param t
//
func TestPlusOne(t *testing.T) {
    data := []int{9, 9, 9}
    head := link.BuildListNode(data)
    dummy := link.PlusOne(head)
    link.ShowData(dummy)
}

//
//  TestOddEvenList
//  @Description: NC133 链表的奇偶重排
//  @param t
//
func TestOddEvenList(t *testing.T) {
    data := []int{1, 4, 6, 3, 7}
    head := link.BuildListNode(data)
    dummy := link.OddEvenListII(head)
    link.ShowData(dummy)
}

//
//  TestCopyRandomList
//  @Description:JZ35 复杂链表的复制
//  @param t
//
func TestCopyRandomList(t *testing.T) {
    data := []int{1, 2, 3, 4, 5}
    head := link.BuildRandomListNode(data)
    link.CopyRandomList(head)
}

//
//  TestMergeKLists
//  @Description: NC51 合并k个已排序的链表
//  @param t
//
func TestMergeKLists(t *testing.T) {
    data := [][]int{{1, 2}, {1, 4, 5}, {6}}
    list := make([]*link.ListNode, 0)
    for _, item := range data {
        list = append(list, link.BuildListNode(item))
    }
    dummy := link.MergeKLists(list)
    link.ShowData(dummy)
}
