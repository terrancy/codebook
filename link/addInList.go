package link

import "awesome"

//
//  addInList
//  @Description:
//  @param head1
//  @param head2
//  @return *ListNode
//
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
    if head1 == nil && head2 == nil {
        return nil
    }
    if head1 == nil && head2 != nil {
        return head2
    }
    if head1 != nil && head2 == nil {
        return head1
    }
    
    cur1 := head1
    data1 := make([]int, 0)
    for cur1 != nil {
        data1 = append(data1, cur1.Val)
        cur1 = cur1.Next
    }
    
    cur2 := head2
    data2 := make([]int, 0)
    for cur2 != nil {
        data2 = append(data2, cur2.Val)
        cur2 = cur2.Next
    }
    data := towSum(data1, data2)
    dummyNode := &ListNode{}
    temp := &ListNode{}
    cur := dummyNode
    for i := len(data) - 1; i >= 0; i-- {
        temp = &ListNode{data[i], nil}
        cur.Next = temp
        cur = cur.Next
    }
    
    return dummyNode.Next
}

//
//  towSum
//  @Description: 大数之和 - 模拟加法
//  @param data1
//  @param data2
//  @return []int
//
func towSum(data1, data2 []int) []int {
    n, m := len(data1), len(data2)
    k := awesome.MaxInt(n, m)
    data := make([]int, k)
    plus := 0
    for i := 0; i < k; i++ {
        a, b, sum := 0, 0, 0
        if i < n {
            a = data1[n-1-i]
        }
        if i < m {
            b = data2[m-1-i]
        }
        sum = plus + a + b
        plus = int(sum / 10)
        data[i] = sum % 10
    }
    if plus == 1 {
        data = append(data, plus)
    }
    return data
}
