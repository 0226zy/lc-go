package besttimetobuyandsellstockii

// MaxProfit 买卖股票的最佳时机 II（标准 DP 数组版）
// 每天可以买入或卖出，最多持有一股，可多次交易，求最大利润。
// dp[i][0] 表示第 i 天结束后不持有股票的最大利润，dp[i][1] 表示持有的最大利润。
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])（不动或今天卖出）
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])（不动或今天买入）
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0] // base case：第 0 天买入，利润为 -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 昨天就不持有，或今天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 昨天就持有，或今天买入
	}
	return dp[n-1][0] // 最后一天一定是不持有更优
}

// MaxProfitAlternative 买卖股票的最佳时机 II（贪心版）
// 上涨行情 a->b->c 整体持有赚 c-a，等价于分段赚 (b-a)+(c-b)，
// 因此累加所有相邻上涨差价即为最大利润（interview150 题单中本题的归类解法）。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfitAlternative(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			profit += diff // 收集每一段上涨
		}
	}
	return profit
}
