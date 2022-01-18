package main

import "encoding/json"

func main() {
    num := []int{2, 3, 4, 2, 6, 2, 5, 1}
    ans := maxInWindows(num, 3)
    json, err := json.Marshal(ans)
    if err != nil {
        println("错误")
    } else {
        println(string(json))
    }
}
