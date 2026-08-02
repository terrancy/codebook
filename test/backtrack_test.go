package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/backtrack"
)

var findWordsCases = []struct {
	name     string
	board    [][]byte
	words    []string
	expected []string
}{
	{
		name: "test_case_1",
		board: [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		},
		words:    []string{"oath", "pea", "eat", "rain"},
		expected: []string{"eat", "oath"},
	},
	{
		name:     "test_case_2",
		board:    [][]byte{{'a', 'b'}, {'c', 'd'}},
		words:    []string{"abcb"},
		expected: []string{},
	},
	{
		name:     "test_case_3",
		board:    [][]byte{{'a'}},
		words:    []string{"a"},
		expected: []string{"a"},
	},
}

// TestFindWords LC212.单词搜索 II
func TestFindWords(t *testing.T) {
	for _, tt := range findWordsCases {
		t.Run(tt.name, func(t *testing.T) {
			res := backtrack.FindWords(tt.board, tt.words)
			sort.Strings(res)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var movingCountCases = []struct {
	name     string
	m        int
	n        int
	k        int
	expected int
}{
	{
		name:     "test_case_1",
		m:        2,
		n:        3,
		k:        1,
		expected: 3,
	},
	{
		name:     "test_case_2",
		m:        3,
		n:        1,
		k:        0,
		expected: 1,
	},
	{
		name:     "test_case_3",
		m:        1,
		n:        1,
		k:        0,
		expected: 1,
	},
}

// TestMovingCount JZ13.机器人的运动范围
func TestMovingCount(t *testing.T) {
	for _, tt := range movingCountCases {
		t.Run(tt.name, func(t *testing.T) {
			res := backtrack.MovingCount(tt.m, tt.n, tt.k)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var hasMatrixPathCases = []struct {
	name     string
	board    [][]byte
	word     string
	expected bool
}{
	{
		name:     "test_case_1",
		board:    [][]byte{{'a', 'b'}, {'c', 'd'}},
		word:     "abcd",
		expected: false,
	},
	{
		name: "test_case_2",
		board: [][]byte{
			{'a', 'b', 'c', 'e'},
			{'s', 'f', 'c', 's'},
			{'a', 'd', 'e', 'e'},
		},
		word:     "abcced",
		expected: true,
	},
	{
		name:     "test_case_3",
		board:    [][]byte{{'a'}},
		word:     "a",
		expected: true,
	},
}

// TestHasMatrixPath JZ12.矩阵中的路径
func TestHasMatrixPath(t *testing.T) {
	for _, tt := range hasMatrixPathCases {
		t.Run(tt.name, func(t *testing.T) {
			res := backtrack.HasMatrixPath(tt.board, tt.word)
			assert.Equal(t, tt.expected, res)
		})
	}
}