package trees

func TreeDepth(pRoot *TreeNode) int {
    // write code here
    if pRoot == nil {
        return 0
    }
    queue := make([]*TreeNode, 0)
    queue = queuePush(queue, pRoot)
    size := 0
    level := 0
    node := &TreeNode{}
    for len(queue) > 0 {
        size = len(queue)
        for i := 0; i < size; i++ {
            node, queue = queueShift(queue)
            if node.Left != nil {
                queue = queuePush(queue, node.Left)
            }
            if node.Right != nil {
                queue = queuePush(queue, node.Right)
            }
        }
        level++
    }
    return level
}

func PrintFromTopToBottom(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    queue := make([]*TreeNode, 0)
    data := make([]int, 0)
    node := &TreeNode{}
    size := 0
    queue = queuePush(queue, root)
    for len(queue) > 0 {
        size = len(queue)
        for size > 0 {
            node, queue = queueShift(queue)
            data = append(data, node.Val)
            if node.Left != nil {
                queue = queuePush(queue, node.Left)
            }
            if node.Right != nil {
                queue = queuePush(queue, node.Right)
            }
            size--
        }
    }
    return data
}
