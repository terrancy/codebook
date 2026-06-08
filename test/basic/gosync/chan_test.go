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
