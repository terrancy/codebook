package base

//
//  FindKTh
//  @Description: NC88 寻找第K大
//  @Description: 有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
//  @Description: 优先队列,维护利用插入排序维护一个排序好队列.队列大小为k
//  @param nums
//  @param n
//  @param k
//  @return int
//
func FindKTh(nums []int, n int, k int) int {
    if n == 0 || n < k {
        return 0
    }
    queue := make([]int, k)
    for i := 0; i < n; i++ {
        queue = priorityQueue(queue, nums[i], k)
    }
    return queue[0]
}

//
//  priorityQueue
//  @Description: 维护一个K长度的递增队列
//  @param queue
//  @param num
//  @param k
//  @return []int
//
func priorityQueue(queue []int, num int, k int) []int {
    n := len(queue)
    for idx := n - 1; idx > 0 && queue[idx] > num; {
        queue[idx+1] = queue[idx]
        idx--
    }
    if n+1 > k {
        queue = queue[n-k+1:]
    }
    return queue
}
