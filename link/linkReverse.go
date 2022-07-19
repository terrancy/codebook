package link

//
// Reverse
// @Description: NC78 反转链表
// @param head
// @return *ListNode
//
func Reverse(head *ListNode) *ListNode {
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

//
// HalfReverse
// @Description:
// @param head
// @return *ListNode
//
func HalfReverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    // 构建虚节点
    dummy := &ListNode{Next: head}
    // 确认节点
    fast, slow := dummy.Next, dummy.Next
    cnt := 0
    for fast != nil {
        fast = fast.Next
        cnt++
    }
    for i := 1; i < (cnt+1)/2; i++ {
        slow = slow.Next
    }
    // 截取
    pre := slow
    cur := slow.Next
    // 翻转合并
    pre.Next = Reverse(cur)
    return dummy.Next
}
