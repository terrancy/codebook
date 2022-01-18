package link

import "sort"

//
//  spiralOrder
//  @Description: NC51 合并k个已排序的链表
//  @param lists
//  @return *ListNode
//
func spiralOrder(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    
    data := make([]int, 0)
    for _, head := range lists {
        for head != nil {
            data = append(data, head.Val)
            head = head.Next
        }
    }
    sort.Ints(data)
    m := len(data)
    head := &ListNode{data[0], nil}
    cur := head
    for i := 1; i < m; i++ {
        cur.Next = &ListNode{data[i], nil}
        cur = cur.Next
    }
    
    return head
}

//func main() {
//    data := [][]int{{1, 2, 3,}, {4, 5, 6, 7,}}
//    list := make([]*ListNode, 0)
//    for _, val := range data {
//        list = append(list, makeListNode(val))
//    }
//    spiralOrder(list)
//}
