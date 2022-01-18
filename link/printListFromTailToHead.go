package main

func printListFromTailToHead(head *ListNode) []int {
    // write code here
    cur := head
    data := make([]int, 0)
    for cur != nil {
        data = append(data, cur.Val)
        cur = cur.Next
    }
    arrayReverse(data)
    return data
}
