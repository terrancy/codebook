package challenge

//
// BestRotation
// @Description: 798. 得分最高的最小轮调798. 得分最高的最小轮调
// @param nums
// @return int
//
func BestRotation(nums []int) int {
    m := 100001
    c := make([]int, m)
    n := len(nums)
    for i := 0; i < n; i++ {
        a := (i - (n - 1) + n) % n
        b := (i - nums[i] + n) % n
        if a <= b {
            addInt(a, b, c)
        } else {
            addInt(0, b, c)
            addInt(a, n-1, c)
        }
    }
    for i := 1; i <= n; i++ {
        c[i] += c[i-1]
    }
    res := 0
    for i := 1; i <= n; i++ {
        if c[i] > c[res] {
            res = i
        }
    }
    return res
}

func addInt(l, r int, c []int) {
    c[l] += 1
    c[r+1] -= 1
}
