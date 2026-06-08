package gosync

import (
	"sync"
)

// 并发等待组 WaitGroup

func WG1(taskNum int) []interface{} {
	var (
		dataCh = make(chan interface{})
		resp   = make([]interface{}, 0, taskNum)
	)

	go func() {
		var wg = sync.WaitGroup{}
		for i := 0; i < taskNum; i++ {
			wg.Add(1)
			go func(ch chan<- interface{}, val int) {
				defer wg.Done()
				ch <- val
			}(dataCh, i)
		}
		wg.Wait()
		close(dataCh)
	}()

	for val := range dataCh {
		resp = append(resp, val)
	}

	return resp
}
