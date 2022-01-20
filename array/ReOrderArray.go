package array

//
//  ReOrderArray
//  @Description: 奇偶位
//  @param nums
//  @return []int
//
func ReOrderArray(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }
    for i, cnt := 1, 0; i <= n; i += 2 {
        cnt++
        target := nums[i-1]
        for j := i - 1; j > cnt-1; j-- {
            nums[j] = nums[j-1]
        }
        nums[cnt-1] = target
    }
    return nums
}

//
//  ReOrderArrayII
//  @Description: NC77 调整数组顺序使奇数位于偶数前面(一)
//  @param nums
//  @return []int
//
func ReOrderArrayII(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }
    for i, pos := 0, 0; i < n; i++ {
        if nums[i]&1 == 0 {
            continue
        }
        target := nums[i]
        for j := i; j > pos; j-- {
            nums[j] = nums[j-1]
        }
        nums[pos] = target
        pos++
    }
    return nums
}
