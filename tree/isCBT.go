package trees
//
//  IsCompleteTree
//  @Description: 判断是否是完全二叉树
//  @param root
//  @return bool
//
func IsCompleteTree(root *TreeNode) bool {
    if root == nil {
        return false
    }
    return isCBT(root)
}

func isCBT(root *TreeNode) bool {
    queue := []*TreeNode{root}
    node := &TreeNode{}
    isFull := false
    for len(queue) > 0 {
        node, queue = queueShift(queue)
        if node.Left != nil {
            queue = queuePush(queue, node.Left)
            if isFull {
                return false
            }
        } else {
            isFull = true
        }
        if node.Right != nil {
            queue = queuePush(queue, node.Right)
            if isFull {
                return false
            }
        } else {
            isFull = true
        }
    }
    return true
}
