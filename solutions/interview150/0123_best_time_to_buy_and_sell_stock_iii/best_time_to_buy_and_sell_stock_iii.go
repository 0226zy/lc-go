package besttimetobuyandsellstockiii

import "math"

// MaxProfit 买卖股票的最佳时机 III（标准 DP 数组版）
// 最多完成两笔交易，同一时间最多持有一股，求最大利润。
// dp[i][k][0] 表示第 i 天结束后、至多完成 k 笔交易、不持有股票的最大利润，
// dp[i][k][1] 表示对应持有股票的最大利润（k 取 0..2）。
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])（不动或今天卖出）
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])（不动或今天买入开启第 k 笔）
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProfit(prices []int) int {
	n := len(prices)
	const K = 2
	dp := make([][K + 1][2]int, n)
	for k := 0; k <= K; k++ {
		dp[0][k][1] = -prices[0] // base case：第 0 天买入，利润为 -prices[0]
	}
	for i := 1; i < n; i++ {
		for k := 1; k <= K; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])   // 昨天就不持有，或今天卖出
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i]) // 昨天就持有，或今天买入
		}
	}
	return dp[n-1][K][0] // 最后一天一定是不持有更优
}

// MaxProfitOptimized 买卖股票的最佳时机 III（四状态滚动变量空间优化版）
// DP 表每天只依赖前一天且 k 只有 2 个取值，用 4 个滚动变量表示状态机：
// buy1 第一次买入后、sell1 第一次卖出后、buy2 第二次买入后、sell2 第二次卖出后的最大利润。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfitOptimized(prices []int) int {
	buy1, buy2 := math.MinInt, math.MinInt
	sell1, sell2 := 0, 0

	for _, price := range prices {
		buy1 = max(buy1, -price)       // 保持不动，或以 price 完成第一次买入
		sell1 = max(sell1, buy1+price) // 保持不动，或以 price 完成第一次卖出
		buy2 = max(buy2, sell1-price)  // 保持不动，或以 price 完成第二次买入
		sell2 = max(sell2, buy2+price) // 保持不动，或以 price 完成第二次卖出
	}

	// 第二次卖出的利润一定不低于第一次卖出（可以在同一天"卖完再买卖"，等价于只做一笔）
	return sell2
}
