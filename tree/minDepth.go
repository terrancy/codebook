package trees

//
//  minDepth
//  @Description: NC234 二叉树的最小深度
//  @param root
//  @return int
//
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := make([]*TreeNode, 0)
    queue = queuePush(queue, root)
    node := &TreeNode{}
    level := 0
    for size := 0; len(queue) > 0; {
        size = len(queue)
        level++
        for size > 0 {
            node, queue = queueShift(queue)
            if node.Left != nil {
                queue = queuePush(queue, node.Left)
            }
            if node.Right != nil {
                queue = queuePush(queue, node.Right)
            }
            if node.Left == nil && node.Right == nil {
                return level
            }
            size--
        }
    }
    return level
}
