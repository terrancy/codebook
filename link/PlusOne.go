package link

func PlusOne(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    head = plusOneReverse(head)
    cur := head
    flag := 0
    plus := 1
    for cur != nil {
        sum := cur.Val + plus + flag
        flag = int(sum / 10)
        cur.Val = sum % 10
        if cur.Next == nil && flag == 1 {
            cur.Next = &ListNode{Val: 1}
            flag = 0
        }
        plus = 0
        cur = cur.Next
    }
    
    return plusOneReverse(head)
}

func plusOneReverse(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
