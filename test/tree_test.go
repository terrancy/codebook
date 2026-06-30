package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome"
	"terrancy/awesome/tree"
)

var serializeCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{
		name:     "test_case_1",
		data:     []int{8, 6, 10, 5, 7, 9, 11},
		expected: []int{8, 6, 10, 5, 7, 9, 11, awesome.INF, awesome.INF, awesome.INF, awesome.INF, awesome.INF, awesome.INF, awesome.INF, awesome.INF},
	},
}

// TestSerialize NC123 序列化二叉树
func TestSerialize(t *testing.T) {
	for _, tt := range serializeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.Serialize(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var isContainsCases = []struct {
	name     string
	data1    []int
	data2    []int
	expected bool
}{
	{
		name:     "test_case_1",
		data1:    []int{1, 2, 3, 4, 5, 6, 7, awesome.INF, 8, 9},
		data2:    []int{2, 4, 5, awesome.INF, 8, 9},
		expected: true,
	},
}

// TestIsContains NC98 判断t1树中是否有与t2树完全相同的子树
func TestIsContains(t *testing.T) {
	for _, tt := range isContainsCases {
		t.Run(tt.name, func(t *testing.T) {
			root1 := trees.BuildTreeNode(tt.data1)
			root2 := trees.BuildTreeNode(tt.data2)
			res := trees.IsContains(root1, root2)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var dspPreMidBuildTreeCases = []struct {
	name     string
	pre      []int
	mid      []int
	expected []int
}{
	{
		name:     "test_case_1",
		pre:      []int{1, 2, 4, 7, 3, 5, 6, 8},
		mid:      []int{4, 7, 2, 1, 5, 3, 8, 6},
		expected: []int{1, 2, 3, 4, awesome.INF, 5, 6, awesome.INF, 7, awesome.INF, awesome.INF, 8, awesome.INF, awesome.INF, awesome.INF, awesome.INF, awesome.INF},
	},
}

// TestDspPreMidBuildTreeII 构建二叉树(先根+中根)
func TestDspPreMidBuildTreeII(t *testing.T) {
	for _, tt := range dspPreMidBuildTreeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.DspPreMidBuildTreeII(tt.pre, tt.mid)
			res := trees.Serialize(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var dspPostMidBuildTreeCases = []struct {
	name     string
	post     []int
	mid      []int
	expected []int
}{
	{
		name:     "test_case_1",
		post:     []int{2, 4, 5, 3, 1},
		mid:      []int{2, 1, 4, 3, 5},
		expected: []int{1, 3, awesome.INF, 5, awesome.INF, 4, awesome.INF, 2, awesome.INF, awesome.INF, awesome.INF},
	},
}

// TestDspPostMidBuildTreeII 构建二叉树(后根+中根)
func TestDspPostMidBuildTreeII(t *testing.T) {
	for _, tt := range dspPostMidBuildTreeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.DspPostMidBuildTreeII(tt.post, tt.mid)
			res := trees.Serialize(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var verifySequenceOfBSTCases = []struct {
	name     string
	data     []int
	expected bool
}{
	{
		name:     "test_case_1",
		data:     []int{7, 4, 6, 5, 9, 11, 10, 8},
		expected: false,
	},
}

// TestVerifySequenceOfBST JZ33 二叉搜索树的后序遍历序列
func TestVerifySequenceOfBST(t *testing.T) {
	for _, tt := range verifySequenceOfBSTCases {
		t.Run(tt.name, func(t *testing.T) {
			res := trees.VerifySequenceOfBST(tt.data)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var hasSubTreeCases = []struct {
	name     string
	data1    []int
	data2    []int
	expected bool
}{
	{
		name:     "test_case_1",
		data1:    []int{8, 8, 7, 9, 2, awesome.INF, awesome.INF, awesome.INF, awesome.INF, 4, 7},
		data2:    []int{8, 9, 2},
		expected: true,
	},
}

// TestHasSubTree JZ26 树的子结构
func TestHasSubTree(t *testing.T) {
	for _, tt := range hasSubTreeCases {
		t.Run(tt.name, func(t *testing.T) {
			root1 := trees.BuildTreeNode(tt.data1)
			root2 := trees.BuildTreeNode(tt.data2)
			res := trees.HasSubTree(root1, root2)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var convertTree2DualLinkCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{
		name:     "test_case_1",
		data:     []int{10, 6, 14, 4, 8, 12, 16},
		expected: []int{4, 6, 8, 10, 12, 14, 16},
	},
}

// TestConvertTree2DualLinkII JZ36 二叉搜索树与双向链表
func TestConvertTree2DualLinkII(t *testing.T) {
	for _, tt := range convertTree2DualLinkCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			head := trees.ConvertTree2DualLinkII(root)
			vals := make([]int, 0)
			for cur := head; cur != nil; cur = cur.Right {
				vals = append(vals, cur.Val)
				if cur.Right != nil {
					assert.Equal(t, cur, cur.Right.Left)
				}
			}
			assert.Equal(t, tt.expected, vals)
		})
	}
}

var isBalancedCases = []struct {
	name     string
	data     []int
	expected bool
}{
	{
		name:     "test_case_1",
		data:     []int{1, 2, 2, 3, 3, awesome.INF, awesome.INF, 4, 4},
		expected: false,
	},
}

// TestIsBalanced NC62 判断是不是平衡二叉树
func TestIsBalanced(t *testing.T) {
	for _, tt := range isBalancedCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.IsBalanced(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var isSymmetricCases = []struct {
	name     string
	data     []int
	expected bool
}{
	{
		name:     "test_case_1",
		data:     []int{2, 3, 3, 4, 5, 5, 4, awesome.INF, awesome.INF, 8, 9, awesome.INF, awesome.INF, 9, 8},
		expected: false,
	},
}

// TestIsSymmetric JZ28 对称的二叉树
func TestIsSymmetric(t *testing.T) {
	for _, tt := range isSymmetricCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.IsSymmetric(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var kthNodeCases = []struct {
	name     string
	data     []int
	k        int
	expected int
}{
	{
		name:     "test_case_1",
		data:     []int{5, 3, 6, 2, 4, awesome.INF, awesome.INF, 1},
		k:        3,
		expected: 4,
	},
}

// TestKthNode 二叉搜索树第K大的值
func TestKthNode(t *testing.T) {
	for _, tt := range kthNodeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.KthNode(root, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var maxPathSumCases = []struct {
	name     string
	data     []int
	expected int
}{
	{
		name:     "test_case_1",
		data:     []int{-10, 9, 20, awesome.INF, awesome.INF, 15, 7},
		expected: 42,
	},
}

// TestMaxPathSum NC6 二叉树中的最大路径和
func TestMaxPathSum(t *testing.T) {
	for _, tt := range maxPathSumCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.MaxPathSum(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var pruneTreeCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{
		name:     "test_case_1",
		data:     []int{1, 0, 1, 0, 0, 0, 1},
		expected: []int{1, awesome.INF, 1, awesome.INF, 1, awesome.INF, awesome.INF},
	},
}

// TestPruneTree LC814.二叉树剪枝
func TestPruneTree(t *testing.T) {
	for _, tt := range pruneTreeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			node := trees.PruneTree(root)
			res := trees.Serialize(node)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var findTargetCases = []struct {
	name     string
	data     []int
	k        int
	expected bool
}{
	{
		name:     "test_case_1",
		data:     []int{8, 6, 10, 5, 7, 9, 11},
		k:        12,
		expected: true,
	},
}

// TestFindTarget 剑指OfferII056.二叉搜索树中两个节点之和
func TestFindTarget(t *testing.T) {
	for _, tt := range findTargetCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.FindTarget(root, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var bstFromPreorderCases = []struct {
	name     string
	data     []int
	expected []int
}{
	{
		name:     "test_case_1",
		data:     []int{8, 10},
		expected: []int{8, awesome.INF, 10, awesome.INF, awesome.INF},
	},
}

// TestBstFromPreorder LC1008.前序遍历构造二叉搜索树
func TestBstFromPreorder(t *testing.T) {
	for _, tt := range bstFromPreorderCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BstFromPreorder(tt.data)
			res := trees.Serialize(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var findPathIIICases = []struct {
	name     string
	data     []int
	k        int
	expected int
}{
	{
		name:     "test_case_1",
		data:     []int{10, 5, -3, 3, 2, awesome.INF, 11, 3, -2, awesome.INF, 1},
		k:        8,
		expected: 3,
	},
	{
		name:     "test_case_2",
		data:     []int{5, 4, 8, 11, awesome.INF, 13, 4, 7, 2, awesome.INF, awesome.INF, 5, 1},
		k:        22,
		expected: 3,
	},
}

// TestFindPathIII LC437.路径总和III
func TestFindPathIII(t *testing.T) {
	for _, tt := range findPathIIICases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.FindPathIII(root, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// TestFindPathIIIPrefixSum LC437.路径总和III(前缀和)
func TestFindPathIIIPrefixSum(t *testing.T) {
	for _, tt := range findPathIIICases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.FindPathIIIPrefixSum(root, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var widthOfBinaryTreeCases = []struct {
	name     string
	data     []int
	expected int
}{
	{
		name:     "test_case_1",
		data:     []int{1, 3, 2, 5, 3, awesome.INF, 9},
		expected: 4,
	},
	{
		name:     "test_case_2",
		data:     []int{1, 3, 2, 5, awesome.INF},
		expected: 2,
	},
	{
		name:     "test_case_3",
		data:     []int{1, 3, 2, 5},
		expected: 2,
	},
}

// TestWidthOfBinaryTree LC662.二叉树最大宽度
func TestWidthOfBinaryTree(t *testing.T) {
	for _, tt := range widthOfBinaryTreeCases {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.BuildTreeNode(tt.data)
			res := trees.WidthOfBinaryTree(root)
			assert.Equal(t, tt.expected, res)
		})
	}
}