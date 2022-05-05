package challenge

//
//  NumSubarrayProductLessThanK
//  @Description: 713. 乘积小于 K 的子数组
//  @param nums
//  @param k
//  @return int
//
func NumSubarrayProductLessThanK(nums []int, k int) int {
    if k == 0 {
        return 0
    }
    res := 0
    n := len(nums)
    for i, j, prod := 0, 0, 1; j < n; j++ {
        prod *= nums[j]
        for prod >= k && i <= j {
            prod /= nums[i]
            i++
        }
        res += j - i + 1
    }
    return res
}
