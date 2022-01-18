package main

//
//  nodeNum
//  @Description: NC84 完全二叉树节点数
//  @param head
//  @return int
//
func nodeNum(head *TreeNode) int {
    return countNode(head)
}

func countNode(head *TreeNode) int {
    if head == nil {
        return 0
    }
    return 1 + countNode(head.Left) + countNode(head.Right)
}
