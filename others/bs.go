package others

import "awesome/array"

//
//  binSearchII
//  @Description: NC105 二分查找-II
//  @Description: 重复数字的升序数组的二分查找第一个出现的数字
//  @param nums
//  @param target
//  @return int
//
func BinSearchII(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    if n == 1 {
        if nums[0] == target {
            return 0
        }
        return -1
    }
    l, r := 0, n-1
    mid := 0
    // 寻找重复数字有序数组的二分
    for l < r {
        mid = l + (r-l)>>1
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    if nums[r] == target {
        return r
    }
    return -1
}

//
//  GetNumberOfK
//  @Description: NC74 数字在升序数组中出现的次数
//  @Solution: 1.边界判断 2.左边界 3.有边界 4.返回
//  @param data
//  @param k
//  @return int
//
func GetNumberOfK(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        if nums[0] == target {
            return 1
        }
        return 0
    }
    // 左边界
    left := 0
    for l, r := 0, n-1; l < r; {
        mid := l + (r-l)>>1
        if nums[mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
        if l == r {
            left = r
        }
    }
    if nums[left] != target {
        return 0
    }
    // 右边界
    right := 0
    for l, r := 0, n-1; l < r; {
        mid := l + (r-l+1)>>1
        if nums[mid] <= target {
            l = mid
        } else {
            r = mid - 1
        }
        if l == r {
            right = l
        }
    }
    return right - left + 1
}

//
//  MinNumberInRotateArray
//  @Description: NC71 旋转数组的最小数字
//  @param rotateArray
//  @return int
//
func MinNumberInRotateArray(rotateArray []int) int {
    n := len(rotateArray)
    if n == 1 {
        return rotateArray[0]
    }
    l, r := 0, n-1
    for l < r {
        mid := l + (r-l)>>1
        if rotateArray[mid] > rotateArray[r] {
            l = mid + 1
        } else if rotateArray[mid] < rotateArray[r] {
            r = mid
        } else {
            r = r - 1
        }
    }
    return rotateArray[r]
}

//
//  TargetNumberInRotateArray
//  @Description: NC48 在旋转过的有序数组中寻找目标值
//  @Solution: 遍历
//  @param nums
//  @return int
//
func TargetNumberInRotateArray(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    if n == 1 && target != nums[0] {
        return -1
    }
    if n == 1 && target == nums[0] {
        return 0
    }
    for i := 0; i < n; i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
}

//
//  TargetNumberInRotateArrayII
//  @Description: NC48 在旋转过的有序数组中寻找目标值
//  @Solution: 二分查找.l<r:是否=;mid:是否+-1;mid和谁比;
//  @param nums
//  @param target
//  @return int
//
func TargetNumberInRotateArrayII(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    if n == 1 && target != nums[0] {
        return -1
    }
    if n == 1 && target == nums[0] {
        return 0
    }
    l, r := 0, n-1
    for l <= r {
        mid := l + (r-l)>>1
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[l] {
            if target > nums[mid] || target < nums[l] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else if nums[mid] < nums[r] {
            if target < nums[mid] || target > nums[r] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}

//
//  FindMedianInTwoSortedArrayBS
//  @Description: NC36 在两个长度相等的排序数组中找到上中位数
//  @param arr1
//  @param arr2
//  @return int
//
func FindMedianInTwoSortedArrayBS(arr1 []int, arr2 []int) int {
    return array.FindMedianInTwoSortedArrayBS(arr1, arr2)
}
