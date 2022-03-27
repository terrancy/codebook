package challenge

import "strconv"

//
//  CalPoints
//  @Description: 682. 棒球比赛
//  @param ops
//  @return int
//
func CalPoints(ops []string) int {
    n := len(ops)
    nums := make([]int, n)
    idx := 0
    for i := 0; i < n; i++ {
        if ops[i] == "+" {
            nums[idx] = nums[idx-2] + nums[idx-1]
        } else if ops[i] == "D" {
            nums[idx] = nums[idx-1] * 2
        } else if ops[i] == "C" {
            idx = idx - 2
        } else {
            digital, err := strconv.Atoi(ops[i])
            if err == nil {
                nums[idx] = digital
            }
        }
        idx++
    }
    res := 0
    for i := 0; i < idx; i++ {
        res += nums[i]
    }
    return res
}
