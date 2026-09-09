package besttimetobuyandsellstock

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1：[7,1,5,3,6,4]", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
	{name: "示例2：价格递减无利润", prices: []int{7, 6, 4, 3, 1}, want: 0},

	// 边界：单元素，无法交易
	{name: "单元素", prices: []int{1}, want: 0},
	{name: "单元素且为0", prices: []int{0}, want: 0},

	// 边界：两天
	{name: "两天上涨", prices: []int{1, 2}, want: 1},
	{name: "两天下跌", prices: []int{2, 1}, want: 0},

	// 边界：价格递增，首尾相减
	{name: "价格递增", prices: []int{1, 2, 3, 4, 5}, want: 4},
	{name: "价格持平", prices: []int{3, 3, 3, 3}, want: 0},

	// 边界：最低价在最后一天
	{name: "最低点在末尾", prices: []int{5, 4, 3, 2, 1, 0}, want: 0},

	// 边界：先涨后跌 / 先跌后涨
	{name: "先涨后跌", prices: []int{2, 10, 1, 3}, want: 8},
	{name: "先跌后涨", prices: []int{3, 2, 6, 5, 0, 3}, want: 4},
	{name: "最低点在最后", prices: []int{4, 7, 2, 1}, want: 3},

	// 边界：极值
	{name: "最大利润为极差", prices: []int{0, 10000}, want: 10000},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

func TestMaxProfitOptimized(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfitOptimized(tt.prices); got != tt.want {
				t.Errorf("MaxProfitOptimized(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

// generatePrices 生成呈周期性波动的价格数组
func generatePrices(n int) []int {
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = (i * 37) % 10000
	}
	return prices
}

var maxProfitBenchmarks = []struct {
	name   string
	prices []int
}{
	{"len=10", []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9}},
	{"len=1000", generatePrices(1000)},
	{"len=100000", generatePrices(100000)},
}

func BenchmarkMaxProfit(b *testing.B) {
	for _, bm := range maxProfitBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfit(bm.prices)
			}
		})
	}
}

func BenchmarkMaxProfitOptimized(b *testing.B) {
	for _, bm := range maxProfitBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfitOptimized(bm.prices)
			}
		})
	}
}
