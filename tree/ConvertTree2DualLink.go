package trees

//
// ConvertTree2DualLink
// @Description: JZ36 二叉搜索树与双向链表
// @param pRootOfTree
// @return *TreeNode
//
var preRoot *TreeNode

func ConvertTree2DualLink(pRootOfTree *TreeNode) *TreeNode {
    if pRootOfTree == nil {
        return nil
    }
    root := pRootOfTree
    for root.Left != nil {
        root = root.Left
    }
    dfsConvertTree2DualLink(pRootOfTree)
    return root
}

//
// dfsConvertTree2DualLink
// @Description: 中根遍历
// @param root
// @return *TreeNode
//
func dfsConvertTree2DualLink(root *TreeNode) {
    if root == nil {
        return
    }
    dfsConvertTree2DualLink(root.Left)
    root.Left = preRoot
    if preRoot != nil {
        preRoot.Right = root
    }
    preRoot = root
    dfsConvertTree2DualLink(root.Right)
}

var head, tail *TreeNode

//
// ConvertTree2DualLinkII
// @Description: JZ36 二叉搜索树与双向链表
// @param pRootOfTree
// @return *TreeNode
//
func ConvertTree2DualLinkII(pRootOfTree *TreeNode) *TreeNode {
    if pRootOfTree == nil {
        return nil
    }
    inOrderConvertTree2DualLink(pRootOfTree)
    return head
}

//
// inOrderConvertTree2DualLink
// @Description: 中根遍历
// @param root
//
func inOrderConvertTree2DualLink(root *TreeNode) {
    if root == nil {
        return
    }
    inOrderConvertTree2DualLink(root.Left)
    if tail == nil {
        head = root
    } else {
        tail.Right = root
        root.Left = tail
    }
    tail = root
    inOrderConvertTree2DualLink(root.Right)
}
