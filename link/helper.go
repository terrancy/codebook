package main

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

func arrayReverse(data []int) []int {
    n := len(data)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
    return data
}
