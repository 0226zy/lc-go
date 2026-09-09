package climbingstairs

// ClimbStairs 爬楼梯（标准 DP 数组版）
// 每次可以爬 1 或 2 个台阶，求爬到第 n 阶的方法数。
// dp[i] 表示爬到第 i 阶的方法数，dp[i] = dp[i-1] + dp[i-2]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2 // base case：1 阶 1 种走法，2 阶 2 种走法
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 最后一步从 i-1 迈 1 步，或从 i-2 迈 2 步
	}
	return dp[n]
}

// ClimbStairsOptimized 爬楼梯（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ClimbStairsOptimized(n int) int {
	if n <= 2 {
		return n
	}
	prev2, prev1 := 1, 2 // 分别代表 dp[i-2]、dp[i-1]
	for i := 3; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
