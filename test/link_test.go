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
    head := link.MakeListNode(data)
    dummy := link.LinkReverse(head)
    link.ShowData(dummy)
}

//
//  TestDeleteDuplicates
//  @Description: NC25 删除有序链表中重复的元素-I
//  @param t
//
func TestDeleteDuplicates(t *testing.T) {
    data := []int{1, 2, 2}
    head := link.MakeListNode(data)
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
    head := link.MakeListNode(data)
    dummy := link.DeleteDuplicatesII(head)
    link.ShowData(dummy)
}

//
//  TestReverseBetween
//  @Description: NC21 链表内指定区间反转
//  @param t
//
func TestReverseBetween(t *testing.T) {
    data := []int{1, 2, 3, 4, 5}
    head := link.MakeListNode(data)
    dummy := link.ReverseBetween(head, 2, 4)
    link.ShowData(dummy)
}

//
//  TestInsertSortList
//  @Description: NC244 对链表进行插入排序
//  @param t
//
func TestInsertSortList(t *testing.T) {
    data := []int{2, 4, 1}
    head := link.MakeListNode(data)
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
    head := link.MakeListNode(data)
    dummy := link.SwapLinkedPair(head)
    link.ShowData(dummy)
}
