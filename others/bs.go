package others

//
//  MinNumberInRotateArray
//  @Description: NC71 旋转数组的最小数字
//  @param rotateArray
//  @return int
//
func MinNumberInRotateArray(rotateArray []int) int {
    n := len(rotateArray)
    if n == 1 {
        return rotateArray[0]
    }
    l, r := 0, n-1
    for l < r {
        mid := l + (r-l)>>1
        if rotateArray[mid] > rotateArray[r] {
            l = mid + 1
        } else if rotateArray[mid] < rotateArray[r] {
            r = mid
        } else {
            r = r - 1
        }
    }
    return rotateArray[r]
}
