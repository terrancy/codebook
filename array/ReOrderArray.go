package array

//
//  ReOrderArray
//  @Description: 奇偶位置 - 不是奇偶数
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
//  @Description: NC77 调整数组顺序使奇数位于偶数前面(一), 并保证奇数和奇数，偶数和偶数之间的相对位置不变(稳定型)
//  @Description: 时间复杂度O(N^2) 空间复杂度O(1)
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

//
//  ReOrderArrayRepeated
//  @Description: 调整数组顺序使奇数位于偶数前面(二)
//  @Description: 使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变
//  @Description: 时间复杂度 O(N),空间复杂度O(1)
//  @param nums
//  @return []int
//
func ReOrderArrayRepeated(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }
    left, right := 0, n-1
    for left < right {
        for left < right && nums[left]&1 == 1 {
            left++
        }
        for left < right && nums[right]&1 == 0 {
            right--
        }
        if left < right {
            nums[left], nums[right] = nums[right], nums[left]
        }
    }
    return nums
}

func ReOrderArrayRepeatedII(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }
    for slow, fast := 0, 0; fast < n; fast++ {
        if nums[fast]&1 == 1 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
    }
    return nums
}
