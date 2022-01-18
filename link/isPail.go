package link

//
//  isPail
//  @Description: NC96 判断一个链表是否为回文结构
//  @solution: 链表翻转 + 遍历对比
//  @param head
//  @return bool
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

//func main() {
//    data := []int{1, 2, 3, 2, 1}
//    head := makeListNode(data)
//    isPail := isPail(head)
//    println(isPail)
//}
