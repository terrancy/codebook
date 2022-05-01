package test

import (
    "awesome/challenge"
    trees "awesome/tree"
    "fmt"
    "testing"
)

func TestGetAllElements(t *testing.T) {
    data1, data2 := []int{2, 1, 4}, []int{1, 0, 3}
    root1 := trees.BuildTreeNode(data1)
    root2 := trees.BuildTreeNode(data2)
    res := challenge.GetAllElements(root1, root2)
    fmt.Println(res)
}
