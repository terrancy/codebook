package dp

import "awesome/array"

//
//  MaxLengthRepeatLess
//  @Description: NC41 最长无重复子数组
//  @Solution: 哈希表 + 双指针, l, r := -1, 0
//  @param nums
//  @return int
//
func MaxLengthRepeatLess(nums []int) int {
    return array.MaxLengthRepeatLess(nums)
}
