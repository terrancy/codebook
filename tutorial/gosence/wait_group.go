package gosence

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wait group add-done-wait
// 为 sync.WaitGroup中Wait函数支持 WaitTimeout 功能.
// docs: https://interview.disign.me/#/question/q013

func WGTimeout(timeout time.Duration) {
	wg := sync.WaitGroup{}
	ch1 := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, ch <-chan struct{}) {
			defer wg.Done()
			<-ch
			fmt.Println(num)
		}(i, ch1)
	}

	if waitTimeout(&wg, timeout) {
		fmt.Println("timeout exit")
	} else {
		fmt.Println("all done")
	}
	close(ch1)
	time.Sleep(time.Second * 1)
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch2 := make(chan struct{})

	go func() {
		wg.Wait()
		fmt.Println("wg wait done..")
		close(ch2)
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true
	case <-ch2:
		return false
	}
}

// WGChannel
// 在 golang 协程和channel配合使用
// 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。
// 最终输出五个随机数。
// https://interview.disign.me/#/question/q009

func WGChannel() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
	}()
	wg.Wait()
	close(ch)
}

// WGPrinter 交替打印数字和字母
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
// https://interview.disign.me/#/question/q001
func WGPrinter() {
	// 使用 wait_group + channel
	var (
		wg         sync.WaitGroup
		number     = make(chan struct{})
		letter     = make(chan struct{})
		numPrint   func()
		alphaPrint func()
	)
	wg.Add(1)

	numPrint = func() {
		i := 0
		for {
			select {
			case _, closed := <-number:
				if !closed {
					return
				}
				i++
				fmt.Print(i)
				i++
				fmt.Print(i)
				letter <- struct{}{}
			}
		}
	}

	alphaPrint = func() {
		defer wg.Done()
		j := 'A'
		for {
			select {
			case <-letter:
				if j > 'Z' {
					return
				}
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				number <- struct{}{}
			}
		}
	}

	go numPrint()
	go alphaPrint()
	number <- struct{}{}
	wg.Wait()
	close(number)
	close(letter)
}
