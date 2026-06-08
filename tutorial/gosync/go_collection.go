package gosync

import "fmt"

//
// Slice 相关
//

func GoSlice1() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1 = append(s1, 2)
	s = append(s, 10)
	fmt.Printf("s1: %v, len: %d, cap:%d \n", s1, len(s1), cap(s1))
}

// 实现并发安全的3种方法
// 1. sync.Map
// 2. map + sync.RWMutex
// 3. 分片锁
// 并发编程尽量减少锁的使用, 在不得不使用锁的时候,可减小锁的颗粒度和持有时间

func GoMap1() {
	hashData := make(map[int]int, 3)
	hashData[2] = 1
}

func GoMapIter1() {
	type student struct {
		Name string
		Age  int
	}
	m := make(map[string]*student)
	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range students {
		m[stu.Name] = &stu
	}
	fmt.Printf("m: %v\n", m)
}
