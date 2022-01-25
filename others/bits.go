package others

// 位运算相关

//
//  foundOnceNumber
//  @Description: NC156 数组中只出现一次的数（其它数出现k次）
//  @Solution:
//  @param arr
//  @param k
//  @return int
//
func FoundOnceNumber(arr []int, k int) int {
    var ans, negative = 0, false
    for i := 31; i >= 0; i-- {
        cnt := 0
        for _, v := range arr {
            // 数组中所有数计算第i位
            cnt += v >> i & 1
        }
        // 二进制转十进制运算，如果这个有余数，说明这个数第i位是存在的
        if cnt%k != 0 {
            if i == 31 {
                negative = true
                continue
            }
            ans += 1 << i
        }
    }
    if negative {
        ans = int(ans - 1<<31)
    }
    return ans
}

//
//  FindNumsAppearOnce
//  @Description: NC75 数组中只出现一次的两个数字
//  @param array
//  @return []int
//
func FindNumsAppearOnce(array []int) []int {
    // 求数组异或
    tmp := arrayXor(array)
    
    // 查找两个数的差异
    cnt := 0
    for tmp>>cnt&1 == 0 {
        cnt++
    }
    
    // 分组再异或求值
    left := make([]int, 0)
    right := make([]int, 0)
    for _, val := range array {
        if val>>cnt&1 == 1 {
            left = append(left, val)
        } else {
            right = append(right, val)
        }
    }
    a, b := arrayXor(left), arrayXor(right)
    if a < b {
        return []int{a, b}
    }
    return []int{b, a}
}

//
//  xorSum
//  @Description: 数组异或
//  @param []int
//  @return int
//
func arrayXor(nums []int) int {
    ans := 0
    for _, val := range nums {
        ans ^= val
    }
    return ans
}

//
//  scaleConvert
//  @Description: NC112 进制转换
//  @param num
//  @param scale
//  @return string
//
func ScaleConvert(num int, scale int) string {
    negative := ""
    if num < 0 {
        num *= -1
        negative = "-"
    }
    ans := ""
    scaleAlpha := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
    for num > 0 {
        a := num % scale
        ans = scaleAlpha[a] + ans
        num = (num - a) / scale
    }
    return negative + ans
}
