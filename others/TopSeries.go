package others

import "strconv"

//
// FindKTh
// @Description: NC88 寻找第K大
// @Description: 有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
// @Description: 优先队列,维护利用插入排序维护一个排序好队列.队列大小为k
// @param nums
// @param n
// @param k
// @return int
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
// priorityQueue
// @Description: 维护一个K长度的递增队列
// @param queue
// @param num
// @param k
// @return []int
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

//
// TopKStrings
// @Description: NC97 字符串出现次数的TopK问题
// @param strings
// @param k
// @return [][]string
//
func TopKStrings(strings []string, k int) [][]string {
    // 统计次数
    tmp := make(map[string]int, 0)
    for _, val := range strings {
        tmp[val]++
    }
    // 字典 字符与次数
    data := make([]pair, len(tmp))
    for key, val := range tmp {
        data = append(data, pair{val: key, cnt: val})
    }
    n := len(data)
    // 堆排序
    for i := n/2 - 1; i >= 0; i-- {
        heapSort(data, n, i)
    }
    ans := make([][]string, 0)
    for i := n - 1; i >= 0; i-- {
        k--
        cnt := strconv.Itoa(data[0].cnt)
        ans = append(ans, []string{data[0].val, cnt})
        if k == 0 {
            break
        }
        data[0], data[i] = data[i], data[0]
        heapSort(data, i, 0)
    }
    return ans
}

type pair struct {
    val string
    cnt int
}

func heapSort(data []pair, n int, idx int) {
    maxIdx := idx
    leftIdx := idx*2 + 1
    rightIdx := idx*2 + 2
    if leftIdx < n && cmp(data[leftIdx], data[maxIdx]) {
        maxIdx = leftIdx
    }
    if rightIdx < n && cmp(data[rightIdx], data[maxIdx]) {
        maxIdx = rightIdx
    }
    if maxIdx != idx {
        // !!交换位置
        data[maxIdx], data[idx] = data[idx], data[maxIdx]
        heapSort(data, n, maxIdx)
    }
}

func cmp(a, b pair) bool {
    if a.cnt > b.cnt {
        return true
    } else if a.cnt == b.cnt && a.val < b.val {
        return true
    }
    return false
}
