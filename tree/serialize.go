package main

import (
    "strconv"
    "strings"
)

//
//  serialize
//  @Description: 二叉树序列化
//  @param root
//  @return string
//
func serialize(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    data := make([]int, 0)
    const INF = 10001
    nullNode := &TreeNode{INF, nil, nil}
    for i := 0; i < len(queue); i++ {
        node := queue[i]
        data = append(data, node.Val)
        if node == nullNode {
            continue
        }
        if node.Left != nil {
            queue = queuePush(queue, node.Left)
        } else {
            queue = queuePush(queue, nullNode)
        }
        if node.Right != nil {
            queue = queuePush(queue, node.Right)
        } else {
            queue = queuePush(queue, nullNode)
        }
    }
    return data
}

func array2String(data []int) string {
    n := len(data)
    str := make([]string, n)
    for i := 0; i < n; i++ {
        str[i] = strconv.Itoa(data[i])
    }
    return strings.Join(str, ",")
}

//
//  deserialize
//  @Description: 二叉树反序列化
//  @param str
//  @return *TreeNode
//
func deserialize(str string) *TreeNode {
    data := string2Array(str)
    n := len(data)
    if n == 0 {
        return &TreeNode{}
    }
    // 极大值 根节点 队列存储
    const INF = 10001
    root := &TreeNode{data[0], nil, nil}
    node := &TreeNode{}
    queue := []*TreeNode{root}
    // 遍历数组 队列出根节点 设置孩子节点
    for i := 1; i < n-1; i = i + 2 {
        node, queue = queueShift(queue)
        a, b := data[i], data[i+1]
        if a != INF {
            node.Left = &TreeNode{a, nil, nil}
            queue = queuePush(queue, node.Left)
        }
        if b != INF {
            node.Right = &TreeNode{b, nil, nil}
            queue = queuePush(queue, node.Right)
        }
    }
    return root
}

func string2Array(str string) []int {
    var strData []string
    strData = strings.Split(str, ",")
    n := len(strData)
    data := make([]int, n)
    for i := 0; i < n; i++ {
        data[i], _ = strconv.Atoi(strData[i])
    }
    return data
}
