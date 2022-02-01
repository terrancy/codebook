package array

import (
    "sort"
    "strconv"
)

func PrintMinNumber(numbers []int) string {
    str := make([]string, 0)
    for _, val := range numbers {
        str = append(str, string(val))
    }
    sort.Slice(str, func(i, j int) bool {
        return strconv.Itoa(numbers[i]) < strconv.Itoa(numbers[j])
    })
    res := ""
    for _, val := range numbers {
        res += string(val)
    }
    if res[0] == '0' {
        res = res[1:]
    }
    return res
}
