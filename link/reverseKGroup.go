package link

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    cur := head
    n := 0
    for cur != nil {
        n++
        cur = cur.Next
    }
    if n < k {
        return head
    }
    
    cur = head
    for i := 0; n >= k; i++ {
        cur = reverseBetween(cur, i*k+1, (i+1)*k)
        n = n - k
    }
    return cur
}

//func main() {
//    data := []int{1, 2, 3, 4, 5}
//    head := makeListNode(data)
//    head = reverseKGroup(head, 2)
//    showData(head, true)
//}
