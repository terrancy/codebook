package test

import (
    "awesome/challenge"
    "fmt"
    "testing"
)

//
//  TestPlatesBetweenCandles
//  @Description: 2055. 蜡烛之间的盘子
//  @param t
//
func TestPlatesBetweenCandles(t *testing.T) {
    s := "**|**|***|"
    queries := [][]int{{2, 5}, {5, 9}}
    res := challenge.PlatesBetweenCandles(s, queries)
    fmt.Println(res)
}

//
//  TestBestRotation
//  @Description: 798. 得分最高的最小轮调798. 得分最高的最小轮调
//  @param t
//
func TestBestRotation(t *testing.T) {
    nums := []int{2, 3, 1, 4, 0}
    res := challenge.BestRotation(nums)
    fmt.Println(res)
}

//
//  TestCountHighestScoreNodes
//  @Description: 2049. 统计最高分的节点数目
//  @param t
//
func TestCountHighestScoreNodes(t *testing.T) {
    nums := []int{-1, 2, 0, 2, 0}
    res := challenge.CountHighestScoreNodes(nums)
    fmt.Println(res)
}

//
//  TestFindRestaurant
//  @Description: 599. 两个列表的最小索引总和
//  @param t
//
func TestFindRestaurant(t *testing.T) {
    list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
    list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
    res := challenge.FindRestaurant(list1, list2)
    fmt.Println(res)
}

//
//  TestCountMaxOrSubsets
//  @Description: 2044. 统计按位或能得到最大值的子集数目
//  @param t
//
func TestCountMaxOrSubsets(t *testing.T) {
    nums := []int{3, 2, 1, 5}
    res := challenge.CountMaxOrSubsets(nums)
    fmt.Println(res)
}

//
//  TestLongestWord
//  @Description: 720. 词典中最长的单词
//  @param t
//
func TestLongestWord(t *testing.T) {
    words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
    res := challenge.LongestWord(words)
    fmt.Println(res)
}

//
//  TestWinnerOfGame
//  @Description: 2038. 如果相邻两个颜色均相同则删除当前颜色
//  @param t
//
func TestWinnerOfGame(t *testing.T) {
    colors := "AAABABB"
    res := challenge.WinnerOfGame(colors)
    fmt.Println(res)
}

//
//  TestFindKthDictOrderNumber
//  @Description: 440. 字典序的第K小数字
//  @param t
//
func TestFindKthDictOrderNumber(t *testing.T) {
    n, k := 2, 2
    res := challenge.FindKthDictOrderNumber(n, k)
    fmt.Println(res)
}

//
//  TestTrailingZeroes
//  @Description: 172. 阶乘后的零
//  @param t
//
func TestTrailingZeroes(t *testing.T) {
    n := 0
    res := challenge.TrailingZeroes(n)
    fmt.Println(res)
}

//
//  TestCalPoints
//  @Description: 682. 棒球比赛
//  @param t
//
func TestCalPoints(t *testing.T) {
    data := []string{"5", "2", "C", "D", "+"}
    res := challenge.CalPoints(data)
    fmt.Println(res)
}

//
//  TestMissingRolls
//  @Description: 2028. 找出缺失的观测数据
//  @param t
//
func TestMissingRolls(t *testing.T) {
    data := []int{3, 2, 4, 3}
    mean := 4
    n := 2
    res := challenge.MissingRolls(data, mean, n)
    fmt.Println(res)
}

//
//  TestReachingPoints
//  @Description: 780. 到达终点
//  @param t
//
func TestReachingPoints(t *testing.T) {
    sx, sy, tx, ty := 1, 1, 2, 2
    res := challenge.ReachingPoints(sx, sy, tx, ty)
    fmt.Println(res)
}

//
//  TestUniqueMorseRepresentations
//  @Description: 804. 唯一摩尔斯密码词
//  @param t
//
func TestUniqueMorseRepresentations(t *testing.T) {
    words := []string{"gin", "zen", "gig", "msg"}
    res := challenge.UniqueMorseRepresentations(words)
    fmt.Println(res)
}

//
//  TestGetTriangleRow
//  @Description: 119. 杨辉三角 II
//  @param t
//
func TestGetTriangleRow(t *testing.T) {
    rowIndex := 3
    res := challenge.GetTriangleRow(rowIndex)
    fmt.Println(res)
}

//
//  TestCountNumbersWithUniqueDigits
//  @Description: 357. 统计各位数字都不同的数字个数
//  @param t
//
func TestCountNumbersWithUniqueDigits(t *testing.T) {
    n := 2
    res := challenge.CountNumbersWithUniqueDigits(n)
    fmt.Println(res)
}
