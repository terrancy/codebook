package gosence

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 并发场景下，只并发读数据 => 没有问题；并发读+写或者并发写+写 => panic
// 并发安全的map, 3个方案
// 1、读多写极少时：sync.map
// 2、读多写少时：map+RWMutex 读写锁
// 3、读少写多时：shardedMap 分片锁

type ISyncMap interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// 方案1、并发读写锁

type RWMutexMap struct {
	sync.RWMutex
	data map[string]interface{}
}

func (m *RWMutexMap) Get(key string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	if val, ok := m.data[key]; ok {
		return val, true
	}
	return nil, false
}

func (m *RWMutexMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}
func NewRwMutexMap() ISyncMap {
	return &RWMutexMap{
		data: make(map[string]interface{}, 4),
	}
}

// 方案2、并发互斥锁
// 场景题：要求实现一个map：
// 1、面向高并发
// 2、只存在插入和查询操作 O（1）
// 3、查询时，若key存在，直接返回val；否则，阻塞直到key val对被放入后，获取val返回；
// 4、写入真实代码，不能有死锁或者 panic 风险。

type MutexMap struct {
	sync.Mutex
	data map[string]interface{}
}

func (m MutexMap) Get(key string) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()
	if val, ok := m.data[key]; ok {
		return val, true
	}
	return nil, false
}

func (m MutexMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func NewMutexMap() ISyncMap {
	return MutexMap{
		data: make(map[string]interface{}, 4),
	}
}

// 方案3、分片锁 https://github.com/orcaman/concurrent-map

type ShardedMap struct {
	shards []*shard
	cnt    int
}

type shard struct {
	sync.RWMutex
	data map[string]interface{}
}

// getShard 根据 key 定位分片
// 使用手写 FNV-1a 32位哈希，避免 fnv.New32a + []byte(key) 的逃逸分配，零 GC 压力
// FNV-1a 算法：hash = FNV_offset_basis，每字节 hash ^= byte，hash *= FNV_prime
func (s *ShardedMap) getShard(key string) *shard {
	h := uint32(2166136261) // FNV-1a 32位 offset basis
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])  // 异或当前字节
		h *= 16777619        // 乘 FNV-1a 32位 prime
	}
	return s.shards[h%uint32(s.cnt)]
}

func (s *ShardedMap) Get(key string) (interface{}, bool) {
	shard := s.getShard(key)
	shard.RLocker()
	defer shard.RUnlock()
	if val, ok := shard.data[key]; ok {
		return val, true
	}
	return nil, false
}

func (s *ShardedMap) Set(key string, value interface{}) {
	shard := s.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

func NewShardedMap() ISyncMap {
	shards := make([]*shard, 4)
	for i := range shards {
		shards[i] = &shard{
			data: make(map[string]interface{}),
		}
	}
	return &ShardedMap{
		shards: shards,
		cnt:    len(shards),
	}
}

// 场景题

// GoAsyncSlice
// title: 多协程查询切片问题
// desc: 假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排序。
// 限时5秒，使用多个goroutine查找切片中是否存在给定的值，在查找到目标值或者超时后立刻结束所有goroutine的执行多协程查询切片问题,
// 会导致切片长度和容量不一致
// link: https://interview.disign.me/#/question/q017
// params: data [23,32,78,43,76,65,345,762,......915,86]
// return:
func GoAsyncSlice(data []int, target int) (found bool, err error) {
	var (
		n            = len(data)
		chunkSize    = 100
		wg           sync.WaitGroup
		ch           = make(chan bool, 1)
		searchTarget func(i, j int)
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchTarget = func(i, j int) {
		defer wg.Done()
		for idx := range data[i:j] {
			// 监听上下文取消信号
			select {
			case <-ctx.Done():
				return
			default:
			}
			if data[idx] != target {
				continue
			}
			// 通道缓冲1，避免阻塞
			select {
			case ch <- true:
				return
			default:
			}
			return
		}
	}

	for i := 0; i < n; i += chunkSize {
		j := i + chunkSize
		if j > n {
			j = n
		}
		wg.Add(1)
		go searchTarget(i, j)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case found = <-ch:
		return found, nil
	}

	return
}

// GoAsyncMapVipVisited 高并发vip限流
// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
// link：https://interview.disign.me/#/question/q011
func GoAsyncMapVipVisited() {

}

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{
		visitIPs: make(map[string]time.Time, 4),
	}
}

func (b *Ban) visit(ip string) bool {
	if _, ok := b.visitIPs[ip]; ok {
		return true
	}
	b.visitIPs[ip] = time.Now()
	return false
}

// 场景题：要求实现一个map：
// 1、面向高并发
// 2、只存在插入和查询操作 O（1）
// 3、查询时，若key存在，直接返回val；否则，阻塞直到key val对被放入后，获取val返回；
// 4、写入真实代码，不能有死锁或者 panic 风险。

type ConcurrentMutexMap struct {
	sync.Mutex
	keyToCh map[string]chan struct{}
	data    map[string]interface{}
}

func (m *ConcurrentMutexMap) Get(key string, maxWaitingDuration time.Duration) (interface{}, error) {
	m.Lock()
	if val, ok := m.data[key]; ok {
		m.Unlock()
		return val, nil
	}
	ch, ok := m.keyToCh[key]
	if !ok {
		ch = make(chan struct{})
		m.keyToCh[key] = ch
	}
	m.Unlock()

	// 大量使用会导致大量僵尸定时器占用内存
	timer := time.NewTimer(maxWaitingDuration)
	defer timer.Stop()

	// 多路复用
	select {
	case <-timer.C:
		// 清理超时的channel(需要再次加锁保护)
		m.Lock()
		if ch == m.keyToCh[key] {
			delete(m.keyToCh, key)
		}
		m.Unlock()
		return nil, fmt.Errorf("超时了：%v", maxWaitingDuration)
	case <-ch:
		// 不操作，程序往下走
	}
	m.Lock()
	val, ok := m.data[key]
	m.Unlock()
	if !ok {
		return nil, fmt.Errorf("key %s 已删除", key)
	}
	return val, nil
}

func (m *ConcurrentMutexMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
	if ch, ok := m.keyToCh[key]; ok {
		close(ch)
		delete(m.keyToCh, key)
	}
}

func NewConcurrentMutexMap() *ConcurrentMutexMap {
	return &ConcurrentMutexMap{
		data:    make(map[string]interface{}, 4),
		keyToCh: make(map[string]chan struct{}, 4),
	}
}