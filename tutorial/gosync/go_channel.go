package gosync

import (
	"fmt"
	"time"
)

// go channel 的常见使用场景

const (
	ch1 = 1 << iota
	ch2
	ch3
	ch4
)

func Chan1() {
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	<-timer.C
	fmt.Println("5秒后执行")

	fmt.Println(ch4)
}

// go channel 的常见问题

// go wait group 的使用
