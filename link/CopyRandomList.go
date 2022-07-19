package link

//
// CopyRandomList
// @Description: JZ35 复杂链表的复制
// @param head
// @return *RandomListNode
//
func CopyRandomList(head *RandomListNode) *RandomListNode {
    if head == nil {
        return nil
    }
    // 1.链表复制
    cur := head
    for cur != nil {
        cloneHead := &RandomListNode{Label: cur.Label}
        cloneHead.Next = cur.Next
        cur.Next = cloneHead
        cur = cur.Next.Next
    }
    // 2.随机指针处理
    cur = head
    for cur != nil {
        if cur.Random != nil {
            cur.Next.Random = cur.Random.Next
        }
        cur = cur.Next.Next
    }
    // 3.链表拆分
    oddHead, evenHead := &RandomListNode{}, &RandomListNode{}
    odd, even := oddHead, evenHead
    cur = head
    for i := 1; cur != nil; i++ {
        if i&1 == 1 {
            odd.Next = cur
            odd = odd.Next
        } else {
            even.Next = cur
            even = even.Next
        }
        cur = cur.Next
    }
    odd.Next = nil
    even.Next = nil
    return evenHead.Next
}
