package trees

// 这里主要介绍二叉树中关于和为某一个值的路径的问题。分为三类

//
// hasPathSum
// @Description: 二叉树中和为某一值的路径(一)
// @Description: 给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径
// @Description: 该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// @Link https://www.nowcoder.com/practice/508378c0823c423baa723ce448cbfd0c?tpId=117&&tqId=37719&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking
// @param root
// @param sum
// @return bool
//
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

//
// findPathII
// @Description: 输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径
// @Description: 该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
// @Description: 叶子节点是指没有子节点的节点,路径只能从父节点到子节点，不能从子节点到父节点,总节点数目为n
// @Link https://www.nowcoder.com/practice/b736e784e3e34731af99065031301bca?tpId=117&&tqId=37718&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking
// @param root
// @param sum
// @return [][]int
//
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

//
// treeSumIII
// @Description: 二叉树中和为某一值的路径(三)
// @Description: 给定一个二叉树root和一个整数值 sum ，求该树有多少路径的的节点值之和等于 sum
// @Description: 该题路径定义不需要从根节点开始，也不需要在叶子节点结束，但是一定是从父亲节点往下到孩子节点
// @Description: 总节点数目为n，保证最后返回的路径个数在整形范围内
// @Link: https://www.nowcoder.com/practice/965fef32cae14a17a8e86c76ffe3131f?tpId=117&&tqId=39297&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking
// @param root
// @param sum
// @return int
//
func findPathIII(root *TreeNode, sum int) int {
    cnt := 0
    dspFindPathTotalIII(root, sum, &cnt)
    return cnt
}

//
// dspFindPathIII
// @Description: 以树的根节点起始向下遍历
// @param root
// @param sum
// @param cnt
//
func dspFindPathIII(root *TreeNode, sum int, cnt *int) {
    if root == nil {
        return
    }
    if root.Val == sum {
        *cnt++
    }
    sum -= root.Val
    dspFindPathIII(root.Left, sum, cnt)
    dspFindPathIII(root.Right, sum, cnt)
}

//
// dspFindPathTotalIII
// @Description: 先根遍历各个节点
// @param root
// @param sum
// @param cnt
//
func dspFindPathTotalIII(root *TreeNode, sum int, cnt *int) {
    if root == nil {
        return
    }
    dspFindPathIII(root, sum, cnt)
    dspFindPathTotalIII(root.Left, sum, cnt)
    dspFindPathTotalIII(root.Right, sum, cnt)
}
