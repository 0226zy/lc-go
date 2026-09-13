package besttimetobuyandsellstockiv

import "math"

// MaxProfit 买卖股票的最佳时机 IV
// 给定数组 prices 和最多可完成的交易次数 k，同一时间最多持有一股，求最大利润。
// 当 k >= n/2 时，最多交易次数不再构成约束，退化为不限次数（累加所有上涨差价）。
// 否则使用状态机 DP：buy[j]/sell[j] 分别表示完成第 j 次买入/卖出后的最大利润。
// 时间复杂度: O(n·min(k,n))  空间复杂度: O(min(k,n))
func MaxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}
	// 一笔完整交易至少占两天，故最多只能完成 n/2 笔有意义的交易
	if k >= n/2 {
		return maxProfitUnlimited(prices)
	}

	// buy[j]：恰好完成第 j 次买入后的最大利润（持股）
	// sell[j]：恰好完成第 j 次卖出后的最大利润（不持股）
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	for j := 1; j <= k; j++ {
		buy[j] = math.MinInt / 2 // 避免后续 +price 溢出；表示尚未买入
	}

	for _, price := range prices {
		for j := 1; j <= k; j++ {
			// 第 j 次买入：沿用旧状态，或以第 j-1 次卖出的利润为成本今天买入
			buy[j] = max(buy[j], sell[j-1]-price)
			// 第 j 次卖出：沿用旧状态，或以第 j 次买入的持仓今天卖出
			sell[j] = max(sell[j], buy[j]+price)
		}
	}
	return sell[k]
}

// maxProfitUnlimited 不限交易次数：累加每一段上涨差价（与 0122 贪心相同）
func maxProfitUnlimited(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			profit += diff
		}
	}
	return profit
}
