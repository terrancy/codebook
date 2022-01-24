package link

//
//  reverseBetween
//  @Description: NC21 链表内指定区间反转
//  @param head
//  @param m
//  @param n
//  @return *ListNode
//
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
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
    for i := 0; i < n-m; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    
    return dummyNode.Next
}

//
//  ReverseBetweenII
//  @Description: 区间翻转
//  @Description: 利用翻转链表实现,节点从1开始,构建虚节点,翻转链表Next
//  @param head
//  @param m
//  @param n
//  @return *ListNode
//
func ReverseBetweenII(head *ListNode, m int, n int) *ListNode {
    if head == nil || head.Next == nil || m == n {
        return head
    }
    // 获取区间头尾节点
    dummy := &ListNode{}
    dummy.Next = head
    tail, pre := dummy, dummy
    for i := 0; i < n; i++ {
        if i < m-1 {
            pre = pre.Next
        }
        tail = tail.Next
    }
    ReverseBetweenNext(pre, tail)
    return dummy.Next
}

//
//  ReverseBetweenNext
//  @Description: 根据链表的头尾指针翻转链表
//  @param pre
//  @param tail
//
func ReverseBetweenNext(pre *ListNode, tail *ListNode) *ListNode {
    if pre.Next == tail {
        return tail
    }
    // 截取
    cur := pre.Next
    next := tail.Next
    tail.Next = nil
    // 翻转合并
    pre.Next = Reverse(cur)
    cur.Next = next
    return cur
}
