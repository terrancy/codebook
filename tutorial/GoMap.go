package tutorial

import "fmt"

func ForMap(dic map[int]int) {
    fmt.Println(dic)
    dic[1] = 119
    fmt.Printf("%p\n", &dic)
}
