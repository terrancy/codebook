package link

//
//  SwapLinkedPair
//  @Description: NC186 遍历:两两交换链表的节点
//  @param head
//  @return *ListNode
//
func SwapLinkedPair(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Val: 0}
    cur := head
    pre := dummy
    for cur != nil && cur.Next != nil {
        // 交换
        // 两两交换
        next := cur.Next
        cur.Next = next.Next
        next.Next = cur
        // 指针指向
        pre.Next = next
        pre = cur
        // 向前走
        cur = cur.Next
    }
    return dummy.Next
}

//
//  SwapLinkedPairII
//  @Description:
//  @param head
//  @return *ListNode
//
func SwapLinkedPairII(head *ListNode) *ListNode {
    return ReverseKGroup(head, 2)
}

//
//  SwapLinkedPairDG
//  @Description: NC186 递归:两两交换链表的节点
//  @param head
//  @return *ListNode
//
func SwapLinkedPairDG(head *ListNode) *ListNode {
    // 单个节点不交换
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    tail := next.Next
    next.Next = head
    head.Next = SwapLinkedPair(tail)
    return next
}
