package challenge

//
//  IsBoomerang
//  @Description: 1037. 有效的回旋镖
//  @param points
//  @return bool
//
func IsBoomerang(points [][]int) bool {
    a, b, c := points[0], points[1], points[2]
    x := (b[1] - a[1]) * (c[0] - a[0])
    y := (c[1] - a[1]) * (b[0] - a[0])
    if x != y {
        return true
    }
    return false
}
