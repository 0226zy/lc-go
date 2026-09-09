package uniquepathsii

// UniquePathsWithObstacles 不同路径 II（标准 DP 数组版）
// 机器人位于 m×n 网格左上角，每次只能向右或向下走，网格中 1 表示障碍物。
// dp[i][j] 表示到达 (i,j) 的路径数；障碍格为 0，否则 dp[i][j] = dp[i-1][j] + dp[i][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0 // 起点或终点是障碍，直接无解
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1 // base case：起点一种走法
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue // 障碍格保持 0，不可达
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j] // 上方来路
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1] // 左方来路
			}
		}
	}
	return dp[m-1][n-1]
}

// UniquePathsWithObstaclesOptimized 不同路径 II（滚动一维数组空间优化版）
// dp[j] 就地更新：未更新时是上一行的值（上方来路），dp[j-1] 是本行新值（左方来路）。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func UniquePathsWithObstaclesOptimized(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0 // 起点或终点是障碍，直接无解
	}
	dp := make([]int, n)
	dp[0] = 1 // base case：起点一种走法
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0 // 障碍格：该列累计路径清零
			} else if j > 0 {
				dp[j] += dp[j-1] // 左方来路累加进上方来路
			}
		}
	}
	return dp[n-1]
}
