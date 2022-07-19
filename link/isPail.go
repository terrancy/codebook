package link

//
// isPail
// @Description: NC96 判断一个链表是否为回文结构
// @solution: 链表翻转 + 遍历对比
// @param head
// @return bool
//
func isPail(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    // 折半翻转
    head = reverseLastPartII(head)
    fast := head
    slow := head
    cnt := 0
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        cnt++
    }
    if fast != nil {
        slow = slow.Next
    }
    cur := head
    for slow != nil {
        if slow.Val != cur.Val {
            return false
        }
        slow = slow.Next
        cur = cur.Next
    }
    return true
}

func isPailII(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    if head.Next == nil {
        return true
    }
    
    slow := head
    fast := head
    data := make([]int, 0)
    for fast != nil && fast.Next != nil {
        data = append(data, slow.Val)
        fast = fast.Next.Next
        slow = slow.Next
    }
    if fast != nil {
        slow = slow.Next
    }
    cnt := len(data)
    for slow != nil {
        if slow.Val != data[cnt-1] {
            return false
        }
        cnt--
        slow = slow.Next
    }
    return true
}

//
// isPailIII
// @Description: NC96 判断一个链表是否为回文结构
// @Solution: 1. 获取前驱节点 2. 链表翻转 3.数据对比
// @Warning: 1. pre := dummy 2. if cur!=nil pre = pre.Next
// @param head
// @return bool
//
func isPailIII(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    // pre 前驱节点
    dummy := &ListNode{Next: head}
    pre, cur := dummy, head
    for cur != nil && cur.Next != nil {
        pre = pre.Next
        cur = cur.Next.Next
    }
    // 细节
    if cur != nil {
        pre = pre.Next
    }
    // 翻转
    pre.Next = Reverse(pre.Next)
    // 回文判断
    left, right := head, pre.Next
    for right != nil {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }
    return true
}
