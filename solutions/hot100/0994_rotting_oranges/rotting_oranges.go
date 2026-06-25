package rottingoranges

import "fmt"

// OrangesRotting 腐烂的橘子
// 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数，如果不可能则返回 -1。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func OrangesRotting(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		visted[i] = make([]int, n)
	}

	var bound func(i, j int) bool
	bound = func(i, j int) bool {
		return i < 0 || i >= m || j < 0 || j >= n
	}

	dires := [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	cost := 0
	var bfs func(i, j int)
	bfs = func(i, j int) {
		if bound(i, j) {
			return
		}

		if grid[i][j] == 0 {

	return ret
}
