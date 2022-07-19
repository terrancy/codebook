package challenge

//
// FindDuplicates
// @Description: 442. 数组中重复的数据
// @param nums [4,3,2,7,8,2,3,1]
// @return []int
//
func FindDuplicates(nums []int) (res []int) {
    for idx := range nums {
        for nums[idx] != nums[nums[idx]-1] {
            nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
        }
    }
    
    for idx, val := range nums {
        if idx+1 != val {
            res = append(res, val)
        }
    }
    return
}
