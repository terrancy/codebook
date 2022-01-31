package array

//
//  FindRepeatNumber
//  @Description: JZ3 数组中重复的数字
//  @param numbers
//  @return int
//
func FindRepeatNumber(numbers []int) int {
    n := len(numbers)
    for i := 0; i < n; {
        if numbers[i] == i {
            i += 1
            continue
        }
        if numbers[numbers[i]] == numbers[i] {
            return numbers[i]
        }
        numbers[numbers[i]], numbers[i] = numbers[i], numbers[numbers[i]]
    }
    return -1
}
