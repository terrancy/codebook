package test

import (
    "awesome/graph"
    "fmt"
    "testing"
)

//
//  TestNumIslands
//  @Description: NC109 岛屿数量
//  @param t
//
func TestNumIslands(t *testing.T) {
    grid := [][]byte{
        {'1', '1', '0', '0', '0'},
        {'0', '1', '0', '1', '1'},
        {'0', '0', '0', '1', '1'},
        {'0', '0', '0', '0', '0'},
        {'0', '0', '1', '1', '1'},
    }
    cnt := graph.NumIslands(grid)
    fmt.Println(cnt)
}
