package others

func LRU(operators [][]int, k int) []int {
    var res []int
    lruCache := initLru(k)
    for _, operator := range operators {
        o := operator[0]
        switch o {
        case 1:
            lruCache.set(operator[1], operator[2])
        case 2:
            res = append(res, lruCache.get(operator[1]))
        }
    }
    return res
}

type LruCache struct {
    Size       int
    Capacity   int
    Cache      map[int]*NodeList
    Head, Tail *NodeList
}

type NodeList struct {
    Key, Value int
    Pre, Next  *NodeList
}

func (l *LruCache) get(key int) int {
    if node, ok := l.Cache[key]; ok {
        l.removeToHead(node)
        return node.Value
    }
    return -1
}

func (l *LruCache) set(key, value int) {
    if node, ok := l.Cache[key]; ok {
        node.Value = value
        l.removeToHead(node)
        return
    }
    
    node := initNode(key, value)
    l.addToHead(node)
    l.Cache[key] = node
    if l.Size == l.Capacity {
        n := l.removeTail()
        delete(l.Cache, n.Key)
        return
    }
    l.Size++
}

func initLru(cap int) *LruCache {
    lruCache := &LruCache{
        Size:     0,
        Capacity: cap,
        Cache:    make(map[int]*NodeList, 0),
        Head:     initNode(0, 0),
        Tail:     initNode(0, 0),
    }
    lruCache.Head.Next = lruCache.Tail
    lruCache.Tail.Pre = lruCache.Head
    return lruCache
}

func (l *LruCache) addToHead(node *NodeList) {
    node.Next = l.Head.Next
    node.Pre = l.Head
    l.Head.Next.Pre = node
    l.Head.Next = node
}

func (l *LruCache) removeToHead(node *NodeList) {
    l.removeNode(node)
    l.addToHead(node)
}

func (l *LruCache) removeNode(node *NodeList) {
    node.Next.Pre = node.Pre
    node.Pre.Next = node.Next
}

func (l *LruCache) removeTail() *NodeList {
    node := l.Tail.Pre
    l.removeNode(node)
    return node
}

func initNode(key, value int) *NodeList {
    return &NodeList{
        Key:   key,
        Value: value,
    }
}
