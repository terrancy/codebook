package challenge

//
// GetTriangleRow
// @Description: 119. 杨辉三角 II
// @param rowIndex
// @return []int
//
func GetTriangleRow(rowIndex int) []int {
    data := make([]int, 1)
    data[0] = 1
    temp := 0
    for i := 1; i <= rowIndex; i++ {
        temp = data[i-1] * (rowIndex - i + 1) / i
        data = append(data, temp)
    }
    return data
}
