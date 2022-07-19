package array

import (
    "sort"
)

//
// ThreeSum
// @Description: NC54 数组中相加和为0的三元组
// @param num
// @return [][]int
//
func ThreeSum(num []int) [][]int {
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
                
                // 继续寻找
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

//
// ThreeSumII
// @Description: NC54 数组中相加和为0的三元组
// @Solution: 1.排序判断最小值; 2. 去重
// @param num
// @return [][]int
//
func ThreeSumII(num []int) [][]int {
    n := len(num)
    // 边界过滤
    if n < 3 {
        return nil
    }
    // 排序
    sort.Ints(num)
    ans := make([][]int, 0)
    // 遍历
    for k := 0; k < n-2; k++ {
        // 去重
        for k > 0 && k < n-2 && num[k-1] == num[k] {
            k++
        }
        for i, j := k+1, n-1; i < j; {
            sum := num[i] + num[j] + num[k]
            // 相等入栈
            if sum == 0 {
                ans = append(ans, []int{num[k], num[i], num[j]})
            }
            // 小于右移
            if sum <= 0 {
                i++
                for i < j && num[i] == num[i-1] {
                    i++
                }
            }
            // 大于左移
            if sum >= 0 {
                j--
                for i < j && num[j] == num[j+1] {
                    j--
                }
            }
        }
    }
    return ans
}
