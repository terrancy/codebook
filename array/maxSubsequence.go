package array

import (
    "awesome"
    "sort"
)

//
//  maxSubsequence
//  @Description: 连续子数组的最大长度
//  @param arr
//  @return int
//
func maxSubsequence(arr []int) int {
    n := len(arr)
    if n == 1 {
        return 1
    }
    sort.Ints(arr)
    cnt := 1
    maxLen := 1
    for i := 1; i < n; i++ {
        if arr[i] == arr[i-1] {
            continue
        } else if arr[i]-arr[i-1] == 1 {
            cnt++
        } else {
            cnt = 1
        }
        maxLen = awesome.MaxInt(maxLen, cnt)
    }
    return maxLen
}
