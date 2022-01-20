package array

//
//  findPeakElement
//  @Description: NC107 寻找峰值
//  @param num
//  @return int
//
func findPeakElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    
    for i := 0; i < n-1; i++ {
        if nums[i] > nums[i+1] {
            return i
        }
    }
    return n - 1
}

//
//  findPeakElementBinarySearch
//  @Description: 寻找峰值 - 二分查找 O(logN)
//  @param nums
//  @return int
//
func findPeakElementBinarySearch(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    left, right := 0, n-1
    mid := 0
    for left < right {
        mid = left + (right-left)>>1
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

//func main() {
//    data := []int{1, 2, 3, 4}
//    res := findPeakElementBinarySearch(data)
//    json, err := json.Marshal(res)
//    if err != nil {
//        fmt.Printf("错误:%v", err)
//    }
//    fmt.Println(string(json))
//}
