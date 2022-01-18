package main

import "fmt"

const INF = 10001

//func main() {
//    data := []int{1, 2, 3, 10001, 10001, 6, 7}
//    root := makeTreeNode(data)
//    temp := serialize(root)
//    fmt.Println(temp)
//    makeTreeNode(temp)
//}

//func main() {
//    data := []int{4, 2, 5, 3, 1}
//    root := makeTreeNode(data)
//    res := findErrorSBT(root)
//    fmt.Println(res)
//}

//func main() {
//    data := []int{7, 1, 12, 0, 4, 11, 14, INF, INF, 3, 5}
//    root := makeTreeNode(data)
//    res := lowestCommonAncestor(root, 12, 11)
//    fmt.Println(res)
//}

// 和为某一个值的路径的问题
func main() {
    data := []int{1, 2, 3, 4, 5, 4, 3, INF, INF, -1}
    root := makeTreeNode(data)
    res := findPathIII(root, 6)
    fmt.Println(res)
}
