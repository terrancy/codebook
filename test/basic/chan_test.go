package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"awesome/tutorial/gosync"
)

func TestChannel(t *testing.T) {
	gosync.Chan1()
	assert.Equal(t, 1, 1)
}
