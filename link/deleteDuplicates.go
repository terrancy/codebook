package link

//
// deleteDuplicates
// @Description: NC25 删除有序链表中重复的元素-I
// @param head
// @return *ListNode
//
func DeleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    cur := head
    for cur != nil && cur.Next != nil {
        // 删除重复值
        if cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
            continue
        }
        cur = cur.Next
    }
    return head
}

//
// deleteDuplicatesII
// @Description: NC24 删除有序链表中重复的元素-II
// @Solution: 1. 设置标志位 2.前驱指针 3.前驱节点 4.标志位判断
// @param head
// @return *ListNode
//
func DeleteDuplicatesII(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Next: head}
    pre := dummy
    cur := head
    isDeleted := false
    for cur.Next != nil {
        if pre.Next.Val == cur.Next.Val {
            isDeleted = true
        } else if isDeleted == true {
            pre.Next = cur.Next
            isDeleted = false
        } else {
            pre = cur
        }
        cur = cur.Next
    }
    // 存在最后两个相等直接退出的情况
    if isDeleted == true {
        pre.Next = nil
    }
    return dummy.Next
}
