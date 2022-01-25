package others

import (
    "sort"
    "strconv"
)

// 模拟

//
//  IsPokerContinuous
//  @Description: NC63 扑克牌顺子
//  @param numbers
//  @return bool
//
func IsPokerContinuous(numbers []int) bool {
    n := len(numbers)
    if n > 5 {
        return false
    }
    sort.Ints(numbers)
    cnt := 0
    for i := 0; i < n-1; i++ {
        if numbers[i] == 0 {
            // 大小王统计
            cnt++
        } else if numbers[i] == numbers[i+1] {
            // 判重
            return false
        } else {
            // 步数
            cnt -= numbers[i+1] - numbers[i] - 1
        }
    }
    return cnt >= 0
}

//
//  IsPalindrome
//  @Description: NC56 回文数字
//  @Solution: 利用字符串的思路
//  @param x
//  @return bool
//
func IsPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    str := strconv.Itoa(x)
    n := len(str)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true
}

//
//  IsPalindromeII
//  @Description: NC56 回文数字
//  @Solution:
//  @param x
//  @return bool
//
func IsPalindromeII(x int) bool {
    if x < 0 {
        return false
    }
    if x == 0 {
        return true
    }
    y := 1
    // 获取最高位
    for x/y > 10 {
        y *= 10
    }
    // 最高低两位比较
    for x > 0 {
        a := x / y
        b := x % 10
        x = x % y / 10
        y = y / 100
        if a != b {
            return false
        }
    }
    return true
}
