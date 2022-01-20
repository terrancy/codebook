package array

//
//  findMedianInTwoSortedArray NC36 在两个长度相等的排序数组中找到上中位数
//  @Description:
//  @param arr1
//  @param arr2
//  @return int
//
func findMedianInTwoSortedArray(arr1 []int, arr2 []int) int {
    n := len(arr1)
    res := 0
    for i, j := 0, 0; i+j < n; {
        if arr1[i] <= arr2[j] {
            res = arr1[i]
            i++
        } else {
            res = arr2[j]
            j++
        }
    }
    return res
}

//func main() {
//    arr1 := []int{1, 2, 3, 4}
//    arr2 := []int{3, 4, 5, 6}
//    res := findMedianInTwoSortedArray(arr1, arr2)
//    println(res)
//}
