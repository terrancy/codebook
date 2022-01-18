package trees

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	data := make([]int, 0)
	DfsPost(root, &data)
	return data
}

//
//  dfsPost
//  @Description: 二叉树后根遍历
//  @param root
//  @param data
//
func DfsPost(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	DfsPost(root.Left, data)
	DfsPost(root.Right, data)
	*data = append(*data, root.Val)
}

//
//  dfsPre
//  @Description: 二叉树先根遍历
//  @param root
//  @param data
//
func DfsPre(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	*data = append(*data, root.Val)
	DfsPre(root.Left, data)
	DfsPre(root.Right, data)
}

//
//  dfsMid
//  @Description: 二叉树中根遍历
//  @param root
//  @param data
//
func DfsMid(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	DfsMid(root.Left, data)
	*data = append(*data, root.Val)
	DfsMid(root.Right, data)
}

//
//  bfs
//  @Description: 先根遍历
//  @param root
//  @param data
//
func Bfs(root *TreeNode, data []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	node := &TreeNode{}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node, queue = queueShift(queue)
			data = append(data, node.Val)
			if node.Left != nil {
				queuePush(queue, node.Left)
			}
			if node.Right != nil {
				queuePush(queue, node.Right)
			}
			size--
		}
	}
}

//
//  PMBuildTree
//  @Description: 根据先根遍历和中根遍历构建树
//  @param pre
//  @param mid
//  @return *TreeNode
//
func PreMidBuildTree(pre []int, mid []int) *TreeNode {
	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}
	midMap := makeMidMap(mid)
	return dspPreMidBuildTree(pre, mid, 0, len(pre)-1, 0, len(mid)-1, midMap)
}

//
//  dspPMBuildTree
//  @Description: 根据先根遍历和中根遍历构建树,递归构建
//  @param pre
//  @param mid
//  @param pl
//  @param pr
//  @param ml
//  @param mr
//  @param midMap
//  @return *TreeNode
//
func dspPreMidBuildTree(pre []int, mid []int, pl int, pr int, ml int, mr int, midMap map[int]int) *TreeNode {
	if pl > pr {
		return nil
	}
	rootVal := pre[pl]
	root := &TreeNode{rootVal, nil, nil}
	// 孩子节点
	midIndex := midMap[rootVal]
	leftLen := midIndex - ml
	root.Left = dspPreMidBuildTree(pre, mid, pl+1, pl+leftLen, ml, midIndex-1, midMap)
	root.Right = dspPreMidBuildTree(pre, mid, pl+leftLen+1, pr, midIndex+1, mr, midMap)
	return root
}

//
//  PostMidBuildTree
//  @Description: 根据后根遍历和中根遍历构建树
//  @param post
//  @param mid
//  @return *TreeNode
//
func PostMidBuildTree(post []int, mid []int) *TreeNode {
	if len(post) == 0 || len(mid) == 0 {
		return nil
	}
	midMap := makeMidMap(mid)
	return dspPostMidBuildTree(post, mid, 0, len(post)-1, 0, len(mid)-1, midMap)
}

//
//  dspPostMidBuildTree
//  @Description: 根据后根遍历和中根遍历构建树,递归构建
//  @param post
//  @param mid
//  @param pl
//  @param pr
//  @param ml
//  @param mr
//  @param midMap
//  @return *TreeNode
//
func dspPostMidBuildTree(post []int, mid []int, pl int, pr int, ml int, mr int, midMap map[int]int) *TreeNode {
	if pl > pr {
		return nil
	}
	rootVal := post[pr]
	root := &TreeNode{rootVal, nil, nil}
	// 孩子节点
	midIndex := midMap[rootVal]
	leftLen := midIndex - ml
	root.Left = dspPostMidBuildTree(post, mid, pl, pl+leftLen-1, ml, midIndex-1, midMap)
	root.Right = dspPostMidBuildTree(post, mid, pl+leftLen, pr-1, midIndex+1, mr, midMap)
	return root
}

func makeMidMap(mid []int) map[int]int {
	midMap := make(map[int]int, len(mid))
	for idx, val := range mid {
		midMap[val] = idx
	}
	return midMap
}

////////////////////////////////////////////////////////////////////////////////
//
//  DspPreMidBuildTreeII
//  @Description:
//  @param pre
//  @param mid
//  @param midMap
//  @return *TreeNode
//
func DspPreMidBuildTreeII(preOrder []int, inOrder []int) *TreeNode {
	// 递归结束
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	// 根节点
	rootVal := preOrder[0]
	// 孩子节点
	root := &TreeNode{rootVal, nil, nil}
	idx := getMidOrderIndex(rootVal, inOrder)
	root.Left = DspPreMidBuildTreeII(preOrder[1:idx+1], inOrder[:idx])
	root.Right = DspPreMidBuildTreeII(preOrder[idx+1:], inOrder[idx+1:])
	return root
}

//
//  DspPostMidBuildTreeII
//  @Description:
//  @param inOrder
//  @param postOrder
//  @return *TreeNode
//
func DspPostMidBuildTreeII(inOrder []int, postOrder []int) *TreeNode {
	if len(postOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	rootVal := inOrder[len(inOrder)-1]
	root := &TreeNode{rootVal, nil, nil}
	idx := getMidOrderIndex(rootVal, inOrder)
	root.Left = DspPostMidBuildTreeII(inOrder[:idx], postOrder[:idx])
	root.Right = DspPostMidBuildTreeII(inOrder[idx+1:], postOrder[idx:len(postOrder)-1])
	return root
}

func getMidOrderIndex(target int, data []int) int {
	for idx, val := range data {
		if val == target {
			return idx
		}
	}
	return -1
}
