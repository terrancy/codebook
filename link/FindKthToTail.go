package link

//
// FindKthToTail
// @Description: NC69 链表中倒数最后k个结点
// @param pHead
// @param k
// @return *ListNode
//
func FindKthToTail(pHead *ListNode, k int) *ListNode {
    // 快指针走K步
    fast := pHead
    for i := 0; i < k; i++ {
        if fast == nil {
            return nil
        }
        fast = fast.Next
    }
    
    // 快指针走到底
    slow := pHead
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    
    return slow
}

func FindKthToTailII(pHead *ListNode, k int) *ListNode {
    if pHead == nil {
        return nil
    }
    
    cur := pHead
    data := make([]*ListNode, 0)
    for cur != nil {
        data = append(data, cur)
        cur = cur.Next
    }
    
    n := len(data)
    if n < k {
        return nil
    }
    
    return data[k]
}
