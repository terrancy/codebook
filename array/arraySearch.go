package array

//
//  arraySearch: NC29 二维数组中的查找
//  @Description: 在一个二维数组array中（每个一维数组的长度相同），
//  每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//  请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数
//  @param target int
//  @param data array
//  @return bool
//
func arraySearch(target int, data [][]int) bool {
    n := len(data)
    m := len(data[0])
    for i, j := n-1, 0; i >= 0 && j < m;
    {
        if target == data[i][j] {
            return true
        }
        
        if target > data[i][j] {
            j++
            continue
        }
        
        if target < data[i][j] {
            i--
            continue
        }
    }
    return false
}

//func main() {
//    target := 10
//    data := [][]int{
//        {1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15},
//    }
//    res := arraySearch(target, data)
//    println(res)
//}
