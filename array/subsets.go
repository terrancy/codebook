package array

//
//  subsets
//  @Description:
//  @param nums
//  @return [][]int
//
func subsets(nums []int) [][]int {
    n := len(nums)
    if n == 0 {
        return [][]int{[]int{}}
    }
    item := nums[n-1]
    res := subsets(nums[0 : n-1])
    
    ln := len(res)
    for i := 0; i < ln; i++ {
        tmp := make([]int, len(res[i]))
        copy(tmp, res[i])
        tmp = append(tmp, item)
        res = append(res, tmp)
    }
    
    return res
}

//func main() {
//    data := []int{1, 2, 3}
//    res := subsets(data)
//    json, err := json.Marshal(res)
//    if err != nil {
//        print("error")
//    } else {
//        println(string(json))
//    }
//}
