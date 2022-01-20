package link

import "sort"

//
//  sortInList
//  @Description: NC70 单链表的排序
//  @param head
//  @return *ListNode
//
func SortInList(head *ListNode) *ListNode {
    data := make([]int, 0)
    for head != nil {
        data = append(data, head.Val)
        head = head.Next
    }
    sort.Ints(data)
    
    node := &ListNode{data[0], nil}
    cur := node
    n := len(data)
    for i := 1; i < n; i++ {
        temp := &ListNode{data[i], nil}
        cur.Next = temp
        cur = cur.Next
    }
    return node
}
