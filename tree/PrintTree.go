package main

func PrintTree(pRoot *TreeNode) [][]int {
    if pRoot == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    queue = queuePush(queue, pRoot)
    size := 0
    node := &TreeNode{}
    data := make([][]int, 0)
    for len(queue) > 0 {
        size = len(queue)
        tmp := make([]int, 0)
        for i := 0; i < size; i++ {
            node, queue = queueShift(queue)
            tmp = append(tmp, node.Val)
            if node.Left != nil {
                queue = queuePush(queue, node.Left)
            }
            if node.Right != nil {
                queue = queuePush(queue, node.Right)
            }
        }
        data = append(data, tmp)
    }
    return data
}

