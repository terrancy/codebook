package array

import "awesome"

//
// maxLength
// @Description: NC41 最长无重复子数组
// @param arr
// @return int
//
func maxLength(arr []int) int {
    // write code here
    n := len(arr)
    if n < 2 {
        return n
    }
    res := 0
    data := make(map[int]int)
    for left, right, val := -1, 0, 0; right < n; right++ {
        val = arr[right]
        if _, ok := data[val]; ok {
            left = awesome.MaxInt(left, data[val])
        }
        data[val] = right
        res = awesome.MaxInt(res, right-left)
    }
    return res
}

//
// MaxLengthRepeatLess
// @Description: NC41 最长无重复子数组
// @Solution: 哈希表 + 双指针, l, r := -1, 0
// @param arr
// @return int
//
func MaxLengthRepeatLess(arr []int) int {
    n := len(arr)
    hashMap := make(map[int]int)
    res := 0
    for left, right := -1, 0; right < n; right++ {
        val := arr[right]
        if _, ok := hashMap[val]; ok {
            left = awesome.MaxInt(left, hashMap[val])
        }
        hashMap[val] = right
        res = awesome.MaxInt(res, right-left)
    }
    return res
}
