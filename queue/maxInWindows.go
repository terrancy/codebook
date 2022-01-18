package main

//
//  maxInWindows
//  @Description: NC82 滑动窗口的最大值 => 单调队列
//  @param num
//  @param size
//  @return []int
//
func maxInWindows(num []int, size int) []int {
    n := len(num)
    if size == 0 || n < size {
        return nil
    }
    deque := make([]int, 0)
    ans := make([]int, 0)
    for i := 0; i < n; i++ {
        for len(deque) > 0 && num[i] > num[deque[len(deque)-1]] {
            deque = queueLastPop(deque)
        }
        deque = queueLastPush(deque, i)
        if deque[0]+size <= i {
            deque = queueFirstPop(deque)
        }
        if i+1 >= size {
            ans = append(ans, num[deque[0]])
        }
    }
    return ans
}

func queueLastPop(queue []int) []int {
    return queue[:len(queue)-1]
}

func queueFirstPop(queue []int) []int {
    return append(queue[:0], queue[1:]...)
}

func queueLastPush(queue []int, num int) []int {
    return append(queue, num)
}
