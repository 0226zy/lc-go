package painthouse

// MinCost 粉刷房子
// 一排 n 个房子，每个房子可刷红、蓝、绿三种颜色之一，costs[i][j] 表示第 i 个房子
// 刷成颜色 j 的花费；要求相邻房子颜色不同，返回刷完所有房子的最小花费。
// 时间复杂度: O(n) 颜色数固定为 3  空间复杂度: O(1) 滚动数组
func MinCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	// dp[j] 表示刷到当前房子、且当前房子刷颜色 j 时的最小总花费
	dp := [3]int{costs[0][0], costs[0][1], costs[0][2]}
	for i := 1; i < len(costs); i++ {
		// 当前房子刷颜色 j 时，前一个房子只能刷另外两种颜色，取较小者
		r := costs[i][0] + min(dp[1], dp[2])
		b := costs[i][1] + min(dp[0], dp[2])
		g := costs[i][2] + min(dp[0], dp[1])
		dp = [3]int{r, b, g}
	}
	return min(dp[0], dp[1], dp[2])
}
