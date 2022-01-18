package link

//
//  partition
//  @Description: NC23 划分链表
//  @param head
//  @param x
//  @return *ListNode
func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    // 生成前置后置两个虚节点
    cur := head
    pre := &ListNode{}
    preHead := pre
    next := &ListNode{}
    nextHead := next
    for cur != nil {
        if cur.Val < x {
            pre.Next = cur
            pre = pre.Next
        } else {
            next.Next = cur
            next = next.Next
        }
        cur = cur.Next
    }
    // 后直节点要关闭,避免容易死循环
    next.Next = nil
    
    // 前置后置拼接
    preHead.Next = nextHead.Next
    
    return preHead.Next
}

//func main() {
//    data := []int{1, 4, 3, 2, 5, 2}
//    head := makeListNode(data)
//    //partition(head, 3)
//    showData(head)
//}
