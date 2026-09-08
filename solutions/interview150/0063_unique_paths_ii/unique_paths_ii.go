package uniquepathsii

// UniquePathsWithObstacles 不同路径 II
// 机器人位于 m×n 网格左上角，每次只能向右或向下移动一步，要到达右下角。
// 网格中 obstacleGrid[i][j] 为 1 表示障碍物（不可通过），为 0 表示空位。
// 返回从左上角到右下角的不同路径总数；起点或终点是障碍物时返回 0。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 滚动数组：dp[j] 表示到达当前行第 j 列的路径数
	dp := make([]int, n)

	// 初始化第一行：遇到第一个障碍物后，其右侧所有格子都不可达
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		// 第一列：只有上一行的第一列可达且当前格不是障碍时保持 1，否则置 0
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				// 障碍物格子不可达，路径数置 0
				dp[j] = 0
			} else {
				// 状态转移：来自上方（dp[j] 的旧值）+ 来自左方（dp[j-1] 的新值）
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}
