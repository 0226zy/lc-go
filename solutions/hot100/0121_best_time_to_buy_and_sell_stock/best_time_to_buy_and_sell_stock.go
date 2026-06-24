package besttimetobuyandsellstock

import "math"

// MaxProfit 买卖股票的最佳时机
// 只能买卖一次，求最大利润。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProfit(prices []int) int {
	ret := 0
	minPrice := math.MaxInt64
	for _, price := range prices {
		ret = max(ret, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return ret
}
