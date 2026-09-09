package besttimetobuyandsellstockwithtransactionfee

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	fee    int
	want   int
}{
	{name: "示例1", prices: []int{1, 3, 2, 8, 4, 9}, fee: 2, want: 8},
	{name: "示例2", prices: []int{1, 3, 7, 5, 10, 3}, fee: 3, want: 6},
	{name: "单日只能不交易", prices: []int{5}, fee: 2, want: 0},
	{name: "两天上涨", prices: []int{1, 5}, fee: 2, want: 2},
	{name: "手续费吃掉利润不交易", prices: []int{1, 3}, fee: 5, want: 0},
	{name: "一路下跌不交易", prices: []int{9, 8, 7, 6, 5}, fee: 1, want: 0},
	{name: "零手续费等同不限次交易", prices: []int{1, 3, 2, 8, 4, 9}, fee: 0, want: 13},
	{name: "大幅波动多次交易", prices: []int{1, 10, 1, 10, 1, 10}, fee: 1, want: 24},
	{name: "手续费为0且单调上涨", prices: []int{1, 2, 3, 4, 5}, fee: 0, want: 4},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices, tt.fee); got != tt.want {
				t.Errorf("MaxProfit(%v, %d) = %d, want %d", tt.prices, tt.fee, got, tt.want)
			}
		})
	}
}

func TestMaxProfitOptimized(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfitOptimized(tt.prices, tt.fee); got != tt.want {
				t.Errorf("MaxProfitOptimized(%v, %d) = %d, want %d", tt.prices, tt.fee, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{1, 3, 2, 8, 4, 9}
	for i := 0; i < b.N; i++ {
		MaxProfit(prices, 2)
	}
}

func BenchmarkMaxProfitOptimized(b *testing.B) {
	prices := []int{1, 3, 2, 8, 4, 9}
	for i := 0; i < b.N; i++ {
		MaxProfitOptimized(prices, 2)
	}
}
