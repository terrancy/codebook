package base

import "container/heap"

/*堆 优先队列*/
type Item struct {
    Value    int // The Value of the item; arbitrary.
    Priority int // The Priority of the item in the queue.
    Index    int // The Index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].Index = i
    pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.Index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.Index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

func (pq *PriorityQueue) update(item *Item, Value int, Priority int) {
    item.Value = Value
    item.Priority = Priority
    heap.Fix(pq, item.Index)
}
