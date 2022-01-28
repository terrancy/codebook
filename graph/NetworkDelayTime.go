package graph

import (
    "awesome/base"
    "container/heap"
    "math"
)

//
//  NetworkDelayTime
//  @Description: 网络延迟问题
//  @param times
//  @param N
//  @param K
//  @return int
//
func NetworkDelayTime(times [][]int, N int, K int) int {
    max := 0
    mGraph := make([][]int, N)
    /*K到各点的延迟时间，dist使用优先队列存储，存在最小堆中*/
    dist := make([]int, N)
    /*初始化优先队列，初始化开始位置dist=0, 其他位置为正无穷*/
    pq := &base.PriorityQueue{}
    for i := 0; i < N; i++ {
        mGraph[i] = make([]int, N)
        for j := 0; j < N; j++ {
            mGraph[i][j] = -1
        }
        dist[i] = math.MaxInt64
    }
    /*初始推的第一个元素*/
    dist[K-1] = 0
    item := &base.Item{
        Value:    K - 1,
        Priority: 0,
    }
    heap.Push(pq, item)
    for _, time := range times {
        mGraph[time[0]-1][time[1]-1] = time[2]
    }
    /*已收集的最短路径的顶点*/
    collected := make(map[int]bool)
    /*每次取堆中最小元素，按照递增的顺序找出各个顶点的最短路*/
    for pq.Len() > 0 {
        item := heap.Pop(pq).(*base.Item)
        v := item.Value
        if collected[v] == true {
            continue
        }
        collected[v] = true
        for w := range mGraph[v] {
            if !collected[w] && mGraph[v][w] >= 0 {
                if dist[v]+mGraph[v][w] < dist[w] {
                    dist[w] = dist[v] + mGraph[v][w]
                    item := &base.Item{
                        Value:    w,
                        Priority: dist[w],
                    }
                    heap.Push(pq, item)
                }
            }
        }
    }
    for _, v := range dist {
        if v == math.MaxInt64 {
            return -1
        }
        if v > max {
            max = v
        }
    }
    return max
}
