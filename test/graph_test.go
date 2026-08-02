package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terrancy/awesome/graph"
)

var evaluateDivisionCases = []struct {
	name      string
	equations [][]string
	values    []float64
	queries   [][]string
	expected  []float64
}{
	{
		name:      "test_case_1",
		equations: [][]string{{"a", "b"}, {"b", "c"}},
		values:    []float64{2.0, 3.0},
		queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		expected:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
	},
	{
		name:      "test_case_2",
		equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		values:    []float64{1.5, 2.5, 5.0},
		queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
		expected:  []float64{3.75, 0.4, 5.0, 0.2},
	},
	{
		name:      "test_case_3",
		equations: [][]string{{"a", "b"}},
		values:    []float64{0.5},
		queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
		expected:  []float64{0.5, 2.0, -1.0, -1.0},
	},
}

// TestEvaluateDivision LC399.除法求值
func TestEvaluateDivision(t *testing.T) {
	for _, tt := range evaluateDivisionCases {
		t.Run(tt.name, func(t *testing.T) {
			res := graph.EvaluateDivision(tt.equations, tt.values, tt.queries)
			assert.Equal(t, len(tt.expected), len(res))
			for i := range res {
				if tt.expected[i] == -1.0 {
					assert.Equal(t, tt.expected[i], res[i])
				} else {
					assert.InDelta(t, tt.expected[i], res[i], 1e-6)
				}
			}
		})
	}
}

var numIslandsCases = []struct {
	name     string
	grid     [][]byte
	expected int
}{
	{
		name: "test_case_1",
		grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		expected: 1,
	},
	{
		name: "test_case_2",
		grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		expected: 3,
	},
	{
		name: "test_case_3",
		grid: [][]byte{
			{'1', '0', '1', '0', '1'},
			{'0', '1', '0', '1', '0'},
			{'1', '0', '1', '0', '1'},
		},
		expected: 8,
	},
	{
		name:     "test_case_4",
		grid:     [][]byte{{'0', '0', '0', '0', '0'}},
		expected: 0,
	},
}

// TestNumIslands LC200.岛屿数量
func TestNumIslands(t *testing.T) {
	for _, tt := range numIslandsCases {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]byte, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			res := graph.NumIslands(grid)
			assert.Equal(t, tt.expected, res)
		})
	}
}