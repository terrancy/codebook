package base

import "fmt"

//
// testRune
// @Description: rune测试
//
func StructRune() {
    str := "这是2022"
    for _, val := range []rune(str) {
        println(string(val))
    }
    str1 := []rune(str)
    fmt.Println(str1)
    fmt.Println(string(str1[:1]))
}
