package besttimetobuyandsellstockii

// MaxProfit 买卖股票的最佳时机 II
// 给定一个数组 prices，prices[i] 表示第 i 天的股票价格。你可以进行任意多次交易
// （买入、卖出可以交替多次），但同一时刻最多持有一股，且必须先卖出后才能再次买入。
// 返回你能获得的最大利润。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func MaxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		// 只要今天比昨天贵，就在昨天买入、今天卖出
		// 所有上涨的"每一段"利润都被累加，等价于在谷底买入、峰顶卖出
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
