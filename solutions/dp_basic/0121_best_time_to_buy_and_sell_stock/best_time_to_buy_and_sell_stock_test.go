package besttimetobuyandsellstock

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	{name: "普通上涨（示例1）", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
	{name: "一路下跌（示例2）", prices: []int{7, 6, 4, 3, 1}, want: 0},
	{name: "只有一天", prices: []int{5}, want: 0},
	{name: "两天上涨", prices: []int{1, 2}, want: 1},
	{name: "两天下跌", prices: []int{2, 1}, want: 0},
	{name: "价格不变", prices: []int{3, 3, 3, 3}, want: 0},
	{name: "先跌后涨", prices: []int{3, 2, 6, 5, 0, 3}, want: 4},
	{name: "最低点在最后", prices: []int{4, 7, 2, 1}, want: 3},
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

var benchPrices = func() []int {
	// 构造 10 万天的价格序列用于基准测试
	prices := make([]int, 100000)
	for i := range prices {
		prices[i] = (i*7919)%10000 + 1
	}
	return prices
}()

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxProfit(benchPrices)
	}
}

func BenchmarkMaxProfitOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxProfitOptimized(benchPrices)
	}
}
