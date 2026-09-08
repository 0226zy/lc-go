package besttimetobuyandsellstock

// MaxProfit 买卖股票的最佳时机
// 给定一个数组 prices，prices[i] 表示第 i 天的股票价格。最多只能完成一笔交易
// （买入一次并卖出一次），且必须先买入后卖出，返回你能获得的最大利润。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func MaxProfit(prices []int) int {
	minPrice := prices[0] // 记录到目前为止看到的最低价格
	maxProfit := 0        // 记录到目前为止能获得的最大利润
	for _, price := range prices {
		if price < minPrice {
			// 出现了更低的买入价，更新最低价格
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// 以历史最低价买入、今天卖出，利润更大则更新
			maxProfit = profit
		}
	}
	return maxProfit
}
