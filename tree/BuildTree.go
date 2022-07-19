package trees

//
// buildTree
// @Description: NC12 重建二叉树 - 先根与中根
// @Solution: 注意事项: 1. 索引idx的获取与切片区间 2.递归终止条件,len(inorder)==0 非 inorder == nil
// @param preorder
// @param inorder
// @return *TreeNode
//
func BuildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    val := preorder[0]
    root := &TreeNode{Val: val}
    idx := getIdx(inorder, val)
    root.Left = BuildTree(preorder[1:idx+1], inorder[:idx])
    root.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])
    return root
}

//
// buildTree
// @Description: NC223 从中序与后序遍历序列构造二叉树
// @param inorder
// @param postorder
// @return *TreeNode
//
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    val := postorder[len(postorder)-1]
    root := &TreeNode{Val: val}
    idx := getIdx(inorder, val)
    root.Left = buildTree(inorder[:idx], postorder[:idx])
    root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
    return root
}

func getIdx(data []int, target int) int {
    for idx, val := range data {
        if val == target {
            return idx
        }
    }
    return -1
}

//
// sortedArrayToBST
// @Description: NC11 将升序数组转化为平衡二叉搜索树
// @Solution: 1.判断条件;2.根的获取 mid := (n + 1) >> 1;3.左节点 num[:mid] 右结点 num[mid+1:]
// @param num
// @return *TreeNode
//
func SortedArrayToBST(num []int) *TreeNode {
    n := len(num)
    if n == 0 {
        return nil
    }
    if n == 1 {
        return &TreeNode{Val: num[0]}
    }
    mid := (n + 1) >> 1
    root := &TreeNode{Val: num[mid]}
    root.Left = SortedArrayToBST(num[:mid])
    root.Right = SortedArrayToBST(num[mid+1:])
    return root
}

//
// bstFromPreorder
// @Description: LC1008. 前序遍历构造二叉搜索树
// @param preorder
// @return *TreeNode
//
func BstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    val := preorder[0]
    root := &TreeNode{Val: val}
    idx := getPreorderIdx(preorder)
    root.Left = BstFromPreorder(preorder[1:idx])
    root.Right = BstFromPreorder(preorder[idx:])
    return root
}

//
// getPreorderIdx
// @Description: 获取idx
// @param data
// @return int
//
func getPreorderIdx(data []int) int {
    // 有左右子树
    for idx, val := range data {
        if val > data[0] {
            return idx
        }
    }
    // 右子树位空 || 左子树为空
    return len(data)
}
