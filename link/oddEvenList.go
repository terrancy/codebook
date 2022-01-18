package main

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
