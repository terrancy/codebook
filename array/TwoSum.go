package array

//
//  TwoSum
//  @Description: NC61 两数之和
//  @param numbers
//  @param target
//  @return []int
//
func TwoSum(numbers []int, target int) []int {
    n := len(numbers)
    data := make(map[int]int)
    
    for i := 0; i < n; i++ {
        number := numbers[i]
        if _, ok := data[number]; !ok {
            data[numbers[i]] = i
        }
        minus := target - number
        if _, ok := data[minus]; ok {
            return []int{data[minus] + 1, i + 1}
        }
    }
    return []int{}
}

//
//  towSumII
//  @Description: 将数组键值反转,查看差值是否存在
//  @param numbers
//  @param target
//  @return []int
//
func TowSumII(numbers []int, target int) []int {
    n := len(numbers)
    data := make(map[int]int)
    
    for i := 0; i < n; i++ {
        if _, ok := data[numbers[i]]; !ok {
            data[numbers[i]] = i
        }
    }
    
    for i := 0; i < n; i++ {
        number := numbers[i]
        minus := target - number
        if _, ok := data[minus]; ok {
            return []int{i + 1, data[minus] + 1}
        }
    }
    return []int{}
}
