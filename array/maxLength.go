package main

func maxLength(arr []int) int {
    // write code here
    n := len(arr)
    if n < 2 {
        return n
    }
    res := 0
    data := make(map[int]int)
    for left, right, val := -1, 0, 0; right < n; right++ {
        val = arr[right]
        if _, ok := data[val]; ok {
            left = maxInt(left, data[val])
        }
        data[val] = right
        res = maxInt(res, right-left)
    }
    return res
}

func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//func main() {
//    data := []int{2, 2, 3, 4, 3}
//    ret := maxLength(data)
//    println(ret)
//}
