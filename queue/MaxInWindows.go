package queue

//
// maxInWindows
// @Description: NC82 滑动窗口的最大值 => 单调队列
// @Solution: 1. 维护单调队列,新元素较大时剔除较小元素再添加 = 类似于插入排序
// @Solution: 2. 判断窗口大小, 队首元素出队
// @Solution: 3. 判断是否记录,满足队首元素
// @param num
// @param size
// @return []int
//
func MaxInWindows(num []int, size int) []int {
    n := len(num)
    if size == 0 || n < size {
        return nil
    }
    deque := make([]int, 0)
    ans := make([]int, 0)
    for i := 0; i < n; i++ {
        // 维护单调队列,新元素比较大时剔除较小的元素
        for len(deque) > 0 && num[i] > num[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        // 滑动窗口大小
        if i-deque[0] >= size {
            deque = deque[1:]
        }
        // 记录队首元素
        if i+1 >= size {
            ans = append(ans, num[deque[0]])
        }
    }
    return ans
}
