package base

//////////////////////////////////////////////////////
// 冒泡排序、快速排序、插入排序、选择排序、堆排序
// 稳定性：排序后 2 个相等键值的顺序和排序之前它们的顺序相同
//////////////////////////////////////////////////////

//
//  BubbleSort
//  @Description: 冒泡排序 O(N2)
//  @Solution: 两两交换,最大的放后面
//  @param nums
//  @return []int
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
//  SelectSort
//  @Description: 选择排序
//  @Solution: 两两交换,最小的放前面
//  @param nums
//  @return []int
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
//  quickSort
//  @Description: 快速排序
//  @Solution: 以首个为基准分为两个数组,递归调用再合并.类似中根遍历.注意递归边界!!
//  @param nums
//  @return []int
//
func QuickSort(nums []int, ) []int {
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
    
    left = QuickSort(left)
    left = append(left, nums[0])
    right = QuickSort(right)
    
    return append(left, right...)
}

//
//  InsertSort
//  @Description: 插入排序
//  @Solution: 将待排序的元素插入已排序的序列中.
//  @param nums
//  @return []int
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

//
//  HeapSortASC
//  @Description: 堆排序
//  @param nums
//  @return []int
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
//  adjustHeap
//  @Description: 大根堆调整
//  @param nums
//  @param n
//  @param pos
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
//  HeapSortDESC
//  @Description: 小根堆
//  @param nums
//  @return []int
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
//  adjustHeapII
//  @Description: 对调整
//  @param nums
//  @param n
//  @param pos
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
