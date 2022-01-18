package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//
//  makeTreeNode
//  @Description: 通过数组层序遍历构建"满二叉树"
//  @param data
//  @return *TreeNode
//
func makeTreeNode(data []int) *TreeNode {
    // 数组优化
    data = dataOptimize(data)
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
func dataOptimize(data []int) []int {
    n := len(data)
    if n&1 == 0 {
        data = append(data, INF)
    }
    return data
}

func queuePop(queue []*TreeNode) (*TreeNode, []*TreeNode) {
    node := queue[len(queue)-1]
    queue = queue[:len(queue)-1]
    return node, queue
}

func queueShift(queue []*TreeNode) (*TreeNode, []*TreeNode) {
    node := queue[0]
    queue = queue[1:]
    return node, queue
}

func queuePush(queue []*TreeNode, node *TreeNode) []*TreeNode {
    queue = append(queue, node)
    return queue
}
