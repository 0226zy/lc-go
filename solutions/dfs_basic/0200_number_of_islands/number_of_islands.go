package numberofislands

// NumIslands 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的二维网格，统计其中岛屿的数量；
// 岛屿由水平或竖直方向相邻的陆地连接而成，四周均被水包围。
// 时间复杂度: O(m*n) 每个格子最多被访问一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	count := 0

	// directions 上下左右四个方向的行、列偏移量
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或遇到水，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		// 淹掉当前陆地，充当“已访问”标记，避免重复计数
		grid[i][j] = '0'
		// 向四个方向继续扩散
		for _, d := range directions {
			dfs(i+d[0], j+d[1])
		}
	}

	// 扫描每个格子，遇到未被淹掉的陆地就是一座新岛
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j) // 把整座岛屿一次性淹掉
			}
		}
	}
	return count
}
