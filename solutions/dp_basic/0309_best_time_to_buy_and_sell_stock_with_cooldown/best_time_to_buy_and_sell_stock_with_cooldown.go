package besttimetobuyandsellstockwithcooldown

// MaxProfit 最佳买卖股票时机含冷冻期（标准 DP 数组版）
// 可以多次买卖，但卖出后有一天冷冻期不能买入，求最大利润。
// 状态机 DP：dp[i][0] 持有股票，dp[i][1] 不持有且处于冷冻期（当天卖出），
// dp[i][2] 不持有且非冷冻期；dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0 // 不足两天，无法完成任何交易
	}
	dp := make([][3]int, n)
	dp[0][0] = -prices[0] // base case：第 0 天买入
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) // 保持持有，或冷冻期结束后买入
		dp[i][1] = dp[i-1][0] + prices[i]                // 昨天持有，今天卖出（进入冷冻期）
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])           // 冷冻期结束，或继续休息
	}
	// 最后一天不应持有股票，在两种不持有状态中取较大者
	return max(dp[n-1][1], dp[n-1][2])
}
