package main

func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    
    min := prices[0]
    left := make([]int, n, n)
    left[0] = 0
    for i := 1; i <= n-1; i++ {
        min = minInt(min, prices[i])
        left[i] = maxInt(left[i-1], prices[i]-min)
    }
    
    max := prices[n-1]
    right := make([]int, n, n)
    right[n-1] = 0
    for j := n - 2; j >= 0; j-- {
        max = maxInt(max, prices[j])
        right[j] = maxInt(right[j+1], max-prices[j])
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

func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minInt(a, b int) int {
    if a > b {
        return b
    }
    return a
}

//func main() {
//    prices := make([]int, n, n)
//    maxProfit(prices)
//}
