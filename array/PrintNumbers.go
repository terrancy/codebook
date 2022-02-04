package array

import "math"

//
//  PrintNumbers
//  @Description: JZ17 打印从1到最大的n位数
//  @param n
//  @return []int
//
func PrintNumbers(n int) []int {
    // 普通解法
    max := int(math.Pow10(n))
    data := make([]int, 0)
    for i := 1; i < max; i++ {
        data = append(data, i)
    }
    return data
}

//
//  PrintNumbersBigData
//  @Description: JZ17 打印从1到最大的n位数
//  @Description: 考虑大数, 模拟大数加法
//  @param n
//  @return []int
//
func PrintNumbersBigData(n int) []int {
    return nil
}
