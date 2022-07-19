package trees

//
// GetNext
// @Description: JZ8 二叉树的下一个结点
// @Requirement: 要求：空间复杂度 O(1)  ，时间复杂度 O(n)
// @param pNode
// @return *TreeLinkNode
//
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
    if pNode == nil {
        return nil
    }
    cur := pNode
    for cur != nil {
        cur = cur.Next
    }
    return cur
}

//
// inOrderGetNext
// @Description: 中根遍历获取目标值
// @param root
// @param target
// @return *TreeLinkNode
//
func inOrderGetNext(root *TreeLinkNode, target *TreeLinkNode) *TreeLinkNode {
    return root
}

//
// GetNextII
// @Description: 中根遍历获取目标值
// @param pNode
// @return *TreeLinkNode
//
func GetNextII(pNode *TreeLinkNode) *TreeLinkNode {
    if pNode == nil {
        return nil
    }
    // 如果当前结点的右子树不为空
    if pNode.Right != nil {
        pNode = pNode.Right
        for pNode.Left != nil {
            pNode = pNode.Left
        }
        return pNode
    }
    
    // 当前结点无右子树时,可能是父节点的左子树或者右子树
    for pNode.Next != nil {
        if pNode.Next.Left == pNode {
            return pNode.Next
        }
        pNode = pNode.Next
    }
    return nil
}
