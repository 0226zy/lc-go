package besttimetobuyandsellstock

// MaxProfit 买卖股票的最佳时机（标准 DP 数组版）
// 最多完成一笔交易，求最大利润。
// minPrice[i] 表示前 i 天的最低价格，dp[i] 表示前 i 天的最大利润；
// dp[i] = max(dp[i-1], prices[i]-minPrice[i])，答案为 dp[n-1]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProfit(prices []int) int {
	n := len(prices)
	minPrice := make([]int, n)
	dp := make([]int, n)
	minPrice[0] = prices[0] // base case：第一天最低成本就是当天价格，dp[0] = 0
	for i := 1; i < n; i++ {
		minPrice[i] = min(minPrice[i-1], prices[i])
		// 今天不卖（沿用历史最优） vs 今天卖（以历史最低价买入）
		dp[i] = max(dp[i-1], prices[i]-minPrice[i])
	}
	return dp[n-1]
}

func Dp(prices []int) int {
	n := len(prices)
	minPrice := prices[0]
	dp := 0
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		dp = max(dp, prices[i]-minPrice)
	}
	return dp
}

// MaxProfitOptimized 买卖股票的最佳时机（滚动变量空间优化版）
// dp[i]、minPrice[i] 都只依赖前一天状态，用两个滚动变量一次遍历即可。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfitOptimized(prices []int) int {
	minPrice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])      // 迄今最低买入价
		profit = max(profit, prices[i]-minPrice) // 今天卖的利润 vs 历史最优
	}
	return profit
}
