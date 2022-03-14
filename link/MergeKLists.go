package link

//
//  MergeKLists
//  @Description: 23. 合并K个升序链表 / NC51 合并k个已排序的链表
//  @param lists
//  @return *ListNode
//
func MergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    // 定义虚节点!!
    var dummy *ListNode
    for _, head := range lists {
        dummy = mergeList(dummy, head)
    }
    return dummy
}

func mergeList(head1, head2 *ListNode) *ListNode {
    if head1 == nil {
        return head2
    }
    if head2 == nil {
        return head1
    }
    dummy := &ListNode{}
    pre := dummy
    cur1, cur2 := head1, head2
    for cur1 != nil && cur2 != nil {
        if cur1.Val < cur2.Val {
            pre.Next = cur1
            cur1 = cur1.Next
        } else {
            pre.Next = cur2
            cur2 = cur2.Next
        }
        pre = pre.Next
    }
    if cur1 != nil {
        pre.Next = cur1
    }
    if cur2 != nil {
        pre.Next = cur2
    }
    return dummy.Next
}
