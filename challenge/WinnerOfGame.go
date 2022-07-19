package challenge

//
// WinnerOfGame
// @Description: 2038. 如果相邻两个颜色均相同则删除当前颜色
// @param colors
// @return bool
//
func WinnerOfGame(colors string) bool {
    n := len(colors)
    bytes := []byte(colors)
    diff := 0
    for i := 1; i < n-1; i++ {
        if bytes[i] == 'A' && bytes[i-1] == 'A' && bytes[i+1] == 'A' {
            diff++
        }
        if bytes[i] == 'B' && bytes[i-1] == 'B' && bytes[i+1] == 'B' {
            diff--
        }
    }
    return diff > 0
}
