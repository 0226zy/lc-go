package triangle

// MinimumTotal 三角形最小路径和
// 给定三角形数组 triangle，从顶部出发，每一步可移动到下一行相邻位置（正下方或右下方），返回自顶向下的最小路径和。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func MinimumTotal(triangle [][]int) int {
	n := len(triangle)

	// dp[j] 表示从当前行第 j 个位置出发到底部的最小路径和
	// 初始化为最后一行的副本，不修改输入 triangle
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	// 自底向上：第 i 行第 j 列的最小路径和 = triangle[i][j] + min(dp[j], dp[j+1])
	// 其中 dp[j]、dp[j+1] 是下一行的结果，本行就地覆盖即可（滚动数组）
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}
