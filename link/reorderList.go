package main

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    // 翻转链表
    reverseLastPart(head)
    //mid := reverseLastPart(head)
    //cur := head
    //curNext := &ListNode{}
    //midNext := &ListNode{}
    //for mid != nil {
    //    curNext = cur.Next
    //    cur.Next = mid
    //    midNext = mid.Next
    //    mid.Next = curNext
    //
    //    // 遍历
    //    mid = midNext
    //    cur = curNext
    //}
}

func reverseLastPart(head *ListNode) *ListNode {
    // 链表中点
    cur := head
    pre := &ListNode{}
    pre.Next = head
    for cur != nil && cur.Next != nil {
        pre = pre.Next
        cur = cur.Next.Next
    }
    mid := pre.Next
    // 翻转链表-四联式
    //cur = pre.Next
    //next := &ListNode{}
    //for pre.Next != nil && cur.Next.Next != nil {
    //   next = cur.Next
    //   cur.Next = next.Next
    //   next.Next = pre.Next
    //   pre.Next = next
    //}
    
    return mid
}

func reverseLastPartII(head *ListNode) *ListNode {
    cur := head
    cnt := 0
    for cur != nil {
        cur = cur.Next
        cnt++
    }
    half := 1 + (cnt+1)>>1
    head = reverseBetween(head, half, cnt)
    return head
}

func reorderListII(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    data := make([]*ListNode, 0)
    cur := head
    for cur != nil {
        data = append(data, cur)
        cur = cur.Next
    }
    
    i, j := 0, len(data)-1
    for i < j {
        data[i].Next = data[j]
        i++
        if i == j {
            break
        }
        data[j].Next = data[i]
        j--
    }
    data[i].Next = nil
}

//func main() {
//    data := []int{1, 2, 3, 4}
//    head := makeListNode(data)
//    reorderListII(head)
//    showData(head, true)
//}
