package trees

//
// bfs
// @Description: NC15 求二叉树的层序遍历
// @param root
// @return [][]int
//
func bfs(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    res := make([][]int, 0)
    for len(queue) > 0 {
        temp := make([]int, 0)
        for size := len(queue); size > 0; size-- {
            node := queue[0]
            queue = queue[1:]
            temp = append(temp, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, temp)
    }
    return res
}
