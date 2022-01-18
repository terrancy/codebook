package main

//
//  EntryNodeOfLoop
//  @Description: NC3 链表中环的入口结点
//  @Solution: 快慢指针相遇(快=2慢),快从头走,慢继续走直到相遇.
//  @Warn: 无环和最大环的判断
//  @param pHead
//  @return *ListNode
//
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
    fast := pHead
    slow := pHead
    for ; fast != nil && fast.Next != nil; {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            break
        }
    }
    // 无环
    if fast == nil || fast.Next == nil {
        return nil
    }
    fast = pHead
    for ; fast != slow; {
        fast = fast.Next
        slow = slow.Next
    }
    return fast
}
