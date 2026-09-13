package maxareaofisland

// MaxAreaOfIsland 岛屿的最大面积
// 给你一个由 1（陆地）和 0（水）组成的二维网格，计算岛屿的最大面积。岛屿面积是值为 1 且四连通的格子数。
// 时间复杂度: O(m*n) 每个格子最多被访问一次  空间复杂度: O(m*n) 递归栈深度最坏情况
func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	maxArea := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 越界或遇到水（含已被淹没的陆地），贡献为 0
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		// 淹掉当前陆地，避免重复访问，并累加四个方向的面积
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if area := dfs(i, j); area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
