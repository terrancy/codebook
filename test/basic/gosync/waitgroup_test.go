package gosync

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/tutorial/gosync"
)

func TestWG1(t *testing.T) {
	//expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//rsp := gosync.WG1(10)
	//assert.ElementsMatch(t, expect, rsp)

	cases := []struct {
		name     string
		taskNum  int
		expected []interface{}
	}{
		{
			name:     "wg1-task10",
			taskNum:  10,
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "wg1-task5",
			taskNum:  5,
			expected: []interface{}{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel() // 并行执行测试
			result := gosync.WG1(tt.taskNum)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
