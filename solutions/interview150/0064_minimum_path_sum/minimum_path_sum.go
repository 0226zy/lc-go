package minimumpathsum

// MinPathSum 最小路径和（标准 DP 数组版）
// 非负整数网格中每次只能向右或向下走，求从左上角到右下角路径数字和的最小值。
// dp[i][j] 表示走到 (i,j) 的最小路径和，dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func MinPathSum(grid [][]int) int {
	// 防御性检查：空网格按题意不会出现，但避免越界 panic
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = grid[i][j] // base case：起点
			case i == 0:
				dp[i][j] = dp[i][j-1] + grid[i][j] // 第一行只能从左边来
			case j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j] // 第一列只能从上边来
			default:
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1]) // 选较小的来路
			}
		}
	}
	return dp[m-1][n-1]
}

// MinPathSumOptimized 最小路径和（滚动一维数组空间优化版）
// dp[j] 就地更新：更新前是上方来路，dp[j-1] 已更新为左方来路，不修改输入网格。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MinPathSumOptimized(grid [][]int) int {
	// 防御性检查：空网格按题意不会出现，但避免越界 panic
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				dp[j] = grid[i][j] // base case：起点
			case i == 0:
				dp[j] = dp[j-1] + grid[i][j] // 第一行只能从左边来
			case j == 0:
				dp[j] += grid[i][j] // 第一列：dp[j] 还是上一行的值，直接累加
			default:
				dp[j] = grid[i][j] + min(dp[j], dp[j-1]) // dp[j]=上方来路，dp[j-1]=左方来路
			}
		}
	}
	return dp[n-1]
}
