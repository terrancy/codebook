package test

import (
    "awesome/link"
    "fmt"
    "testing"
)

//
//  TestDeleteDuplicates
//  @Description: NC25 删除有序链表中重复的元素-I
//  @param t
//
func TestDeleteDuplicates(t *testing.T) {
    data := []int{1, 2, 2}
    head := link.MakeListNode(data)
    node := link.DeleteDuplicates(head)
    tmp := link.Serialize(node)
    fmt.Print(tmp)
}

//
//  TestDeleteDuplicatesII
//  @Description: NC24 删除有序链表中重复的元素-II
//  @param t
//
func TestDeleteDuplicatesII(t *testing.T) {
    data := []int{1, 2, 2}
    head := link.MakeListNode(data)
    node := link.DeleteDuplicatesII(head)
    tmp := link.Serialize(node)
    fmt.Print(tmp)
}

//
//  TestReverseBetween
//  @Description: NC21 链表内指定区间反转
//  @param t
//
func TestReverseBetween(t *testing.T) {
    data := []int{1, 2, 3, 4, 5}
    head := link.MakeListNode(data)
    node := link.ReverseBetween(head, 2, 4)
    tmp := link.Serialize(node)
    fmt.Println(tmp)
}
