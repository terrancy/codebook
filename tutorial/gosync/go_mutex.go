package gosync

import "sync"

// 单机锁相关知识

func mutex1() {
	m := sync.Mutex{}
	m.Lock()

	rw := sync.RWMutex{}
	rw.Lock()
}
