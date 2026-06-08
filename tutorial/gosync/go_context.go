package gosync

import (
	"context"
	"time"
)

func context1() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}

func context2() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
}
