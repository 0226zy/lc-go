package minimumpathsum

// MinPathSum 最小路径和
// 给定一个 m×n 的非负整数网格 grid，每次只能向下或向右移动一步，
// 找出从左上角到右下角路径数字总和最小的路径，返回该最小总和。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MinPathSum(grid [][]int) int {
	// 防御性检查：空网格按题意不会出现，但避免越界 panic
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	// dp[j] 表示从起点走到当前行第 j 列的最小路径和（一维滚动数组，不修改输入）
	dp := make([]int, n)
	dp[0] = grid[0][0]
	// 第一行只能从左边走过来，依次累加
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		// 第一列只能从上边走下来，直接累加
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			// dp[j] 旧值是上方结果，dp[j-1] 新值是左方结果
			dp[j] = grid[i][j] + min(dp[j], dp[j-1])
		}
	}
	return dp[n-1]
}
