package mincostclimbingstairs

// MinCostClimbingStairs 使用最小花费爬楼梯（标准 DP 数组版）
// cost[i] 是从第 i 级台阶出发的花费，每次可爬 1 或 2 级，求到达顶部（下标 n）的最小花费。
// dp[i] 表示到达第 i 级台阶的最低花费，dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	// base case：dp[0] = dp[1] = 0，可以从下标 0 或 1 免费出发
	for i := 2; i <= n; i++ {
		// 最后一步从 i-1 爬 1 步，或从 i-2 爬 2 步，取花费较小者
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// MinCostClimbingStairsOptimized 使用最小花费爬楼梯（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MinCostClimbingStairsOptimized(cost []int) int {
	prev2, prev1 := 0, 0 // 分别代表 dp[i-2]、dp[i-1]
	for i := 2; i <= len(cost); i++ {
		prev2, prev1 = prev1, min(prev1+cost[i-1], prev2+cost[i-2])
	}
	return prev1
}
