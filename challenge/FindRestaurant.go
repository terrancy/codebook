package challenge

//
// FindRestaurant
// @Description: 599. 两个列表的最小索引总和
// @param list1
// @param list2
// @return []string
//
func FindRestaurant(list1 []string, list2 []string) []string {
    dict := make(map[string]int)
    for idx, str := range list1 {
        dict[str] = idx
    }
    
    min := 2000
    minData := make([]string, 0)
    for idx, str := range list2 {
        if _, ok := dict[str]; ok {
            if dict[str]+idx < min {
                min = dict[str] + idx
                minData = minData[0:0]
            }
            if dict[str]+idx <= min {
                minData = append(minData, str)
            }
        }
    }
    return minData
}
