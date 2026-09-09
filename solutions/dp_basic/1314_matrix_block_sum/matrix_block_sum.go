package matrixblocksum

// MatrixBlockSum 矩阵区域和（二维前缀和）
// dp[i][j] 表示 mat 前 i 行前 j 列的元素和，
// dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]。
// 每个 answer[i][j] 通过容斥原理 O(1) 查询：大矩形 - 上 - 左 + 左上。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func MatrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	// dp[i][j] = mat 前 i 行前 j 列的元素和，多开一行一列简化边界
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 当前格 + 上方矩形 + 左方矩形 - 重复计算的左上矩形
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	answer := make([][]int, m)
	for i := range answer {
		answer[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// 把以 (i,j) 为中心、半径 k 的区域边界裁剪到矩阵内
			r1, r2 := max(0, i-k), min(m-1, i+k)
			c1, c2 := max(0, j-k), min(n-1, j+k)
			// 容斥：大矩形 - 上 - 左 + 左上
			answer[i][j] = dp[r2+1][c2+1] - dp[r1][c2+1] - dp[r2+1][c1] + dp[r1][c1]
		}
	}
	return answer
}
