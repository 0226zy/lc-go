package besttimetobuyandsellstockiii

import "math"

// MaxProfit 买卖股票的最佳时机 III
// 给定一个数组 prices，其中 prices[i] 是股票第 i 天的价格。
// 最多可以完成两笔交易（买入并卖出算一笔完整交易），且同一时间最多持有一股股票，
// 返回能获取的最大利润。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfit(prices []int) int {
	// 四个状态，均表示"当前账面上的最大利润"：
	// buy1  第一次买入后（值为负，表示花出去的最少钱）
	// sell1 第一次卖出后（第一笔交易的最大利润）
	// buy2  第二次买入后（用 sell1 的利润抵扣买入成本后的最大净收益）
	// sell2 第二次卖出后（两笔交易的最大利润）
	buy1, buy2 := math.MinInt32, math.MinInt32
	sell1, sell2 := 0, 0

	for _, price := range prices {
		buy1 = max(buy1, -price)      // 保持不动，或以 price 完成第一次买入
		sell1 = max(sell1, buy1+price) // 保持不动，或以 price 完成第一次卖出
		buy2 = max(buy2, sell1-price)  // 保持不动，或以 price 完成第二次买入
		sell2 = max(sell2, buy2+price) // 保持不动，或以 price 完成第二次卖出
	}

	// 第二次卖出的利润一定不低于第一次卖出（可以在同一天"卖完再买卖"，等价于只做一笔）
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
