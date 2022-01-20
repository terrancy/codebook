package link

//
//  SortLinkedList
//  @Description: NC207 排序奇升偶降链表
//  @Description: 给定一个奇数位升序，偶数位降序的链表，返回对其排序后的链表
//  @param head
//  @return *ListNode
//
func SortLinkedList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 拆分
    odd, even := getOddEvenLink(head)
    // 翻转
    even = sortLinkedListReverse(even)
    // 合并
    return sortLinkMerge(odd, even)
}

//
//  sortLinkedListReverse
//  @Description: 翻转
//  @param head
//  @return *ListNode
//
func sortLinkedListReverse(head *ListNode) *ListNode {
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
//  getOddEvenLink
//  @Description: 拆分
//  @param head
//  @return *ListNode
//  @return *ListNode
//
func getOddEvenLink(head *ListNode) (*ListNode, *ListNode) {
    odd, oddHead := head, head
    even, evenHead := head.Next, head.Next
    cur := head.Next.Next
    for i := 1; cur != nil; i++ {
        if i&1 == 1 {
            odd.Next = cur
            odd = odd.Next
        } else {
            even.Next = cur
            even = even.Next
        }
        cur = cur.Next
    }
    odd.Next = nil
    even.Next = nil
    return oddHead, evenHead
}

//
//  sortLinkMerge
//  @Description: 合并
//  @param odd
//  @param even
//  @return *ListNode
//
func sortLinkMerge(odd *ListNode, even *ListNode) *ListNode {
    if odd == nil {
        return even
    }
    if even == nil {
        return odd
    }
    
    dummy := &ListNode{}
    cur := dummy
    for odd != nil && even != nil {
        if odd.Val > even.Val {
            cur.Next = even
            even = even.Next
        } else {
            cur.Next = odd
            odd = odd.Next
        }
        cur = cur.Next
    }
    // 剩下部分
    if odd != nil {
        cur.Next = odd
    }
    if even != nil {
        cur.Next = even
    }
    return dummy.Next
}
