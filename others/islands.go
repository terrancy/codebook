package others

// 岛屿问题

var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var n, m int

//
//  NumIslands
//  @Description: NC109 岛屿数量
//  @param grid
//  @return int
//
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

//
//  dfsNumIslands
//  @Description: 将节点赋值为'0',上下左右都赋值为'0'
//  @param grid
//  @param i
//  @param j
//
func dfsNumIslands(grid [][]byte, i int, j int) {
    if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    for _, direction := range directions {
        dfsNumIslands(grid, i+direction[0], j+direction[1])
    }
}
