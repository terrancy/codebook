package link

//
//  oddEvenList
//  @Description: NC133 链表的奇偶重排
//  @param head
//  @return *ListNode
//
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    evenHead := head.Next
    odd, even := head, head.Next
    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head
}

func OddEvenListII(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    oddHead, evenHead := head, head.Next
    odd, even := oddHead, evenHead
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
    odd.Next = evenHead
    return oddHead
}
