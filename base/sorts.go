package base

// ////////////////////////////////////////////////////
// 冒泡排序、快速排序、插入排序、选择排序、堆排序
// 稳定性：排序后 2 个相等键值的顺序和排序之前它们的顺序相同
// ////////////////////////////////////////////////////

// BubbleSort
// @Description: 冒泡排序 O(N2)
// @Solution: 两两交换,最大的放后面
// @param nums
// @return []int
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

// SelectSort
// @Description: 选择排序 O(N2)
// @Solution: 每轮选出最小值索引,放到未排序序列首位
// @param nums
// @return []int
func SelectSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

// InsertSort
// @Description: 插入排序
// @Solution: 将待排序的元素插入已排序的序列中.
// @param nums
// @return []int
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

// —— 模块一: 朴素快排（便于理解快排思想） ——

// QuickSortOld
// @Description: 朴素快速排序
// @Solution: 以首个元素为基准,分为两个子数组,递归再合并
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

// —— 模块二: 快排核心（仅调度,不含优化细节） ——

// QuickSort
// @Description: 快排优化版
// @Solution: 三数取中选基准 + 尾递归 + 双路分区
func QuickSort(nums []int, left int, right int) {
	for left < right {
		pivot := selectPivot(nums, left, right)
		mid := partitionLomuto(nums, left, right, pivot)
		if mid-left < right-mid {
			QuickSort(nums, left, mid-1)
			left = mid + 1
		} else {
			QuickSort(nums, mid+1, right)
			right = mid - 1
		}
	}
}

// QuickSortII
// @Description: 三路快排（处理重复值,减少无效分区）
// @Solution: 三数取中选基准 + 三路分区（荷兰国旗）
func QuickSortII(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := selectPivot(nums, left, right)
	leftMid, rightMid := threeWayPartition(nums, left, right, pivot)
	QuickSortII(nums, left, leftMid)
	QuickSortII(nums, rightMid, right)
}

// —— 优化组件一: 基准选择（三数取中） ——

// selectPivot
// @Description: 三数取中选基准,并交换到首位
// @Solution: 取 left/right/mid 三者的中位数,交换到 nums[left] 位置
// @return int 选中的基准值
func selectPivot(nums []int, left int, right int) int {
	mid := left + (right-left)>>1
	pivot := threeSumMedian(nums[left], nums[right], nums[mid])
	idx := getPivotIdx(nums, left, right, mid, pivot)
	nums[left], nums[idx] = nums[idx], nums[left]
	return pivot
}

// —— 优化组件二: 双路分区（Lomuto 双向扫描） ——

// partitionLomuto
// @Description: Lomuto 双路分区
// @Solution: 基准已在首位,双向扫描将数组分为 小于基准 | 大于基准 两部分
// @return int 基准最终所在的索引
func partitionLomuto(nums []int, left int, right int, pivot int) int {
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

// —— 优化组件三: 三路分区（荷兰国旗问题） ——

// threeWayPartition
// @Description: 三路分区（荷兰国旗问题）
// @Solution: 将数组划分为 小于基准 | 等于基准 | 大于基准 三部分
// @return (int, int) 左右边界,使得 nums[left:right+1] 全等于 pivot
func threeWayPartition(nums []int, left int, right int, pivot int) (int, int) {
	first, last := left, right
	leftLen, rightLen := 0, 0
	leftPos, rightPos := left, right

	// —— 阶段一: 双向扫描分区 ——
	// 右指针从右向左: 跳过 >= pivot 的元素,遇到 ==pivot 的暂存到 rightPos
	// 左指针从左向右: 跳过 <= pivot 的元素,遇到 ==pivot 的暂存到 leftPos
	// 扫描结束后: [==pivot | <pivot | >pivot | ==pivot]
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

	// —— 阶段二: 基准归位 + 重复值归位 ——
	// 将基准放到最终位置,然后将两侧暂存的 ==pivot 元素移到基准左右
	// 左侧归位: 将 [first, leftPos) 区间的 ==pivot 元素交换到基准左边
	i, j := left-1, first
	for j < leftPos && nums[i] != pivot {
		nums[i], nums[j] = nums[j], nums[i]
		i--
		j++
	}
	// 右侧归位: 将 (rightPos, last] 区间的 ==pivot 元素交换到基准右边
	i, j = left+1, last
	for j > rightPos && nums[i] != pivot {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	// 返回等于基准区间的左右边界: nums[left-1-leftLen ... left+1+rightLen] 全等于 pivot
	return left - 1 - leftLen, left + 1 + rightLen
}

// —— 辅助函数 ——

// getPivotIdx
// @Description: 根据目标值返回在三数中的索引
func getPivotIdx(nums []int, a, b, c, target int) int {
	for _, idx := range []int{a, b, c} {
		if target == nums[idx] {
			return idx
		}
	}
	return -1
}

// threeSumMedian
// @Description: 三数取中
func threeSumMedian(a, b, c int) int {
	if b > a {
		a, b = b, a
	}
	if a < c {
		return a
	}
	if b > c {
		return b
	}
	return c
}

// ////////////////////////////////////////////////////
// 堆排序
// ////////////////////////////////////////////////////

// HeapSortASC
// @Description: 堆排序
// @param nums
// @return []int
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

// adjustHeap
// @Description: 大根堆调整。每次获取最大值
// @param nums
// @param n
// @param pos
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

// HeapSortDESC
// @Description: 小根堆
// @param nums
// @return []int
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

// adjustHeapII
// @Description: 堆调整。每次获取最小值
// @param nums
// @param n
// @param pos
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
		adjustHeap(nums, n, i)
	}
	for i, cnt := n-1, 0; i > 0 && cnt < k; i, cnt = i-1, cnt+1 {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		adjustHeap(nums, n, 0)
	}
	m := len(nums)
	if m < k {
		return nums
	}
	return nums[m-k:]
}