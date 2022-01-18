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
	root := trees.MakeTreeNode(data)
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
	root1 := trees.MakeTreeNode(data1)
	root2 := trees.MakeTreeNode(data2)
	ans := trees.IsContains(root1, root2)
	fmt.Println(ans)
}

func TestBuildTree(t *testing.T) {
	// 中根 + 先根
	//pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	//mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
	//root := trees.DspPreMidBuildTreeII(pre, mid)
	//temp := trees.Serialize(root)
	//fmt.Println(temp)

	// 中根 + 后根
	//mid := []int{2, 1, 4, 3, 5}
	//post := []int{2, 4, 5, 3, 1}
	//root := trees.DspPostMidBuildTreeII(post, mid)
	//temp := trees.Serialize(root)
	//fmt.Println(temp)
}
