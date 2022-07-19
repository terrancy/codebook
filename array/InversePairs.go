package array

//
// InversePairs
// @Description: JZ51 数组中的逆序对 - 快排思路
// @param data
// @return int
//
func InversePairsQuickSort(data []int) int {
    n := len(data)
    if n < 2 {
        return 0
    }
    const Mod = 1000000007
    cnt := 0
    small := make([]int, 0)
    large := make([]int, 0)
    for _, val := range data[1:] {
        if val > data[0] {
            large = append(large, val)
        } else {
            small = append(small, val)
            cnt = (cnt + len(large) + 1) % Mod
        }
    }
    return (cnt + InversePairsQuickSort(large) + InversePairsQuickSort(small)) % Mod
}

//
// InversePairsMergeSort
// @Description: JZ51 数组中的逆序对 - 归并排序思路
// @param data
// @return int
//
func InversePairsMergeSort(data []int) int {
    return 1
}
