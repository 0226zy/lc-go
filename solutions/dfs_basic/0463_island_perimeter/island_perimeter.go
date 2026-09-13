package islandperimeter

// IslandPerimeter 岛屿的周长
// 给定一个二维网格地图，1 表示陆地、0 表示水域，网格中恰好有一个岛屿，计算该岛屿的周长。
// 时间复杂度: O(m*n) 每个格子最多被染色一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func IslandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	perimeter := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或踩到水：这一步对应岛屿的一条边
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			perimeter++
			return
		}
		// 已访问过的陆地：直接返回，避免重复计数与死循环
		if grid[i][j] == -1 {
			return
		}
		// 染色标记当前陆地
		grid[i][j] = -1
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j) // 题目保证只有一座岛，遍历完即可返回
				return perimeter
			}
		}
	}
	return 0
}

// IslandPerimeterCount 岛屿的周长（计数解法）
// 每块陆地贡献 4 条边，每对相邻陆地共享 2 条边；只检查上方和左方的邻居避免重复扣除。
// 时间复杂度: O(m*n)  空间复杂度: O(1)，不修改输入网格
func IslandPerimeterCount(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			perimeter += 4
			if i > 0 && grid[i-1][j] == 1 {
				perimeter -= 2
			}
			if j > 0 && grid[i][j-1] == 1 {
				perimeter -= 2
			}
		}
	}
	return perimeter
}
