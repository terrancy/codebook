package challenge

import trees "awesome/tree"

//
// IsUniqueTree
// @Description: 965. 单值二叉树
// @param root
// @return bool
//
func IsUniqueTree(root *trees.TreeNode) bool {
    data := -1
    return dfs(root, &data)
}

func dfs(root *trees.TreeNode, data *int) bool {
    if root == nil {
        return true
    }
    if *data == -1 {
        *data = root.Val
    }
    if root.Val != *data {
        return false
    }
    return dfs(root.Left, data) && dfs(root.Right, data)
}
