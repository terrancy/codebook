package test

import (
    "awesome/challenge"
    trees "awesome/tree"
    "fmt"
    "testing"
)

//
//  TestGetAllElements
//  @Description: 1305. 两棵二叉搜索树中的所有元素
//  @param t
//
func TestGetAllElements(t *testing.T) {
    data1, data2 := []int{2, 1, 4}, []int{1, 0, 3}
    root1 := trees.BuildTreeNode(data1)
    root2 := trees.BuildTreeNode(data2)
    res := challenge.GetAllElements(root1, root2)
    fmt.Println(res)
}

//
//  TestTagsValidator
//  @Description: 591. 标签验证器
//  @param t
//
func TestTagsValidator(t *testing.T) {
    code := "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
    res := challenge.TagsValidator(code)
    fmt.Println(res)
}

//
//  TestReorderLogFiles
//  @Description: 937. 重新排列日志文件
//  @param t
//
func TestReorderLogFiles(t *testing.T) {
    logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
    res := challenge.ReorderLogFiles(logs)
    fmt.Println(res)
}

//
//  TestFindTheWinner
//  @Description: 1823. 找出游戏的获胜者
//  @param t
//
func TestFindTheWinner(t *testing.T) {
    n, m := 5, 2
    res := challenge.FindTheWinner(n, m)
    fmt.Println(res)
}
