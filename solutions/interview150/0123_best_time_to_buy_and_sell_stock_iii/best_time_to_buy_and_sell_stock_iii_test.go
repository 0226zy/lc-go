package besttimetobuyandsellstockiii

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: [3,3,5,0,0,3,1,4] 两笔交易", []int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{"示例2: [1,2,3,4,5] 一次买入卖出", []int{1, 2, 3, 4, 5}, 4},
		{"示例3: [7,6,4,3,1] 一路下跌不交易", []int{7, 6, 4, 3, 1}, 0},
		{"示例4: 单元素 [1] 无法交易", []int{1}, 0},

		// 边界：只有一个或两个元素
		{"两个元素上涨 [1,2]", []int{1, 2}, 1},
		{"两个元素下跌 [2,1]", []int{2, 1}, 0},

		// 边界：全部价格相同，利润为 0
		{"价格全部相同 [5,5,5,5]", []int{5, 5, 5, 5}, 0},

		// 两笔交易分别在两段上涨区间
		{"两段上涨 [1,5,2,8]", []int{1, 5, 2, 8}, 10},
		{"先涨后跌再涨 [2,4,1,7]", []int{2, 4, 1, 7}, 8},

		// 一笔大交易优于两笔小交易，但两笔叠加更优
		{"经典叠加 [1,2,4,2,5,7,2,4,9]", []int{1, 2, 4, 2, 5, 7, 2, 4, 9}, 13},

		// 边界：含 0 价格
		{"含0价格 [0,5,0,5]", []int{0, 5, 0, 5}, 10},

		// 边界：极大价格值（约束上限 10^5）
		{"极大值价差 [0,100000]", []int{0, 100000}, 100000},
		{"极大值两笔 [0,100000,0,100000]", []int{0, 100000, 0, 100000}, 200000},

		// 同一天卖完再买卖等价于不做操作
		{"锯齿波动 [1,2,1,2,1,2]", []int{1, 2, 1, 2, 1, 2}, 2},

		// 先大跌后大涨
		{"V型反转 [9,1,2,8,6,9]", []int{9, 1, 2, 8, 6, 9}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}

// TestMaxProfitStress 压力场景：10^5 长度数组
func TestMaxProfitStress(t *testing.T) {
	n := 100000

	// 压力场景1：持续上涨 0,1,2,...,99999，一笔交易拿下全部涨幅
	up := make([]int, n)
	for i := range up {
		up[i] = i
	}
	t.Run("压力: 10万长度持续上涨", func(t *testing.T) {
		if got := MaxProfit(up); got != n-1 {
			t.Errorf("MaxProfit(持续上涨) = %v, want %v", got, n-1)
		}
	})

	// 压力场景2：持续下跌，利润为 0
	down := make([]int, n)
	for i := range down {
		down[i] = n - i
	}
	t.Run("压力: 10万长度持续下跌", func(t *testing.T) {
		if got := MaxProfit(down); got != 0 {
			t.Errorf("MaxProfit(持续下跌) = %v, want 0", got)
		}
	})

	// 压力场景3：0/1 交替，两次交易最多各赚 1
	zigzag := make([]int, n)
	for i := range zigzag {
		zigzag[i] = i % 2
	}
	t.Run("压力: 10万长度0/1交替", func(t *testing.T) {
		if got := MaxProfit(zigzag); got != 2 {
			t.Errorf("MaxProfit(0/1交替) = %v, want 2", got)
		}
	})
}

func BenchmarkMaxProfit(b *testing.B) {
	benchmarks := []struct {
		name string
		make func(n int) []int
	}{
		{"1万长度持续上涨", func(n int) []int {
			prices := make([]int, n)
			for i := range prices {
				prices[i] = i
			}
			return prices
		}},
		{"1万长度锯齿波动", func(n int) []int {
			prices := make([]int, n)
			for i := range prices {
				prices[i] = i % 1000
			}
			return prices
		}},
		{"10万长度随机模式", func(n int) []int {
			prices := make([]int, n)
			for i := range prices {
				// 用确定性的伪随机模式，保证基准可复现
				prices[i] = (i*7919 + 104729) % 100000
			}
			return prices
		}},
	}

	for _, bm := range benchmarks {
		n := 10000
		if bm.name == "10万长度随机模式" {
			n = 100000
		}
		prices := bm.make(n)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfit(prices)
			}
		})
	}
}
