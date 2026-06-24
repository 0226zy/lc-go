package numberofislands

// NumIslands 岛屿数量 (DFS)
// 计算二维网格中岛屿的数量。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func NumIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' || visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// NumIslandsBFS 岛屿数量 (BFS)
// 计算二维网格中岛屿的数量，使用广度优先搜索。
// 时间复杂度: O(m*n)  空间复杂度: O(min(m,n)) —— BFS 队列最大长度为短边
func NumIslandsBFS(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	type pos struct{ i, j int }
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			count++
			grid[i][j] = '0' // 入队前标记，避免重复入队
			queue := []pos{{i, j}}

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]

				for _, d := range dirs {
					ni, nj := curr.i+d[0], curr.j+d[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= n {
						continue
					}
					if grid[ni][nj] == '0' {
						continue
					}
					grid[ni][nj] = '0'
					queue = append(queue, pos{ni, nj})
				}
			}
		}
	}

	return count
}
