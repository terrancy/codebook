package main

import (
    "encoding/json"
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func makeListNode(data []int) *ListNode {
    head := &ListNode{}
    n := len(data)
    if n == 0 {
        return head
    }
    cur := head
    for _, val := range data {
        temp := &ListNode{val, nil}
        cur.Next = temp
        cur = cur.Next
    }
    return head.Next
}

//
//  circleSingleListNode
//  @Description: 环形单向链表
//  @param n
//  @return *ListNode
//
func circleSingleListNode(n int) *ListNode {
    head := &ListNode{}
    cur := head
    temp := &ListNode{}
    for i := 1; i <= n; i++ {
        temp = &ListNode{i, nil}
        cur.Next = temp
        cur = cur.Next
    }
    cur.Next = head.Next
    return head.Next
}

//
//  list2Array
//  @Description: 将链表转为数组
//  @param head
//  @return []int
//
func list2Array(head *ListNode) []int {
    if head == nil {
        return nil
    }
    
    cur := head
    data := make([]int, 0)
    for cur != nil {
        data = append(data, cur.Val)
        cur = cur.Next
    }
    
    return data
}

//
//  showData
//  @Description: 将链表数据打印出来
//  @param head
//
func showData(head *ListNode, isPrint bool) []int {
    data := list2Array(head)
    if isPrint == false {
        return data
    }
    if len(data) == 0 {
        data = []int{}
    }
    json, err := json.Marshal(data)
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Println(string(json))
    return nil
}
