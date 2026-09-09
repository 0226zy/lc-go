package triangle

// MinimumTotal 三角形最小路径和（标准 DP 数组版，自底向上）
// dp[i][j] 表示从 (i, j) 出发走到底边的最小路径和，
// dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])，答案为 dp[0][0]。
// 时间复杂度: O(n²)  空间复杂度: O(n²)（n 为三角形行数）
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	copy(dp[n-1], triangle[n-1]) // base case：底边的路径和就是自身
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// 下一步只能走到下一行的 j 或 j+1，选更小的
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

// MinimumTotalOptimized 三角形最小路径和（一维滚动数组优化版）
// dp[j] 滚动表示"当前行第 j 列走到底边的最小路径和"，每行从左到右原地更新。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func MinimumTotalOptimized(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	copy(dp, triangle[n-1]) // base case：底边
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// dp[j]、dp[j+1] 此时仍是下一行的旧值，可原地覆盖
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}
