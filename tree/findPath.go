package trees

// 这里主要介绍二叉树中关于和为某一个值的路径的问题。分为三类

// hasPathSum
// @Title: LC112.路径总和
// @Description: 给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径
// @Description: 该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// @Link: https://leetcode.cn/problems/path-sum/
// @param root
// @param sum
// @return bool
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSum(root, sum)
}

func dspHasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	sum -= root.Val
	return dspHasPathSum(root, sum) || dspHasPathSum(root, sum)
}

// findPathII
// @Title: LC113.路径总和II
// @Description: 输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径
// @Description: 该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// @Description: 叶子节点是指没有子节点的节点,路径只能从父节点到子节点，不能从子节点到父节点,总节点数目为n
// @Link: https://leetcode.cn/problems/path-sum-ii/
// @param root
// @param sum
// @return [][]int
func findPathII(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	dspFindPathII(root, sum, []int{}, &res)
	return res
}

func dspFindPathII(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	// 先序遍历
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	dspFindPathII(root.Left, sum-root.Val, path, res)
	dspFindPathII(root.Right, sum-root.Val, path, res)
}

// FindPathIII
// @Title: LC437.路径总和III
// @Description: 给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum
// @Description: 该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
// @Description: 总节点数目为n，保证最后返回的路径个数在整形范围内
// @Link: https://leetcode.cn/problems/path-sum-iii/
// @param root
// @param k
// @return int
func FindPathIII(root *TreeNode, k int) int {
	var (
		cnt     = 0
		dfs     func(root *TreeNode, x int)
		dfsPath func(root *TreeNode, x int)
	)

	dfs = func(root *TreeNode, x int) {
		if root == nil {
			return
		}
		if x == root.Val {
			cnt++
		}
		dfs(root.Left, x-root.Val)
		dfs(root.Right, x-root.Val)
	}

	dfsPath = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		dfs(root, k)
		dfsPath(root.Left, k)
		dfsPath(root.Right, k)
	}

	dfsPath(root, k)
	return cnt
}

// FindPathIIIPrefixSum
// @Title: LC437.路径总和III(前缀和)
// @Description: 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// @Link: https://leetcode.cn/problems/path-sum-iii/
// @param root
// @param k
// @return int
func FindPathIIIPrefixSum(root *TreeNode, k int) int {
	var (
		cnt    = 0
		preSum = map[int]int{0: 1}
		dfs    func(root *TreeNode, total int)
	)

	dfs = func(root *TreeNode, total int) {
		if root == nil {
			return
		}

		total += root.Val
		if val, ok := preSum[total-k]; ok {
			cnt += val
		}
		preSum[total] += 1
		dfs(root.Left, total)
		dfs(root.Right, total)
		preSum[total]--
	}

	dfs(root, 0)
	return cnt
}
