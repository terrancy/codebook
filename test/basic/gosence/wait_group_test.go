package gosence

import (
	"testing"
	"time"

	"terrancy/awesome/tutorial/gosence"
)

func TestWGTimeout(t *testing.T) {
	gosence.WGTimeout(1 * time.Second)
}

func TestWGPrinter(t *testing.T) {
	gosence.WGPrinter()
}
