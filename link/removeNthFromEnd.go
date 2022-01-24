package link

//
//  removeNthFromEnd
//  @Description: NC53 删除链表的倒数第n个节点
//  @param pHead
//  @param n
//  @return *ListNode
//
func RemoveNthFromEnd(pHead *ListNode, n int) *ListNode {
    if pHead == nil || pHead.Next == nil {
        return nil
    }
    // 倒数第K+1个节点,待删除节点的前一个节点
    pKthNode := FindKthToTail(pHead, n+1)
    if pKthNode == nil {
        pHead = pHead.Next
    } else {
        pKthNode.Next = pKthNode.Next.Next
    }
    return pHead
}

func RemoveNthFromEndII(pHead *ListNode, n int) *ListNode {
    if pHead == nil || pHead.Next == nil {
        return pHead
    }
    cur := pHead
    for i := 0; i < n; i++ {
        if cur == nil {
            return nil
        }
        cur = cur.Next
    }
    dummy := &ListNode{Next: pHead}
    pre := dummy
    for cur != nil {
        cur = cur.Next
        pre = pre.Next
    }
    pre.Next = pre.Next.Next
    return dummy.Next
}
