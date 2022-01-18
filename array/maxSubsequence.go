package main

import "sort"

//
//  maxSubsequence
//  @Description:
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
        maxLen = maxInt(maxLen, cnt)
    }
    return maxLen
}

//func main() {
//    data := []int{1, 1, 1}
//    ret := maxSubsequence(data)
//    print(ret)
//}
