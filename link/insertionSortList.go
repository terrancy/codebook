package link

//
//  InsertionSortList
//  @Description: NC244 对链表进行插入排序
//  @Solution: 在循环体内，dummy是一个有序的链表，将target节点追加到链表的合适位置.
//  @param head
//  @return *ListNode
//
func InsertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Val: -10000}
    target := head
    for target != nil {
        pre := dummy
        for pre.Next != nil && pre.Next.Val < target.Val {
            pre = pre.Next
        }
        temp := target.Next
        next := pre.Next
        pre.Next = target
        target.Next = next
        target = temp
    }
    return dummy.Next
}
