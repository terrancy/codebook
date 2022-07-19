package others

import (
    "awesome"
    "sort"
    "strconv"
)

//
// SQRT
// @Description: NC32 求平方根
// @param num
// @return int
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
// Gcd
// @Description: NC151 最大公约数
// @Description: 辗转相减法: A = a*X,B = b*X => A-B = (a - b)*X => X为公约数,ab互质
// @param a
// @param b
// @return int
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
// GcdII
// @Description: NC151 最大公约数
// @Description: 辗转相除法:
// @param a
// @param b
// @return int
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

//
// MaxNums
// @Description: NC111 最大数
// @Description: 给定一个长度为n的数组nums，数组由一些非负整数组成，现需要将他们进行排列并拼接，每个数不可拆分，使得最后的结果最大
// @param nums
// @return string
//
func MaxNums(nums []int) string {
    str := make([]string, 0)
    for _, num := range nums {
        str = append(str, strconv.Itoa(num))
    }
    // 按照拼接的字符排序
    sort.Slice(str, func(i, j int) bool {
        str1 := str[i] + str[j]
        str2 := str[j] + str[i]
        num1, _ := strconv.Atoi(str1)
        num2, _ := strconv.Atoi(str2)
        return num1 > num2
    })
    res := ""
    for _, val := range str {
        res += val
    }
    if res[0] == '0' {
        return "0"
    }
    return res
}

//
// NumberOfZeroFactorial
// @Description: NC129 阶乘末尾0的数量
// @Description: 给定一个非负整数 n ，返回 n! 结果的末尾为 0 的数量
// @Tips: 有5因子的数比有2因子的数要少，所以我们就看能拆出来多少5就可以了，因为肯定能有足够数量的因子2来匹配
// @Tips: 这样的话 1\sim n1∼n 中有n/5个5的倍数，n/25个25的倍数，n/125个125的倍数!!
// @Tips: 25 = 5 * 5.对于5而言只算1个25.但是有2个5呢!!这里就要用25再回来.同理,125也是一样
// @param n
// @return int64
//
func NumberOfZeroFactorial(n int64) int64 {
    if n < 5 {
        return 0
    }
    var ans, factor int64 = 0, 5
    for n >= factor {
        ans += n / factor
        factor *= 5
    }
    return ans
}

//
// TwoSum
// @Description: NC1 大数加法
// @Description: 以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回
// @param data1
// @param data2
// @return []int
//
func BigNumberPlus(s, t string) string {
    n, m := len(s), len(t)
    if m == 0 {
        return s
    }
    if n == 0 {
        return t
    }
    k := awesome.MaxInt(n, m)
    data := make([]byte, k+1)
    plus := 0
    for i := 0; i < k; i++ {
        a, b := 0, 0
        if i < n {
            a = int(s[n-1-i] - '0')
        }
        if i < m {
            b = int(t[m-1-i] - '0')
        }
        sum := plus + a + b
        plus = sum / 10
        data[k-i] = byte(sum%10 + '0')
    }
    if plus == 1 {
        data[0] = '1'
        return string(data)
    }
    return string(data[1:])
}

//
// BigNumberProduct
// @Description: NC10 大数乘法
// @Description: 以字符串的形式读入两个数字，编写一个函数计算它们的乘积，以字符串形式返回。
// @param s
// @param t
// @return string
//
func BigNumberProduct(s string, t string) string {
    if s == "0" || t == "0" {
        return "0"
    }
    if s == "1" {
        return t
    }
    if t == "1" {
        return s
    }
    m := len(s)
    n := len(t)
    ns := make([]int, m)
    nt := make([]int, n)
    for i := 0; i < m; i++ {
        ns[i] = int(s[i] - '0')
    }
    for i := 0; i < n; i++ {
        nt[i] = int(t[i] - '0')
    }
    raw := make([]int, m+n)
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            raw[i+j+1] += ns[i] * nt[j]
        }
    }
    k := len(raw)
    res := make([]rune, k)
    carry := 0
    for i := k - 1; i >= 0; i-- {
        tmp := carry + raw[i]
        carry = tmp / 10
        res[i] = rune(tmp%10 + '0')
    }
    if res[0] == '0' {
        return string(res[1:])
    }
    return string(res)
}
