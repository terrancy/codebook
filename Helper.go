package awesome

import (
    "sort"
)

// 一些最值
// math.MaxInt32, math.MinInt32

// 比较最值
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func MinInt(a, b int) int {
    if a > b {
        return b
    }
    return a
}

// 翻转问题

//
// arrayReverse
// @Description: 数组翻转
// @param data
// @return []int
//
func DataReverse(data []int) []int {
    n := len(data)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
    return data
}

//
// DataDualReverse
// @Description: 二维数组翻转
// @param data
// @return [][]int
//
func DataDualReverse(data [][]int) [][]int {
    for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
    return data
}

//
// StrReverse
// @Description: 字符翻转
// @param str
// @return string
//
func StrReverse(str string) string {
    n := len(str)
    if n < 2 {
        return str
    }
    nums := []byte(str)
    for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
        nums[l], nums[r] = nums[r], nums[l]
    }
    return string(nums)
}

// 排序相关

//
// DataSort
// @Description: 数组排序
// @param data
// @return []int
//
func DataSort(data []int) []int {
    n := len(data)
    if n < 2 {
        return data
    }
    sort.Slice(data, func(i, j int) bool {
        return data[i] > data[j]
    })
    return data
}

//
// StrSort
// @Description: 字符排序
// @param str
// @return string
//
func StrSort(str string) []byte {
    n := len(str)
    bytes := []byte(str)
    if n < 2 {
        return bytes
    }
    sort.Slice(bytes, func(i, j int) bool {
        return bytes[i] < bytes[j]
    })
    return bytes
}

//
// Str2Sort
// @Description: 字符串排序
// @param strList
// @return []byte
//
func Str2Sort(strList []string) []string {
    n := len(strList)
    if n < 2 {
        return strList
    }
    sort.Slice(strList, func(i, j int) bool {
        return strList[i]+strList[j] < strList[j]+strList[i]
    })
    return strList
}
