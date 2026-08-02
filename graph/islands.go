package graph

// 岛屿问题

var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var n, m int

// NumIslands
// @Title: LC200.岛屿数量
// @Description: 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// @Description: 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// @Description: 此外，你可以假设该网格的四条边均被水包围。
// @Link: https://leetcode.cn/problems/number-of-islands/
// @param grid
// @return int
func NumIslands(grid [][]byte) int {
	n = len(grid)
	m = len(grid[0])
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfsNumIslands(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

// dfsNumIslands
// @Description: 将节点赋值为'0',上下左右都赋值为'0'
// @param grid
// @param i
// @param j
func dfsNumIslands(grid [][]byte, i int, j int) {
	if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for _, direction := range directions {
		dfsNumIslands(grid, i+direction[0], j+direction[1])
	}
}