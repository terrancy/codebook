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
		if right-left < 10 {
			InsertSort(nums[left : right+1])
			return
		}
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
// @Solution: 三数取中选基准 + 三路分区（荷兰国旗）+ 尾递归优化 + 小数组插入排序
func QuickSortII(nums []int, left int, right int) {
	for left < right {
		if right-left < 10 {
			InsertSort(nums[left : right+1])
			return
		}
		pivot := selectPivot(nums, left, right)
		leftMid, rightMid := threeWayPartition(nums, left, right, pivot)
		if leftMid-left < right-rightMid {
			QuickSortII(nums, left, leftMid)
			left = rightMid
		} else {
			QuickSortII(nums, rightMid, right)
			right = leftMid
		}
	}
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
// @Description: 三路分区（荷兰国旗 DNF 算法）
// @Solution: 一次遍历, 将数组划分为 [小于基准 | 等于基准 | 大于基准] 三部分
// @return (int, int) 小于段的右端、大于段的左端, 中间 [lt, gt] 全等于 pivot
func threeWayPartition(nums []int, left int, right int, pivot int) (int, int) {
	lt, i, gt := left, left, right
	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	return lt - 1, gt + 1
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
// 归并排序
// ////////////////////////////////////////////////////

// MergeSort
// @Description: 归并排序 O(NlogN) 稳定
// @Solution: 分治拆小数组 + 合并有序数组
func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := n >> 1
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

// merge
// @Description: 合并两个有序数组
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

// MergeSortBottomUp
// @Description: 迭代版归并排序（自底向上, 无递归栈）
// @Solution: 按 1,2,4,8... 的步长两两合并
func MergeSortBottomUp(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	tmp := make([]int, n)
	for step := 1; step < n; step <<= 1 {
		for i := 0; i < n-step; i += step << 1 {
			left, mid, right := i, i+step-1, min(i+2*step-1, n-1)
			mergeRange(nums, tmp, left, mid, right)
		}
	}
	return nums
}

// mergeRange
// @Description: 合并 nums[left:mid+1] 和 nums[mid+1:right+1] (两子段各自有序)
func mergeRange(nums, tmp []int, left, mid, right int) {
	copy(tmp[left:right+1], nums[left:right+1])
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		if i > mid {
			nums[k] = tmp[j]
			j++
		} else if j > right {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
	for {
		largeIndex := pos
		if 2*pos+1 < n && nums[2*pos+1] > nums[largeIndex] {
			largeIndex = 2*pos + 1
		}
		if 2*pos+2 < n && nums[2*pos+2] > nums[largeIndex] {
			largeIndex = 2*pos + 2
		}
		if largeIndex == pos {
			break
		}
		nums[pos], nums[largeIndex] = nums[largeIndex], nums[pos]
		pos = largeIndex
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