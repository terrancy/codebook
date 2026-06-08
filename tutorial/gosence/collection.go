package gosence

import (
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

func (m RWMutexMap) Get(key string) (interface{}, bool) {
	m.RLocker()
	defer m.RUnlock()
	if val, ok := m.data[key]; ok {
		return val, true
	}
	return nil, false
}

func (m RWMutexMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}
func NewRwMutexMap() ISyncMap {
	return RWMutexMap{
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

// 方案3、分片锁

type ShardedMap struct{}

func (s ShardedMap) Get(key string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}

func (s ShardedMap) Set(key string, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func NewShardedMap() ISyncMap {
	return ShardedMap{}
}

// 场景题

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
	timer := time.NewTicker(maxWaitingDuration)
	defer timer.Stop()

	// 多路复用
	select {
	case <-timer.C:
		return nil, fmt.Errorf("超时了：%v", maxWaitingDuration)
	case <-ch:
		// 不操作，程序往下走
	}
	m.Lock()
	val := m.data[key]
	m.Unlock()
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
