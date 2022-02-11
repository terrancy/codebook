package array

//
//  FindTarget
//  @Description: NC29 二维数组中的查找
//  @Solution: 1. 左下角或者右上角 2. if + continue
//  @param target
//  @param nums
//  @return bool
//
func FindTarget(target int, nums [][]int) bool {
    n := len(nums)
    if n == 0 {
        return false
    }
    m := len(nums[0])
    for i, j := n-1, 0; i >= 0 && j < m; {
        if nums[i][j] == target {
            return true
        }
        if nums[i][j] > target {
            i--
            continue
        }
        if nums[i][j] < target {
            j++
        }
    }
    return false
}
