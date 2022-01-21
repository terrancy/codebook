package array

//
//  findMedianInTwoSortedArray NC36 在两个长度相等的排序数组中找到上中位数
//  @Description:
//  @param arr1
//  @param arr2
//  @return int
//
func FindMedianInTwoSortedArray(arr1 []int, arr2 []int) int {
    n := len(arr1)
    res := 0
    for i, j := 0, 0; i+j < n; {
        if arr1[i] <= arr2[j] {
            res = arr1[i]
            i++
        } else {
            res = arr2[j]
            j++
        }
    }
    return res
}

//
//  findMedianInTwoSortedArrayBS
//  @Description: NC36 在两个长度相等的排序数组中找到上中位数
//  @param arr1
//  @param arr2
//  @return int
//
func FindMedianInTwoSortedArrayBS(arr1 []int, arr2 []int) int {
    n := len(arr1)
    if n == 1 {
        return minInt(arr1[0], arr2[0])
    }
    l1, r1, l2, r2 := 0, n-1, 0, n-1
    for mid1, mid2 := 0, 0; l1 < r1; {
        mid1 = (l1 + r1) >> 1
        mid2 = (l2 + r2) >> 1
        offset := ((r1 - l1 + 1) & 1) ^ 1
        //println("<<")
        //println(r1 - l1 + 1)
        //println(offset)
        if arr1[mid1] == arr2[mid2] {
            return arr1[mid1]
        }
        if arr1[mid1] > arr2[mid2] {
            r1 = mid1
            l2 = mid2 + offset
        }
        if arr1[mid1] < arr2[mid2] {
            r2 = mid2
            l1 = mid1 + offset
        }
    }
    return minInt(arr1[l1], arr2[l2])
}

func minInt(a, b int) int {
    if a > b {
        return b
    }
    return a
}
