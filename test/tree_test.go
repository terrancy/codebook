package test

import (
    "awesome"
    "awesome/tree"
    "fmt"
    "testing"
)

//
//  testSerialize
//  @Description: NC123 序列化二叉树
//  @param t
//
func TestSerialize(t *testing.T) {
    data := []int{8, 6, 10, 5, 7, 9, 11}
    root := trees.BuildTreeNode(data)
    temp := trees.Serialize(root)
    fmt.Println(temp)
    str := trees.Array2String(temp)
    fmt.Println(str)
    //trees.Deserialize(str)
}

//
//  TestIsContains
//  @Description: NC98 判断t1树中是否有与t2树完全相同的子树
//  @param t
//
func TestIsContains(t *testing.T) {
    data1 := []int{1, 2, 3, 4, 5, 6, 7, awesome.INF, 8, 9}
    data2 := []int{2, 4, 5, awesome.INF, 8, 9}
    root1 := trees.BuildTreeNode(data1)
    root2 := trees.BuildTreeNode(data2)
    ans := trees.IsContains(root1, root2)
    fmt.Println(ans)
}

//
//  TestBuildTree
//  @Description: 构建二叉树
//  @param t
//
func TestBuildTree(t *testing.T) {
    // 中根 + 先根
    t.Run("preTest", func(t *testing.T) {
        pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
        mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
        root := trees.DspPreMidBuildTreeII(pre, mid)
        temp := trees.Serialize(root)
        fmt.Println(temp)
    })
    
    // 中根 + 后根
    t.Run("PostTest", func(t *testing.T) {
        mid := []int{2, 1, 4, 3, 5}
        post := []int{2, 4, 5, 3, 1}
        root := trees.DspPostMidBuildTreeII(post, mid)
        temp := trees.Serialize(root)
        fmt.Println(temp)
    })
}

//
//  TestVerifySequenceOfBST
//  @Description: JZ33 二叉搜索树的后序遍历序列
//  @param t
//
func TestVerifySequenceOfBST(t *testing.T) {
    data := []int{7, 4, 6, 5, 9, 11, 10, 8}
    res := trees.VerifySequenceOfBST(data)
    fmt.Println(res)
}

//
//  TestHadSubTree
//  @Description:JZ26 树的子结构
//  @param t
//
func TestHasSubTree(t *testing.T) {
    data := []int{8, 8, 7, 9, 2, awesome.INF, awesome.INF, awesome.INF, awesome.INF, 4, 7}
    tmp := []int{8, 9, 2}
    head := trees.BuildTreeNode(data)
    subHead := trees.BuildTreeNode(tmp)
    res := trees.HasSubTree(head, subHead)
    if res == true {
        fmt.Println("yes")
    } else {
        fmt.Println("no")
    }
}

//
//  TestConvertTree2DualLink
//  @Description: JZ36 二叉搜索树与双向链表
//  @param t
//
func TestConvertTree2DualLink(t *testing.T) {
    data := []int{10, 6, 14, 4, 8, 12, 16}
    head := trees.BuildTreeNode(data)
    trees.ConvertTree2DualLinkII(head)
}

//
//  TestIsSymmetric
//  @Description: JZ28 对称的二叉树
//  @param t
//
func TestIsSymmetric(t *testing.T) {
    data := []int{2, 3, 3, 4, 5, 5, 4, awesome.INF, awesome.INF, 8, 9, awesome.INF, awesome.INF, 9, 8}
    root := trees.BuildTreeNode(data)
    res := trees.IsSymmetric(root)
    fmt.Println(res)
}

func TestKthNode(t *testing.T) {
    data := []int{5, 3, 6, 2, 4, awesome.INF, awesome.INF, 1}
    root := trees.BuildTreeNode(data)
    res := trees.KthNode(root, 3)
    fmt.Println(res)
}

//
//  TestMaxPathSum
//  @Description: NC6 二叉树中的最大路径和
//  @param t
//
func TestMaxPathSum(t *testing.T) {
    data := []int{-10, 9, 20, awesome.INF, awesome.INF, 15, 7}
    root := trees.BuildTreeNode(data)
    res := trees.MaxPathSum(root)
    fmt.Println(res)
}

//
//  TestPruneTree
//  @Description: 814. 二叉树剪枝 || 剑指 Offer II 047. 二叉树剪枝
//  @param t
//
func TestPruneTree(t *testing.T) {
    data := []int{1, 0, 1, 0, 0, 0, 1}
    root := trees.BuildTreeNode(data)
    node := trees.PruneTree(root)
    res := trees.Serialize(node)
    fmt.Println(res)
}

//
//  TestFindTarget
//  @Description: 剑指 Offer II 056. 二叉搜索树中两个节点之和
//  @param t
//
func TestFindTarget(t *testing.T) {
    data := []int{8, 6, 10, 5, 7, 9, 11}
    k := 12
    root := trees.BuildTreeNode(data)
    res := trees.FindTarget(root, k)
    fmt.Println(res)
}

//
//  TestBstFromPreorder
//  @Description: 1008. 前序遍历构造二叉搜索树
//  @param t
//
func TestBstFromPreorder(t *testing.T) {
    data := []int{8, 10}
    root := trees.BstFromPreorder(data)
    res := trees.Serialize(root)
    fmt.Println(res)
}
