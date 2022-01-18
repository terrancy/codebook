package main

//
//  FindFirstCommonNode
//  @Description: NC66 两个链表的第一个公共结点
//  @param pHead1
//  @param pHead2
//  @return *ListNode
//
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
    if pHead1 == nil || pHead2 == nil {
        return nil
    }
    cur1, cur2 := pHead1, pHead2
    for cur1 != cur2 {
        if cur1 == nil {
            cur1 = pHead2
        } else {
            cur1 = cur1.Next
        }
        if cur2 == nil {
            cur2 = pHead1
        } else {
            cur2 = cur2.Next
        }
    }
    return cur1
}
