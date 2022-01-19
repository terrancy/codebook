package link

//
//  InsertionSortList
//  @Description: NC244 对链表进行插入排序
//  @Solution: 两个指针,一个指向表头,一个待插入指针.从表头开始找第一个比他大的位置放置在他的前面(next.val > cur.val)
//  @param head
//  @return *ListNode
//
func InsertionSortList(head *ListNode) *ListNode {
    // 定义头结点 找插入的位置 插入操作
    dummy := &ListNode{Val: -10000}
    target := head
    for target != nil {
        pre := dummy
        for pre.Next != nil && pre.Next.Val < target.Val {
            pre = pre.Next
        }
        cur := pre.Next
        pre.Next = target
        target = target.Next
        pre.Next.Next = cur
    }
    return dummy.Next
}
