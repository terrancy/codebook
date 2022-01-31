package test

import (
    "awesome/others"
    "fmt"
    "testing"
)

//
//  TestFoundOnceNumber
//  @Description: NC156 数组中只出现一次的数（其它数出现k次）
//  @param t
//
func TestFoundOnceNumber(t *testing.T) {
    data := []int{5, -4, 1, 1, 5, 1, 5}
    res := others.FoundOnceNumber(data, 3)
    fmt.Println(res)
}

//
//  TestFindNumsAppearOnce
//  @Description: NC75 数组中只出现一次的两个数字
//  @param t
//
func TestFindNumsAppearOnce(t *testing.T) {
    data := []int{1, 2, 3, 3, 2, 9}
    res := others.FindNumsAppearOnce(data)
    fmt.Println(res)
}

//
//  TestScaleConvert
//  @Description: 进制转换
//  @param t
//
func TestScaleConvert(t *testing.T) {
    res := others.ScaleConvert(-4, 3)
    fmt.Println(res)
}

//
//  TestTopKStrings
//  @Description: NC97 字符串出现次数的TopK问题
//  @param t
//
func TestTopKStrings(t *testing.T) {
    data := []string{"123", "123", "231", "32"}
    res := others.TopKStrings(data, 2)
    fmt.Println(res)
}

//
//  TestSqrt
//  @Description: NC32 求平方根
//  @param t
//
func TestSQRT(t *testing.T) {
    res := others.SQRT(5)
    fmt.Println(res)
}

//
//  TestIsPokerContinuous
//  @Description: NC63 扑克牌顺子
//  @param t
//
func TestIsPokerContinuous(t *testing.T) {
    data := []int{6, 0, 2, 0, 4}
    res := others.IsPokerContinuous(data)
    fmt.Println(res)
}

//
//  TestIsPalindrome
//  @Description: NC56 回文数字
//  @param t
//
func TestIsPalindrome(t *testing.T) {
    num := 121
    res := others.IsPalindromeII(num)
    fmt.Println(res)
}

//
//  TestUglyNumber
//  @Description: 丑数
//  @param t
//
func TestUglyNumber(t *testing.T) {
    res := others.UglyNumber(7)
    fmt.Println(res)
}

//
//  TestMinNumberInRotateArray
//  @Description: NC71 旋转数组的最小数字
//  @param t
//
func TestMinNumberInRotateArray(t *testing.T) {
    data := []int{3, 4, 5, 1, 2}
    res := others.MinNumberInRotateArray(data)
    fmt.Println(res)
}

//
//  TestTargetNumberInRotateArray
//  @Description: NC48 在旋转过的有序数组中寻找目标值
//  @param t
//
func TestTargetNumberInRotateArray(t *testing.T) {
    data := []int{1, 3}
    res := others.TargetNumberInRotateArrayII(data, 3)
    fmt.Println(res)
}

//
//  TestGcd
//  @Description: NC151 最大公约数
//  @param t
//
func TestGcd(t *testing.T) {
    res := others.GcdII(3, 4)
    fmt.Println(res)
}

//
//  TestMaxNums
//  @Description:NC111 最大数
//  @param t
//
func TestMaxNums(t *testing.T) {
    data := []int{2, 20, 23, 4, 8}
    res := others.MaxNums(data)
    fmt.Println(res)
}

//
//  TestNumberOfZeroFactorial
//  @Description: NC129 阶乘末尾0的数量
//  @param t
//
func TestNumberOfZeroFactorial(t *testing.T) {
    res := others.NumberOfZeroFactorial(1000000000)
    fmt.Println(res)
}

//
//  TestBigNumberPlus
//  @Description: NC1 大数加法
//  @param t
//
func TestBigNumberPlus(t *testing.T) {
    res := others.BigNumberPlus("1", "99")
    fmt.Println(res)
}

//
//  TestBigNumberProduct
//  @Description: NC10 大数乘法
//  @param t
//
func TestBigNumberProduct(t *testing.T) {
    source := "11"
    target := "99"
    res := others.BigNumberProduct(source, target)
    fmt.Printf("%s*%s=%s", source, target, res)
}

//
//  TestFindNthDigit
//  @Description: JZ44 数字序列中某一位的数字
//  @param t
//
func TestFindNthDigit(t *testing.T) {
    n := 13
    res := others.FindNthDigit(n)
    fmt.Println(res)
}

//
//  TestAddNumber
//  @Description: JZ65 不用加减乘除做加法
//  @param t
//
func TestAddNumber(t *testing.T) {
    a, b := 1001, 2020
    res := others.AddNumber(a, b)
    fmt.Println(res)
}

//
//  TestMyPow
//  @Description:
//  @param t
//
func TestMyPow(t *testing.T) {
    var x float64 = 100
    var n = 2
    res := others.MyPow(x, n)
    fmt.Println(res)
}
