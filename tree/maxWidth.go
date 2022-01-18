package main

//
//  widthOfBinaryTree
//  @Description: NC204 二叉树的最大宽度
//  @param root
//  @return int
//
func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := make([]*TreeNode, 0)
    root.Val = 1
    queue = queuePush(queue, root)
    node := &TreeNode{}
    maxWidth := 1
    for size := 0; len(queue) > 0; {
        size = len(queue)
        maxWidth = maxInt(maxWidth, queue[len(queue)-1].Val-queue[0].Val+1)
        for size > 0 {
            node, queue = queueShift(queue)
            if node.Left != nil {
                node.Left.Val = 2 * node.Val
                queue = queuePush(queue, node.Left)
            }
            if node.Right != nil {
                node.Right.Val = 2*node.Val + 1
                queue = queuePush(queue, node.Right)
            }
            size--
        }
    }
    return maxWidth
}
