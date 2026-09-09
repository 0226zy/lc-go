package minimumfallingpathsum

// MinFallingPathSum 下降路径最小和（标准 DP 数组版）
// 从第一行任意位置出发，每步走到下一行的正下/左下/右下，求到最后一行的最小路径和。
// dp[i][j] 表示以 (i,j) 为终点的下降路径最小和：
// dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])（越界忽略）。
// 时间复杂度: O(n²)  空间复杂度: O(n²)
func MinFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	copy(dp[0], matrix[0]) // base case：第一行的路径和就是元素本身
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			best := dp[i-1][j] // 正上方
			if j > 0 {
				best = min(best, dp[i-1][j-1]) // 左上方
			}
			if j < n-1 {
				best = min(best, dp[i-1][j+1]) // 右上方
			}
			dp[i][j] = matrix[i][j] + best
		}
	}
	// 终点可以是最后一行任意位置，取最小值
	ans := dp[n-1][0]
	for j := 1; j < n; j++ {
		ans = min(ans, dp[n-1][j])
	}
	return ans
}

// MinFallingPathSumOptimized 下降路径最小和（滚动数组空间优化版）
// 第 i 行只依赖第 i-1 行，用 prev/cur 两个一维数组滚动代替二维 dp 数组。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func MinFallingPathSumOptimized(matrix [][]int) int {
	n := len(matrix)
	prev := make([]int, n)
	copy(prev, matrix[0])
	cur := make([]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			best := prev[j]
			if j > 0 {
				best = min(best, prev[j-1])
			}
			if j < n-1 {
				best = min(best, prev[j+1])
			}
			cur[j] = matrix[i][j] + best
		}
		prev, cur = cur, prev // 交换复用，避免每行重新分配
	}
	ans := prev[0]
	for j := 1; j < n; j++ {
		ans = min(ans, prev[j])
	}
	return ans
}
