package main

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
    k := maxInt(n, m)
    data := make([]int, k)
    sum, plus := 0, 0
    a, b := 0, 0
    for i := 0; i < k; i++ {
        sum = 0
        if i < n {
            a = data1[n-1-i]
        } else {
            a = 0
        }
        if i < m {
            b = data2[m-1-i]
        } else {
            b = 0
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

//func dataSwag(data1, data2 []int) ([]int, []int) {
//    n, m := len(data1), len(data2)
//    temp := make([]int, 0)
//    if n < m {
//        temp = data1
//        data1 = data2
//        data2 = temp
//    }
//    return data1, data2
//}

//func main() {
//    data1 := []int{9, 3, 7}
//    node1 := makeListNode(data1)
//    data2 := []int{6, 3}
//    node2 := makeListNode(data2)
//    data := addInList(node1, node2)
//    showData(data, true)
//}
