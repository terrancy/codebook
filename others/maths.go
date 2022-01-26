package others

//
//  SQRT
//  @Description: NC32 求平方根
//  @param num
//  @return int
//
func SQRT(num int) int {
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

//
//  Gcd
//  @Description: NC151 最大公约数
//  @Description: 辗转相减法: A = a*X,B = b*X => A-B = (a - b)*X => X为公约数,ab互质
//  @param a
//  @param b
//  @return int
//
func Gcd(a int, b int) int {
    if a < b {
        a, b = b, a
    }
    if b == 0 {
        return a
    }
    return Gcd(b, a-b)
}

//
//  GcdII
//  @Description: NC151 最大公约数
//  @Description: 辗转相除法:
//  @param a
//  @param b
//  @return int
//
func GcdII(a int, b int) int {
    if a < b {
        a, b = b, a
    }
    if b == 0 {
        return a
    }
    return GcdII(b, a%b)
}
