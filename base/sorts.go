package base

// ////////////////////////////////////////////////////
// 冒泡排序、快速排序、插入排序、选择排序、堆排序
// 稳定性：排序后 2 个相等键值的顺序和排序之前它们的顺序相同
// ////////////////////////////////////////////////////

//
// BubbleSort
// @Description: 冒泡排序 O(N2)
// @Solution: 两两交换,最大的放后面
// @param nums
// @return []int
//
func BubbleSort(nums []int) []int {
    n := len(nums)
    for i := n - 1; i >= 0; i-- {
        for j := 0; j < i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
    return nums
}

//
// SelectSort
// @Description: 选择排序
// @Solution: 两两交换,最小的放前面
// @param nums
// @return []int
//
func SelectSort(nums []int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    return nums
}

//
// InsertSort
// @Description: 插入排序
// @Solution: 将待排序的元素插入已排序的序列中.
// @param nums
// @return []int
//
func InsertSort(nums []int) []int {
    n := len(nums)
    for i := 1; i < n; i++ {
        preIdx := i - 1
        cur := nums[i]
        for preIdx >= 0 && nums[preIdx] > cur {
            nums[preIdx+1] = nums[preIdx]
            preIdx--
        }
        nums[preIdx+1] = cur
    }
    return nums
}

// ////////////////////////////////////////////////////
// 快速排序
// ////////////////////////////////////////////////////

//
// quickSort
// @Description: 朴素快速排序
// @Solution: 以首个为基准分为两个数组,递归调用再合并.类似中根遍历.注意递归边界!!
// @param nums
// @return []int
//
func QuickSortOld(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return nil
    }
    left, right := make([]int, 0), make([]int, 0)
    for i := 1; i < n; i++ {
        if nums[i] < nums[0] {
            left = append(left, nums[i])
        } else {
            right = append(right, nums[i])
        }
    }
    
    left = QuickSortOld(left)
    left = append(left, nums[0])
    right = QuickSortOld(right)
    
    return append(left, right...)
}

//
// QuickSortII
// @Description: 快排优化
// @Solution: 1.基准选中 2.尾递归优化 3.重复数值的处理
// @param nums
// @param left 左边界
// @param right 有边界
//
func QuickSort(nums []int, left int, right int) {
    for left < right {
        mid := partition(nums, left, right)
        QuickSort(nums, left, mid-1)
        left = mid + 1
    }
}

//
// partition
// @Description: 获取基准值
// @param nums
// @param left
// @param right
// @return int
//
func partition(nums []int, left int, right int) int {
    // 三数取中
    mid := left + (right-left)>>1
    pivot := threeSumMedian(nums[left], nums[right], nums[mid])
    // 替换
    idx := getPivotIdx(nums, left, right, mid, pivot)
    nums[left], nums[idx] = nums[idx], nums[left]
    // 基准分区，重复数值的处理
    for left < right {
        for left < right && pivot <= nums[right] {
            right--
        }
        nums[left] = nums[right]
        for left < right && pivot >= nums[left] {
            left++
        }
        nums[right] = nums[left]
    }
    nums[left] = pivot
    return left
}

//
// QuickSortII
// @Description: 快排优化
// @Solution: 1.基准选中 2.尾递归优化 3.重复数值的处理
// @param nums
// @param left 左边界
// @param right 有边界
//
func QuickSortII(nums []int, left int, right int) {
    partitionII(nums, left, right)
    // for left < right {
    //   // 重复值处理
    //   leftMid, rightMid := partitionII(nums, left, right)
    //   QuickSortII(nums, left, leftMid)
    //   left = rightMid
    // }
}

