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
