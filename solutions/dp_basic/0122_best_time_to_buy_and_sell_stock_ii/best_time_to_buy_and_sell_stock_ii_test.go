package besttimetobuyandsellstockii

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	{name: "示例1：两次交易", prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
	{name: "示例2：持续上涨", prices: []int{1, 2, 3, 4, 5}, want: 4},
	{name: "示例3：持续下跌", prices: []int{7, 6, 4, 3, 1}, want: 0},
	{name: "只有一天", prices: []int{5}, want: 0},
	{name: "两天上涨", prices: []int{1, 5}, want: 4},
	{name: "两天下跌", prices: []int{5, 1}, want: 0},
	{name: "价格全部相同", prices: []int{3, 3, 3, 3}, want: 0},
	{name: "当天买当天卖", prices: []int{2, 4, 1, 7}, want: 8},
	{name: "V形波动", prices: []int{6, 1, 3, 2, 4, 7}, want: 7},
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
	prices := []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9}
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}

func BenchmarkMaxProfitAlternative(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4, 2, 8, 1, 9}
	for i := 0; i < b.N; i++ {
		MaxProfitAlternative(prices)
	}
}
