package besttimetobuyandsellstockwithcooldown

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	{name: "空数组", prices: []int{}, want: 0},
	{name: "单日无法交易", prices: []int{1}, want: 0},
	{name: "官方示例", prices: []int{1, 2, 3, 0, 2}, want: 3},
	{name: "两天上涨", prices: []int{1, 2}, want: 1},
	{name: "持续下跌不交易", prices: []int{5, 4, 3, 2, 1}, want: 0},
	{name: "先跌后涨", prices: []int{2, 1, 4}, want: 3},
	{name: "多次交易含冷冻期", prices: []int{6, 1, 3, 2, 4, 7}, want: 6},
	{name: "价格相同", prices: []int{2, 2, 2, 2}, want: 0},
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

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{6, 1, 3, 2, 4, 7, 1, 2, 3, 0, 2, 5, 8, 3, 6, 9}
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}
