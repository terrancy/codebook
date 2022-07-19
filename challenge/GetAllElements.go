package challenge

import trees "awesome/tree"

//
// GetAllElements
// @Description: 1305. 两棵二叉搜索树中的所有元素
// @param root1
// @param root2
// @return []int
//
func GetAllElements(root1, root2 *trees.TreeNode) []int {
    var data1, data2, res []int
    getAllElementsDFS(root1, &data1)
    getAllElementsDFS(root2, &data2)
    if data1 == nil {
        return data2
    }
    if data2 == nil {
        return data1
    }
    i, j := 0, 0
    n, m := len(data1), len(data2)
    for i < n && j < m {
        if data1[i] < data2[j] {
            res = append(res, data1[i])
            i++
        } else {
            res = append(res, data2[j])
            j++
        }
    }
    if i < n {
        res = append(res, data1[i:]...)
    }
    if j < n {
        res = append(res, data2[j:]...)
    }
    
    return res
}

func getAllElementsDFS(root *trees.TreeNode, data *[]int) {
    if root == nil {
        return
    }
    getAllElementsDFS(root.Left, data)
    *data = append(*data, root.Val)
    getAllElementsDFS(root.Right, data)
}
