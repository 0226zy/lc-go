package besttimetobuyandsellstockii

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1：两次交易", prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
	{name: "示例2：持续上涨", prices: []int{1, 2, 3, 4, 5}, want: 4},
	{name: "示例3：持续下跌", prices: []int{7, 6, 4, 3, 1}, want: 0},

	// 边界：元素过少，无法交易或只够一笔
	{name: "单元素", prices: []int{3}, want: 0},
	{name: "两个元素下跌", prices: []int{2, 1}, want: 0},
	{name: "两个元素上涨", prices: []int{1, 2}, want: 1},

	// 边界：价格持平，无利可图
	{name: "价格持平", prices: []int{3, 3, 3, 3}, want: 0},

	// 边界：锯齿形波动，每次上涨都赚
	{name: "锯齿波动", prices: []int{1, 5, 1, 5, 1, 5}, want: 12},
	{name: "当天买当天卖", prices: []int{2, 4, 1, 7}, want: 8},
	{name: "V形波动", prices: []int{6, 1, 3, 2, 4, 7}, want: 7},

	// 边界：极值
	{name: "最大利润累加", prices: []int{0, 10000, 0, 10000}, want: 20000},
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

func TestMaxProfitAlternative(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfitAlternative(tt.prices); got != tt.want {
				t.Errorf("MaxProfitAlternative(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	benchmarks := []struct {
		name   string
		prices []int
	}{
		{"len=10", []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9}},
		{"len=1000", generatePrices(1000)},
		{"len=30000", generatePrices(30000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfit(bm.prices)
			}
		})
	}
}

func BenchmarkMaxProfitAlternative(b *testing.B) {
	benchmarks := []struct {
		name   string
		prices []int
	}{
		{"len=10", []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9}},
		{"len=1000", generatePrices(1000)},
		{"len=30000", generatePrices(30000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfitAlternative(bm.prices)
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
