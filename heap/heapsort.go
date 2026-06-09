package heap

// heapify 调整以 i 为根的子树为最大堆
func heapify(arr []int, n, i int) {
	priority := i    // 初始化最大值为根
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点比根大
	if left < n && less(arr, priority, left) {
		priority = left
	}

	// 如果右子节点比最大值大
	if right < n && less(arr, priority, right) {
		priority = right
	}

	// 如果最大值不是根
	if priority != i {
		arr[i], arr[priority] = arr[priority], arr[i]
		// 递归调整受影响的子树
		heapify(arr, n, priority)
	}
}

// less 最大堆 或者 最小堆
func less(arr []int, i, j int) bool {
	return arr[i] < arr[j]
}

// Push 向最大堆中添加元素
func Push(arr *[]int, val int) {
	*arr = append(*arr, val)
	n := len(*arr)
	// 从最后一个元素开始向上调整
	for i := n - 1; i > 0; {
		parent := (i - 1) / 2
		if (*arr)[parent] >= (*arr)[i] {
			break
		}
		(*arr)[parent], (*arr)[i] = (*arr)[i], (*arr)[parent]
		i = parent
	}
}

// Pop 从最大堆中移除并返回堆顶元素（最大值）
func Pop(arr *[]int) (int, bool) {
	n := len(*arr)
	if n == 0 {
		return 0, false
	}

	mx := (*arr)[0]
	// 将最后一个元素移到堆顶
	(*arr)[0] = (*arr)[n-1]
	*arr = (*arr)[:n-1]

	// 向下调整堆
	if len(*arr) > 0 {
		heapify(*arr, len(*arr), 0)
	}

	return mx, true
}

func Heapsort(arr []int) {
	n := len(arr)
	// 建立最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 从堆顶开始，每次将堆顶元素与堆尾元素交换，调整堆，重复 n-1 次
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
