package main

func findMedianInTwoSortedArrayBS(arr1 []int, arr2 []int) int {
    n := len(arr1)
    if n == 1 {
        return minInt(arr1[0], arr2[0])
    }
    l1, r1, l2, r2 := 0, n-1, 0, n-1
    for mid1, mid2 := 0, 0; l1 < r1; {
        mid1 = (l1 + r1) >> 1
        mid2 = (l2 + r2) >> 1
        offset := ((r1 - l1 + 1) & 1) ^ 1
        println("<<")
        println(r1 - l1 + 1)
        println(offset)
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

//func main() {
//    arr1 := []int{1, 9, 19}
//    arr2 := []int{2, 5, 10}
//    res := findMedianInTwoSortedArrayBS(arr1, arr2)
//    println(res)
//}
