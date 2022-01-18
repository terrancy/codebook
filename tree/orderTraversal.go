package main

func postOrderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    data := make([]int, 0)
    dfsPost(root, &data)
    return data
}

//
//  dfsPost
//  @Description: 二叉树后根遍历
//  @param root
//  @param data
//
func dfsPost(root *TreeNode, data *[]int) {
    if root == nil {
        return
    }
    dfsPost(root.Left, data)
    dfsPost(root.Right, data)
    *data = append(*data, root.Val)
}

//
//  dfsPre
//  @Description: 二叉树先根遍历
//  @param root
//  @param data
//
func dfsPre(root *TreeNode, data *[]int) {
    if root == nil {
        return
    }
    
    *data = append(*data, root.Val)
    dfsPre(root.Left, data)
    dfsPre(root.Right, data)
}

//
//  dfsMid
//  @Description: 二叉树中根遍历
//  @param root
//  @param data
//
func dfsMid(root *TreeNode, data *[]int) {
    if root == nil {
        return
    }
    
    dfsMid(root.Left, data)
    *data = append(*data, root.Val)
    dfsMid(root.Right, data)
}

func buildTree() {

}