//
// partitionII
// @Description: 重复数值的优化
// @Link: https://www.cnblogs.com/vipchenwei/p/7460293.html
// @param nums
// @param left
// @param right
// @return int
//
func partitionII(nums []int, left int, right int) (int, int) {
    first, last := left, right
    // 三数取中
    mid := left + (right-left)>>1
    pivot := threeSumMedian(nums[left], nums[right], nums[mid])
    // 基准分区，重复数值的处理
    leftLen, rightLen := 0, 0
    leftPos, rightPos := left, right
    nums[left], pivot = pivot, nums[left]
    for left < right {
        for left < right && pivot <= nums[right] {
            if pivot == nums[right] {
                nums[rightPos], nums[right] = nums[right], nums[rightPos]
                rightPos--
                rightLen++
            }
            right--
        }
        nums[left] = nums[right]
        for left < right && pivot >= nums[left] {
            if pivot == nums[left] {
                nums[leftPos], nums[left] = nums[left], nums[leftPos]
                leftPos++
                leftLen++
            }
            left++
        }
        nums[right] = nums[left]
    }
    nums[left] = pivot
    
    i, j := left-1, first
    for j < leftPos && nums[i] != pivot {
        nums[i], nums[j] = nums[j], nums[i]
        i--
        j++
    }
    i, j = left+1, last
    for j > rightPos && nums[i] != pivot {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
    
    return left - 1 - leftLen, left + 1 + rightLen
}

func getPivotIdx(nums []int, a, b, c, target int) int {
    for _, idx := range []int{a, b, c} {
        if target == nums[idx] {
            return idx
        }
    }
    return -1
}

//
// threeSumMedian
// @Description: 三数取中
// @param a
// @param b
// @param c
// @return int
//
func threeSumMedian(a, b, c int) int {
    if b > a {
        a, b = b, a
    }
    if a < c {
        return a
    } else {
        if b > c {
            return b
        } else {
            return c
        }
    }
}

// ////////////////////////////////////////////////////
// 堆排序
// ////////////////////////////////////////////////////

//
// HeapSortASC
// @Description: 堆排序
// @param nums
// @return []int
//
func HeapSortASC(nums []int) []int {
    n := len(nums)
    for i := n/2 - 1; i >= 0; i-- {
        adjustHeap(nums, n, i)
    }
    for i := n - 1; i > 0; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        n--
        adjustHeap(nums, n, 0)
    }
    return nums
}

//
// adjustHeap
// @Description: 大根堆调整。每次获取最大值
// @param nums
// @param n
// @param pos
//
func adjustHeap(nums []int, n int, pos int) {
    largeIndex := pos
    if 2*pos+1 < n && nums[2*pos+1] > nums[largeIndex] {
        largeIndex = 2*pos + 1
    }
    if 2*pos+2 < n && nums[2*pos+2] > nums[largeIndex] {
        largeIndex = 2*pos + 2
    }
    if largeIndex != pos {
        nums[pos], nums[largeIndex] = nums[largeIndex], nums[pos]
        adjustHeap(nums, n, largeIndex)
    }
}

//
// HeapSortDESC
// @Description: 小根堆
// @param nums
// @return []int
//
func HeapSortDESC(nums []int) []int {
    n := len(nums)
    for i := n/2 - 1; i >= 0; i-- {
        adjustHeapII(nums, n, i)
    }
    // 循环交换根节点和最后一个节点
    for i := n - 1; i > 0; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        // 每次交换后需要重新调整
        n--
        adjustHeapII(nums, n, 0)
    }
    return nums
}

//
// adjustHeapII
// @Description: 堆调整。每次获取最小值
// @param nums
// @param n
// @param pos
//
func adjustHeapII(nums []int, n int, pos int) {
    smallIdx := pos
    leftIdx, rightIdx := 2*smallIdx+1, 2*smallIdx+2
    if leftIdx < n && nums[smallIdx] > nums[leftIdx] {
        smallIdx = leftIdx
    }
    if rightIdx < n && nums[smallIdx] > nums[rightIdx] {
        smallIdx = rightIdx
    }
    if smallIdx != pos {
        nums[smallIdx], nums[pos] = nums[pos], nums[smallIdx]
        adjustHeapII(nums, n, smallIdx)
    }
}
func FindLastKth(nums []int, k int) []int {
    n := len(nums)
    for i := n/2 - 1; i >= 0; i-- {
        adjustHeapII(nums, n, i)
    }
    // 循环交换根节点和最后一个节点
    for i, cnt := n-1, 0; i > 0 && cnt < k; i, cnt = i-1, cnt+1 {
        nums[0], nums[i] = nums[i], nums[0]
        // 每次交换后需要重新调整
        n--
        adjustHeapII(nums, n, 0)
    }
    m := len(nums)
    if m < k {
        return nums
    }
    return nums[m-k:]
}
