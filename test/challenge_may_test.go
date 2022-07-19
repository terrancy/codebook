package test

import (
    "awesome"
    "awesome/challenge"
    trees "awesome/tree"
    "fmt"
    "testing"
)

//
// TestGetAllElements
// @Description: 1305. 两棵二叉搜索树中的所有元素
// @param t
//
func TestGetAllElements(t *testing.T) {
    data1, data2 := []int{2, 1, 4}, []int{1, 0, 3}
    root1 := trees.BuildTreeNode(data1)
    root2 := trees.BuildTreeNode(data2)
    res := challenge.GetAllElements(root1, root2)
    fmt.Println(res)
}

//
// TestTagsValidator
// @Description: 591. 标签验证器
// @param t
//
func TestTagsValidator(t *testing.T) {
    code := "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
    res := challenge.TagsValidator(code)
    fmt.Println(res)
}

//
// TestReorderLogFiles
// @Description: 937. 重新排列日志文件
// @param t
//
func TestReorderLogFiles(t *testing.T) {
    logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
    res := challenge.ReorderLogFiles(logs)
    fmt.Println(res)
}

//
// TestFindTheWinner
// @Description: 1823. 找出游戏的获胜者
// @param t
//
func TestFindTheWinner(t *testing.T) {
    n, m := 5, 2
    res := challenge.FindTheWinner(n, m)
    fmt.Println(res)
}

//
// TestNumSubarrayProductLessThanK
// @Description: 713. 乘积小于 K 的子数组
// @param t
//
func TestNumSubarrayProductLessThanK(t *testing.T) {
    nums := []int{10, 5, 2, 6}
    k := 100
    res := challenge.NumSubarrayProductLessThanK(nums, k)
    fmt.Println(res)
}

//
// TestFindDuplicates
// @Description: 442. 数组中重复的数据
// @param t
//
func TestFindDuplicates(t *testing.T) {
    nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
    res := challenge.FindDuplicates(nums)
    fmt.Println(res)
}

//
// TestOneEditAway
// @Description: 面试题 01.05. 一次编辑
// @param t
//
func TestOneEditAway(t *testing.T) {
    first := "pale"
    second := "ple"
    result := challenge.OneEditAway(first, second)
    fmt.Println(result)
}

//
// TestInorderSuccessor
// @Description: 面试题 04.06. 后继者
// @param t
//
func TestInorderSuccessor(t *testing.T) {
    data := []int{2, 1, 3}
    root := trees.BuildTreeNode(data)
    p := &trees.TreeNode{Val: 1}
    node := challenge.InorderSuccessor(root, p)
    fmt.Println(node.Val)
}

//
// TestIsUniqueTree
// @Description: 965. 单值二叉树
// @param t
//
func TestIsUniqueTree(t *testing.T) {
    data := []int{1, 1, 1, 2, 1, awesome.INF, 1}
    root := trees.BuildTreeNode(data)
    result := challenge.IsUniqueTree(root)
    fmt.Println(result)
}
