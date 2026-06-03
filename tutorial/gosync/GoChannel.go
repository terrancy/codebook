package gosync

import (
	"fmt"
	"time"
)

// go channel 的常见使用场景

func Chan1() {
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	<-timer.C
	fmt.Println("5秒后执行")
}

// go channel 的常见问题

// go wait group 的使用
