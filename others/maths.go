package others

//
//  Sqrt
//  @Description: NC32 求平方根
//  @param num
//  @return int
//
func Sqrt(num int) int {
    if num < 2 {
        return num
    }
    l, r := 1, num
    for l <= r {
        mid := l + (r-l)>>1
        if mid <= num/mid && (mid+1) > num/(mid+1) {
            return mid
        } else if mid > num/mid {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return 1
}
