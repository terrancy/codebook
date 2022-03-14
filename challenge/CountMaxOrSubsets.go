package challenge

//
//  CountMaxOrSubsets
//  @Description: 2044. 统计按位或能得到最大值的子集数目
//  @param nums
//  @return int
//
func CountMaxOrSubsets(nums []int) int {
    var n int = len(nums)
    
    var maxOr int = 0
    var cnt int = 0
    
    for state := 0; state < (1 << n); state++ {
        var curOr int = 0
        for i := 0; i < n; i++ {
            if ((state >> i) & 1) == 1 {
                curOr |= nums[i]
            }
        }
        if curOr == maxOr {
            cnt++
        } else if curOr > maxOr {
            maxOr = curOr
            cnt = 1
        }
    }
    
    return cnt
}
