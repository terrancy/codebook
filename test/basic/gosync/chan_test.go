package gosync

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/tutorial/gosync"
)

func TestChannel(t *testing.T) {
	gosync.Chan1()
	assert.Equal(t, 1, 1)
}

func TestChannel2(t *testing.T) {
	var ch chan int
	ch = make(chan int)
	ch <- 1
	close(ch)
}
