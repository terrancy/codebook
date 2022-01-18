package link

//
//  removeNthFromEnd
//  @Description: NC53 删除链表的倒数第n个节点
//  @param pHead
//  @param n
//  @return *ListNode
//
func removeNthFromEnd(pHead *ListNode, n int) *ListNode {
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
