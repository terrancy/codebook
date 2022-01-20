package array

//
//  MajorityElement
//  @Description: NC73 数组中出现次数超过一半的数字
//  @param nums
//  @return int
//
func MajorityElement(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    cnt := 0
    vote := -1
    for i := 0; i < n; i++ {
        if cnt == 0 || vote == nums[i] {
            vote = nums[i]
            cnt++
        } else {
            cnt--
        }
        if cnt > n/2 {
            break
        }
    }
    return vote
}
