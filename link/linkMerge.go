package main

func linkMerge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
    node := &ListNode{}
    pHead := node
    for pHead1 != nil && pHead2 != nil {
        if pHead1.Val < pHead2.Val {
            node.Next = pHead1
            pHead1 = pHead1.Next
        } else {
            node.Next = pHead2
            pHead2 = pHead2.Next
        }
        node = node.Next
    }
    if pHead1 == nil {
        node.Next = pHead2
    }
    if pHead2 == nil {
        node.Next = pHead1
    }
    return pHead.Next
}
