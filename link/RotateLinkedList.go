package link

//
//  RotateLinkedList
//  @Description: NC211 旋转链表
//  @param head
//  @param k
//  @return *ListNode
//
func RotateLinkedList(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 取模
    cnt := 0
    cur := head
    for cur != nil {
        cnt++
        cur = cur.Next
    }
    k = k % cnt
    if k == 0 {
        return head
    }
    // 确定移动k位后的头尾节点,确定倒数第K+1个节点
    fast := head
    for i := 0; fast != nil && i < k+1; i++ {
        fast = fast.Next
    }
    slow := head
    for fast != nil {
        if fast.Next == nil {
            fast.Next = head
            slow = slow.Next
            break
        }
        fast = fast.Next
        slow = slow.Next
    }
    next := slow.Next
    slow.Next = nil
    return next
}
