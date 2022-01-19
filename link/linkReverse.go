package link

func LinkReverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
