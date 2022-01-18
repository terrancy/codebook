package main

//
//  ysfLink
//  @Description: 约瑟夫环链表
//  @param n 1~n人组成圈
//  @param m 报数步长step
//
func ysfLink(n, m int) int {
    if n == 1 {
        return 1
    }
    if m == 1 {
        return n
    }
    // 构建环形单向链表
    head := circleSingleListNode(n)
    cur := head
    temp := head
    for cur != cur.Next {
        temp = cur
        for i := 0; i < m-1; i++ {
            cur = cur.Next
        }
        if cur.Next == cur {
            return cur.Val
        }
        temp.Next = cur.Next
    }
    return cur.Val
}

func ysfII(n, m int) int {
    if n == 1 {
        return 1
    }
    if m == 1 {
        return n
    }
    ans := 0
    for i := 2; i <= n; i++ {
        ans = (ans + m) % i
    }
    return ans + 1
}

//func main() {
//    res := ysfLink(3, 1)
//    println(res)
//}
