package main

//
//  deleteDuplicates
//  @Description: NC25 删除有序链表中重复的元素-I
//  @param head
//  @return *ListNode
//
func deleteDuplicates(head *ListNode) *ListNode {
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
//  deleteDuplicatesII
//  @Description: NC24 删除有序链表中重复的元素-II
//  @param head
//  @return *ListNode
//
func deleteDuplicatesII(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    preHead := &ListNode{}
    preHead.Next = head
    pre := preHead
    cur := head
    isDeleted := false
    for cur != nil && cur.Next != nil {
        if cur.Val == cur.Next.Val {
            cur.Next = cur.Next.Next
            isDeleted = true
        } else {
            if isDeleted == true {
                pre.Next = cur.Next
                isDeleted = false
            } else {
                pre = cur
            }
            cur = cur.Next
        }
    }
    // 存在最后两个相等直接退出的情况
    if isDeleted == true {
        pre.Next = nil
    }
    return preHead.Next
}
