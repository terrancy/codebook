package array

import "awesome"

//
// MaxSubArray
// @Description: JZ42 连续子数组的最大和
// @param nums
// @return int
//
func MaxSubArray(nums []int) int {
    sum := 0
    max := nums[0]
    for _, val := range nums {
        sum = awesome.MaxInt(sum+val, val)
        max = awesome.MaxInt(max, sum)
    }
    return max
}

func MaxSubArrayII(nums []int) int {
    sum := 0
    max := nums[0]
    for _, val := range nums {
        if sum+val < val {
            sum = val
        } else {
            sum = sum + val
        }
        if max < sum {
            max = sum
        }
    }
    return max
}
