package trees

//
// HasSubTree
// @Description: JZ26 树的子结构
// @param pRoot1
// @param pRoot2
// @return bool
//
func HasSubTree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
    if pRoot1 == nil || pRoot2 == nil {
        return false
    }
    return dfsHasSubTree(pRoot1, pRoot2) || HasSubTree(pRoot1.Left, pRoot2) || HasSubTree(pRoot1.Right, pRoot2)
}

//
// dfsHasSubTree
// @Description:JZ26 树的子结构
// @param pRoot1
// @param pRoot2
// @return bool
//
func dfsHasSubTree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
    if pRoot2 == nil {
        return true
    }
    if pRoot1 == nil {
        return false
    }
    if pRoot1.Val != pRoot2.Val {
        return false
    }
    return dfsHasSubTree(pRoot1.Left, pRoot2.Left) && dfsHasSubTree(pRoot1.Right, pRoot2.Right)
}
