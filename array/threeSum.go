package array

import (
    "sort"
)

func threeSum(num []int) [][]int {
    n := len(num)
    if n < 3 {
        return nil
    }
    // 排序
    sort.Ints(num)
    
    i, j := 0, 0
    ans := make([][]int, 0)
    for k, sum := 0, 0; k < n-2; k++ {
        if num[k] > 0 {
            break
        }
        // 重复值跳过
        for k > 0 && k < n-2 {
            if num[k] != num[k-1] {
                break
            }
            k++
        }
        i = k + 1
        j = n - 1
        sum = 0
        for i < j {
            sum = num[k] + num[i] + num[j]
            if sum < 0 {
                // 重复值跳过
                for i < j {
                    i++
                    if num[i-1] != num[i] {
                        break
                    }
                }
            } else if sum > 0 {
                // 重复值跳过
                for i < j {
                    j--
                    if num[j+1] != num[j] {
                        break
                    }
                }
            } else {
                // 存储
                ans = append(ans, []int{num[k], num[i], num[j]})
                
                //继续寻找
                for i < j {
                    i++
                    if num[i-1] != num[i] {
                        break
                    }
                }
                for i < j {
                    j--
                    if num[j+1] != num[j] {
                        break
                    }
                }
            }
        }
    }
    return ans
}

func threeSumII(num []int) [][]int {
    n := len(num)
    // 边界过滤
    if n < 3 {
        return nil
    }
    // 排序
    sort.Ints(num)
    var i, j, sum int
    ans := make([][]int, 0)
    // 遍历
    for k := 0; k < n-2; k++ {
        // 去重
        for k > 0 && k < n-2 {
            if num[k-1] != num[k] {
                break
            }
            k++
        }
        i = k + 1
        j = n - 1
        for i < j {
            sum = num[i] + num[j] + num[k]
            // 相等入栈
            if sum == 0 {
                ans = append(ans, []int{num[k], num[i], num[j]})
            }
            // 小于右移
            if sum <= 0 {
                for i < j {
                    i++
                    if num[i-1] != num[i] {
                        break
                    }
                }
            }
            // 大于左移
            if sum >= 0 {
                for i < j {
                    j--
                    if num[j+1] != num[j] {
                        break
                    }
                }
            }
        }
    }
    return ans
}

//func main() {
//    data := []int{-10, 0, 10, 20, -10, -40}
//    //data := []int{0, 0, 0, 0}
//    res := threeSumII(data)
//    json, err := json.Marshal(res)
//    if err != nil {
//        println("错误")
//    } else {
//        println(string(json))
//    }
//}
