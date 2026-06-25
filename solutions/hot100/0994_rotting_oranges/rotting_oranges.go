package rottingoranges

// OrangesRotting 腐烂的橘子 (多源 DFS)
// 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数，如果不可能则返回 -1。
// 思路：对每个腐烂源分别 DFS，time[i][j] 记录格子被腐烂的最早时间（取 min），
//
//	最终答案为所有新鲜橘子 time 的最大值。
//
// 时间复杂度: O(m*n*源数)，剪枝后接近 O(m*n)  空间复杂度: O(m*n)
func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// time[i][j] 记录 (i,j) 被腐烂的最早时间，初始为 INF
	const INF = 1 << 30
	time := make([][]int, m)
	for i := range time {
		time[i] = make([]int, n)
		for j := range time[i] {
			time[i][j] = INF
		}
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// dfs 从一个腐烂源出发，更新每个可达格子的最短腐化时间
	var dfs func(i, j, dist int)
	dfs = func(i, j, dist int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 { // 空格子，不可通过
			return
		}
		if dist >= time[i][j] { // 剪枝：当前路径不更优
			return
		}
		time[i][j] = dist
		for _, d := range dirs {
			dfs(i+d[0], j+d[1], dist+1)
		}
	}

	// 统计新鲜橘子数，收集所有腐烂源
	fresh := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				dfs(i, j, 0) // 每个腐烂源分别 DFS
			}
		}
	}

	// 没有新鲜橘子，直接返回 0
	if fresh == 0 {
		return 0
	}

	// 取所有新鲜橘子 time 的最大值
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if time[i][j] == INF {
					return -1 // 有新鲜橘子永远无法被腐烂
				}
				if time[i][j] > ans {
					ans = time[i][j]
				}
			}
		}
	}
	return ans
}

// OrangesRottingBFS 腐烂的橘子 (多源 BFS)
// 返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数，如果不可能则返回 -1。
// 思路：所有腐烂橘子同时入队，分层 BFS 扩散，层数即为分钟数。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func OrangesRottingBFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh, minutes := 0, 0
	type pos struct{ i, j int }
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := []pos{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				// 多个腐烂源同时入队，做为第一分钟到源，即第一层
				queue = append(queue, pos{i, j})
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	// 层序，先处理第一层
	for len(queue) > 0 && fresh > 0 {
		size := len(queue) // 当前层的个数
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				ni, nj := curr.i+dir[0], curr.j+dir[1]
				if ni < 0 || ni >= m || nj < 0 || nj >= n {
					continue
				}
				if grid[ni][nj] == 1 {
					fresh--
					grid[ni][nj] = 2
					queue = append(queue, pos{ni, nj})
				}

			}

		}			// 一层结束，也是一分钟结束
			minutes++

	// 还有新鲜的
	if fresh > 0 {
		return -1
	}
	return minutes

}
