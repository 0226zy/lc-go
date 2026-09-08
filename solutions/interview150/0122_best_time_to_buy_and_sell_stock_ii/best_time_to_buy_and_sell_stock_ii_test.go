package besttimetobuyandsellstockii

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: [7,1,5,3,6,4]", []int{7, 1, 5, 3, 6, 4}, 7},
		{"示例2: [1,2,3,4,5]", []int{1, 2, 3, 4, 5}, 4},
		{"示例3: [7,6,4,3,1]", []int{7, 6, 4, 3, 1}, 0},

		// 边界：单元素，无法交易
		{"单元素", []int{3}, 0},
		{"两个元素下跌", []int{2, 1}, 0},
		{"两个元素上涨", []int{1, 2}, 1},

		// 边界：价格持平
		{"价格持平", []int{3, 3, 3, 3}, 0},

		// 边界：锯齿形波动，每次都赚钱
		{"锯齿波动", []int{1, 5, 1, 5, 1, 5}, 12},

		// 边界：极值
		{"最大利润累加", []int{0, 10000, 0, 10000}, 20000},
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

// generatePrices 生成呈周期性波动的价格数组
func generatePrices(n int) []int {
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = (i * 37) % 10000
	}
	return prices
}
