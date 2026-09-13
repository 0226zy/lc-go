package maximalsquare

// MaximalSquare 最大正方形
// 在一个由 '0' 和 '1' 组成的二维矩阵 matrix 中，找到只包含 '1' 的最大正方形，并返回其面积。
// dp[i][j] 表示以 matrix[i-1][j-1] 为右下角的最大正方形边长：
// 若 matrix[i-1][j-1] == '1'，则 dp[i][j] = min(左, 上, 左上) + 1。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	// dp 多开一行一列，下标偏移后可避免边界特判
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

// MaximalSquareOptimized 最大正方形（滚动一维数组空间优化版）
// 只用一维 dp，更新时用 prev 保存左上角旧值。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MaximalSquareOptimized(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	maxSide := 0

	for i := 1; i <= m; i++ {
		prev := 0 // 上一行的 dp[j-1]，即左上角
		for j := 1; j <= n; j++ {
			temp := dp[j] // 更新前的 dp[j] 是上一行同列，稍后成为下一列的左上
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(dp[j], dp[j-1], prev) + 1
				if dp[j] > maxSide {
					maxSide = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return maxSide * maxSide
}
