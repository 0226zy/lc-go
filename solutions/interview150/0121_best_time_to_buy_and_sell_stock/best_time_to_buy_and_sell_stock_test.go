package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: [7,1,5,3,6,4]", []int{7, 1, 5, 3, 6, 4}, 5},
		{"示例2: 价格递减无利润", []int{7, 6, 4, 3, 1}, 0},

		// 边界：单元素，无法交易
		{"单元素", []int{1}, 0},
		{"单元素且为0", []int{0}, 0},

		// 边界：价格递增，首尾相减
		{"价格递增", []int{1, 2, 3, 4, 5}, 4},
		{"价格持平", []int{3, 3, 3, 3}, 0},

		// 边界：最低价在最后一天
		{"最低点在末尾", []int{5, 4, 3, 2, 1, 0}, 0},

		// 边界：先涨后跌
		{"先涨后跌", []int{2, 10, 1, 3}, 8},

		// 边界：极值
		{"最大利润为极差", []int{0, 10000}, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
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
		{"len=100000", generatePrices(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfit(bm.prices)
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
