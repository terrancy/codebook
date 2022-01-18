package main

//
//  reverseBetween
//  @Description: NC21 链表内指定区间反转
//  @param head
//  @param m
//  @param n
//  @return *ListNode
//
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    // 构建虚节点
    dummyNode := &ListNode{}
    dummyNode.Next = head
    
    // 前置节点
    pre := dummyNode
    for i := 0; i < m-1; i++ {
        pre = pre.Next
    }
    
    // 翻转四联式
    cur := pre.Next
    next := &ListNode{}
    for i := 0; i < n-m; i++ {
        next = cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    
    return dummyNode.Next
}
