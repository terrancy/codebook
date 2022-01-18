package main

//
//  findErrorSBT
//  @Description: NC58 找到搜索二叉树中两个错误的节点
//  @param root
//  @return []int
//
func findErrorSBT(root *TreeNode) []int {
    var data []int
    dsp(root, &data)
    return findErrorNode(data)
}

func dsp(root *TreeNode, data *[]int) {
    if root == nil {
        return
    }
    dsp(root.Left, data)
    *data = append(*data, root.Val)
    dsp(root.Right, data)
}

func findErrorNode(data []int) []int {
    n := len(data)
    if n == 0 {
        return data
    }
    
    var i, j int
    for i = 0; i < n-1; i++ {
        if data[i] > data[i+1] {
            break
        }
    }
    
    for j = n - 1; j >= 1; j-- {
        if data[j] < data[j-1] {
            break
        }
    }
    
    if data[i] > data[j] {
        return []int{data[j], data[i]}
    } else {
        return []int{data[i], data[j]}
    }
}
