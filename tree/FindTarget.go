package trees

//
//  FindTarget
//  @Description: 剑指 Offer II 056. 二叉搜索树中两个节点之和
//  @param root
//  @param k
//  @return bool
//
func FindTarget(root *TreeNode, k int) bool {
    findTargetFlag := false
    findTargetHash := make(map[int]int)
    inorderFindTarget(root, k, &findTargetFlag, findTargetHash)
    return findTargetFlag
}

func inorderFindTarget(root *TreeNode, k int, findTargetFlag *bool, findTargetHash map[int]int) {
    if root == nil || *findTargetFlag == true {
        return
    }
    inorderFindTarget(root.Left, k, findTargetFlag, findTargetHash)
    tmp := k - root.Val
    if _, ok := findTargetHash[tmp]; ok {
        *findTargetFlag = true
        return
    }
    findTargetHash[root.Val] = 1
    inorderFindTarget(root.Right, k, findTargetFlag, findTargetHash)
}
