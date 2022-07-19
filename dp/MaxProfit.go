package dp

import "awesome"

// 连续子数组的最大和

//
// MaxSubArray
// @Description: NC19 连续子数组的最大和
// @param nums
// @return int
//
func MaxSubArray(nums []int) int {
    max := nums[0]
    sum := 0
    for _, val := range nums {
        sum = awesome.MaxInt(sum+val, val)
        max = awesome.MaxInt(max, sum)
    }
    return max
}

//
// MaxSubArrayII
// @Description: JZ85 连续子数组的最大和(二)
// @Description: 找到一个具有最大和且数组长度最长的连续子数组
// @param nums
// @return []int
//
func MaxSubArrayII(nums []int) []int {
    max := nums[0]
    sum := 0
    left, right, n := 0, 0, 0
    for idx, val := range nums {
        if sum < 0 {
            sum = val
            left = idx
        } else {
            sum = sum + val
        }
        if max > sum {
            continue
        }
        if sum > max || idx-left > n {
            right = idx
            n = right - left
        }
        max = sum
    }
    left = right - n
    return nums[left : right+1]
}

// 买卖股票的最好时机

//
// MaxProfit
// @Description: JZ63 买卖股票的最好时机(一)
// @Description: 你可以买入一次股票和卖出一次股票,请根据这个价格数组，返回买卖股票能获得的最大收益
// @Solution: 利润是整数，max初始值为0
// @param prices
// @return int
//
func MaxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    // 利润是整数，max初始值为0
    sum, max := 0, 0
    for i := 0; i < n-1; i++ {
        val := prices[i+1] - prices[i]
        sum = awesome.MaxInt(sum+val, val)
        max = awesome.MaxInt(max, sum)
    }
    return max
}

//
// MaxProfitII
// @Description: 122. 买卖股票的最佳时机 II
// @Description: 你在任何时候最多只能持有一股股票,可以多次卖出再买入
// @param prices
// @return int
//
func MaxProfitII(prices []int) int {
    return 0
}

//
// MaxProfitIII
// @Description: NC135 买卖股票的最好时机(三)
// @param prices
// @return int
//
func MaxProfitIII(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    
    min := prices[0]
    left := make([]int, n, n)
    left[0] = 0
    for i := 1; i <= n-1; i++ {
        min = awesome.MinInt(min, prices[i])
        left[i] = awesome.MaxInt(left[i-1], prices[i]-min)
    }
    
    max := prices[n-1]
    right := make([]int, n, n)
    right[n-1] = 0
    for j := n - 2; j >= 0; j-- {
        max = awesome.MaxInt(max, prices[j])
        right[j] = awesome.MaxInt(right[j+1], max-prices[j])
    }
    
    max = 0
    var temp int
    for k := 0; k < n-1; k++ {
        temp = left[k] + right[k]
        if max < temp {
            max = temp
        }
    }
    
    return max
}
