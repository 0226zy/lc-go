package uniquepaths

// UniquePaths 不同路径（标准 DP 数组版）
// 机器人在 m x n 网格左上角，每次只能向右或向下走一步，求到达右下角的不同路径数。
// dp[i][j] 表示到达格子 (i,j) 的路径数，dp[i][j] = dp[i-1][j] + dp[i][j-1]。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1 // base case：第一行/第一列只有一种走法
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] // 只能从上方或左方走来
			}
		}
	}
	return dp[m-1][n-1]
}

// UniquePathsAlternative 不同路径（数学组合数法）
// 一共走 m+n-2 步，其中选 m-1 步向下，答案即组合数 C(m+n-2, m-1)。
// 逐步乘除 ans = ans * (n-1+i) / i，每步都能整除，避免阶乘溢出。
// 时间复杂度: O(min(m,n))  空间复杂度: O(1)
func UniquePathsAlternative(m int, n int) int {
	if m > n {
		m, n = n, m // 用较小的一边做循环，减少次数
	}
	ans := 1
	for i := 1; i <= m-1; i++ {
		ans = ans * (n - 1 + i) / i // 逐步计算 C(m+n-2, m-1)
	}
	return ans
}
