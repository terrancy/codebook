package main

import "sort"

//
//  stringSort
//  @Description: 字符排序
//  @param str
//  @return string
//
func stringSort(str string) []byte {
    bytes := []byte(str)
    sort.Slice(bytes, func(i, j int) bool {
        return bytes[i] < bytes[j]
    })
    return bytes
}
