package slide

//
//  MaxWater
//  @Description: NC128 接雨水问题
//  @Solution: 水桶的短板效应,比较两边较小值.左边向右,右边向左.遇到大的替换,遇到小的盛水
//  @param []int
//  @return int
//
func MaxWater(nums []int) int64 {
    n := len(nums)
    if n < 3 {
        return int64(0)
    }
    res := int64(0)
    leftMax, rightMax := 0, 0
    for left, right := 0, n-1; left < right; {
        // 寻找短板
        if nums[left] <= nums[right] {
            // 看看能不能接水
            if nums[left] > leftMax {
                leftMax = nums[left]
            } else {
                res += int64(leftMax - nums[left])
            }
            left++
        } else {
            if nums[right] > rightMax {
                rightMax = nums[right]
            } else {
                res += int64(rightMax - nums[right])
            }
            right--
        }
    }
    return res
}
