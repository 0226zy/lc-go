package besttimetobuyandsellstockwithtransactionfee

// MaxProfit 买卖股票的最佳时机含手续费（标准 DP 数组版）
// 每天可以无限次交易，但最多持有一股，每笔交易（买入+卖出）支付 fee 手续费。
// dp[i][0] 表示第 i 天结束后不持股的最大利润，dp[i][1] 表示持股的最大利润。
// 转移：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
//
//	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
//
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0] // base case：第 0 天不买 / 买入
	for i := 1; i < n; i++ {
		// 不持股：昨天就不持有，或昨天持有今天卖出（付手续费）
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		// 持股：昨天就持有，或昨天不持有今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0] // 最终手里不能留股票
}

// MaxProfitOptimized 买卖股票的最佳时机含手续费（滚动变量空间优化版）
// 每天的状态只依赖前一天，用 cash/hold 两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfitOptimized(prices []int, fee int) int {
	cash, hold := 0, -prices[0] // 分别代表前一天的不持股、持股最大利润
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee) // 今天卖出（或不动）
		hold = max(hold, cash-prices[i])     // 今天买入（或继续持有）
		// 注意 cash 先更新：同一天「卖出又买回」等效于没交易，利润不变，不影响正确性
	}
	return cash
}
