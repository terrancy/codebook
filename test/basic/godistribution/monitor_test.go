package godistribution

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestGoLeak1(t *testing.T) {
	defer goleak.VerifyNone(t)
	// 测试代码
	go func() {
		time.Sleep(time.Second)
	}()
	fmt.Println(11)
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
