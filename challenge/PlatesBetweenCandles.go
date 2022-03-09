package challenge

//
//  PlatesBetweenCandles
//  @Description: 2055. 蜡烛之间的盘子
//  @param s string
//  @param queries [][]int
//  @return []int
//
func PlatesBetweenCandles(s string, queries [][]int) []int {
    bytes := []byte(s)
    n, m := len(s), len(queries)
    ans, sum := make([]int, m), make([]int, n+1)
    list := make([]int, 0)
    for i := 0; i < n; i++ {
        if bytes[i] == '|' {
            list = append(list, i)
        }
        temp := 0
        if bytes[i] == '*' {
            temp = 1
        }
        sum[i+1] = sum[i] + temp
    }
    if len(list) == 0 {
        return ans
    }
    
    for i := 0; i < m; i++ {
        a, b := queries[i][0], queries[i][1]
        c, d := -1, -1
        l, r := 0, len(list)-1
        // 找到 a 右边最近的蜡烛
        for l < r {
            mid := (l + r) >> 1
            if list[mid] >= a {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if list[r] >= a {
            c = list[r]
        } else {
            continue
        }
        // 找到 b 左边最近的蜡烛
        l, r = 0, len(list)-1
        for l < r {
            mid := (l + r + 1) >> 1
            if list[mid] <= b {
                l = mid
            } else {
                r = mid - 1
            }
        }
        if list[r] <= b {
            d = list[r]
        } else {
            continue
        }
        if c <= d {
            ans[i] = sum[d+1] - sum[c]
        }
    }
    return ans
}
