package main

import "fmt"

//func main() {
//    data := []int{3, 3, 0, 3}
//    res := permuteUnique(data)
//    fmt.Println(res)
//}

func main() {
    temp := []int{0, 1, 2, 3, 4, 5}
    data := make([][]int, 0)
    data = append(data, temp)
    temp = temp[1:]
    fmt.Println(temp, data)
}
