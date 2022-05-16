package challenge

import trees "awesome/tree"

//
//  InorderSuccessor
//  @Description: 面试题 04.06. 后继者
//  @param root
//  @param p
//  @return *trees.TreeNode
//
func InorderSuccessor(root, p *trees.TreeNode) *trees.TreeNode {
    var successor *trees.TreeNode
    if p.Right != nil {
        successor = p.Right
        for successor.Left != nil {
            successor = successor.Left
        }
        return successor
    }
    node := root
    for node != nil {
        if node.Val > p.Val {
            successor = node
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return successor
}
