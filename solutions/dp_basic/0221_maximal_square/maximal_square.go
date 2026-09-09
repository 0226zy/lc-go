package maximalsquare

// MaximalSquare 最大正方形（标准 DP 数组版）
// 在 0/1 矩阵中找全 1 的最大正方形，返回面积。
// dp[i][j] 表示以 matrix[i-1][j-1] 为右下角的全 1 正方形最大边长，
// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1（当前格为 '1' 时）。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func MaximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1) // 第 0 行、第 0 列为哨兵，默认 0
	}
	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				// 木桶效应：边长受制于左、上、左上三个方向中最短的一块
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
			// 当前格为 '0' 时 dp[i][j] 保持 0
		}
	}
	return maxSide * maxSide
}

// MaximalSquareOptimized 最大正方形（一维滚动数组空间优化版）
// 转移只依赖上一行和当前行左侧，用一维数组滚动覆盖，prev 保存左上角的旧值。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MaximalSquareOptimized(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	maxSide := 0
	for i := 1; i <= m; i++ {
		prev := 0 // 保存 dp[i-1][j-1]（左上角）
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 覆盖前先记住上一行的 dp[i-1][j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(dp[j], dp[j-1], prev) + 1
				maxSide = max(maxSide, dp[j])
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return maxSide * maxSide
}
