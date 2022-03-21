package array

//
//  FirstMissingPositive
//  @Description: 41. 缺失的第一个正数
//  @param nums
//  @return int
//
func FirstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] <= 0 {
            nums[i] = n + 1
        }
    }
    for i := 0; i < n; i++ {
        num := nums[i]
        if num > 0 && num < n && nums[num-1] > 0 {
            nums[num-1] = - nums[num-1]
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            return i + 1
        }
    }
    return n + 1
}
