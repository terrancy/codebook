package gosync

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/tutorial/gosync"
)

func TestGoSlice1(t *testing.T) {
	gosync.GoSlice1()
}

func TestGoMap1(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestGoMapIter1(t *testing.T) {
	gosync.GoMapIter1()
}
