package link

//
// ReverseKGroup
// @Description: NC50 链表中的节点每k个一组翻转
// @param head
// @param k
// @return *ListNode
//
func ReverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 1 {
        return head
    }
    
    cur := head
    n := 0
    for cur != nil {
        n++
        cur = cur.Next
    }
    if n < k {
        return head
    }
    
    cur = head
    for i := 0; i < n/k; i++ {
        cur = ReverseBetweenII(cur, i*k+1, (i+1)*k)
    }
    return cur
}

//
// ReverseKGroupII
// @Description: NC50 链表中的节点每k个一组翻转
// @param head
// @param k
// @return *ListNode
//
func ReverseKGroupII(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 1 {
        return head
    }
    cur := head
    n := 0
    for cur != nil {
        n++
        cur = cur.Next
    }
    if n < k {
        return head
    }
    dummy := &ListNode{Next: head}
    pre := dummy
    cur = pre
    m := n - n%k
    for i := 1; i <= m; i++ {
        cur = cur.Next
        if i%k == 0 {
            tail := cur
            cur = ReverseBetweenNext(pre, tail)
            pre = cur
        }
    }
    return dummy.Next
}
