package dp

import "awesome"
//
//  maxProfit
//  @Description:
//  @param prices
//  @return int
//
func maxProfit(prices []int) int {
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
