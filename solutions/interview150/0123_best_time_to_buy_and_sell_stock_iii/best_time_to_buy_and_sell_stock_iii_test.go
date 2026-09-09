package besttimetobuyandsellstockiii

import "testing"

var maxProfitCases = []struct {
	name   string
	prices []int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1：两笔交易", prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, want: 6},
	{name: "示例2：一次买入卖出", prices: []int{1, 2, 3, 4, 5}, want: 4},
	{name: "示例3：一路下跌不交易", prices: []int{7, 6, 4, 3, 1}, want: 0},
	{name: "示例4：单元素无法交易", prices: []int{1}, want: 0},

	// 边界：两个元素
	{name: "两个元素上涨", prices: []int{1, 2}, want: 1},
	{name: "两个元素下跌", prices: []int{2, 1}, want: 0},

	// 边界：全部价格相同，利润为 0
	{name: "价格全部相同", prices: []int{5, 5, 5, 5}, want: 0},

	// 两笔交易分别在两段上涨区间
	{name: "两段上涨", prices: []int{1, 5, 2, 8}, want: 10},
	{name: "先涨后跌再涨", prices: []int{2, 4, 1, 7}, want: 8},

	// 一笔大交易优于两笔小交易，但两笔叠加更优
	{name: "经典叠加", prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9}, want: 13},

	// 边界：含 0 价格
	{name: "含0价格", prices: []int{0, 5, 0, 5}, want: 10},

	// 边界：极大价格值（约束上限 10^5）
	{name: "极大值价差", prices: []int{0, 100000}, want: 100000},
	{name: "极大值两笔", prices: []int{0, 100000, 0, 100000}, want: 200000},

	// 同一天卖完再买卖等价于不做操作
	{name: "锯齿波动", prices: []int{1, 2, 1, 2, 1, 2}, want: 2},

	// 先大跌后大涨
	{name: "V型反转", prices: []int{9, 1, 2, 8, 6, 9}, want: 10},
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

// TestMaxProfitStress 压力场景：10^5 长度数组，两个版本分别验证
func TestMaxProfitStress(t *testing.T) {
	n := 100000

	// 压力场景1：持续上涨 0,1,2,...,99999，一笔交易拿下全部涨幅
	up := make([]int, n)
	for i := range up {
		up[i] = i
	}
	t.Run("压力：10万长度持续上涨", func(t *testing.T) {
		if got := MaxProfit(up); got != n-1 {
			t.Errorf("MaxProfit(持续上涨) = %d, want %d", got, n-1)
		}
		if got := MaxProfitOptimized(up); got != n-1 {
			t.Errorf("MaxProfitOptimized(持续上涨) = %d, want %d", got, n-1)
		}
	})

	// 压力场景2：持续下跌，利润为 0
	down := make([]int, n)
	for i := range down {
		down[i] = n - i
	}
	t.Run("压力：10万长度持续下跌", func(t *testing.T) {
		if got := MaxProfit(down); got != 0 {
			t.Errorf("MaxProfit(持续下跌) = %d, want 0", got)
		}
		if got := MaxProfitOptimized(down); got != 0 {
			t.Errorf("MaxProfitOptimized(持续下跌) = %d, want 0", got)
		}
	})

	// 压力场景3：0/1 交替，两次交易最多各赚 1
	zigzag := make([]int, n)
	for i := range zigzag {
		zigzag[i] = i % 2
	}
	t.Run("压力：10万长度0/1交替", func(t *testing.T) {
		if got := MaxProfit(zigzag); got != 2 {
			t.Errorf("MaxProfit(0/1交替) = %d, want 2", got)
		}
		if got := MaxProfitOptimized(zigzag); got != 2 {
			t.Errorf("MaxProfitOptimized(0/1交替) = %d, want 2", got)
		}
	})
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := benchmarkPrices(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}

func BenchmarkMaxProfitOptimized(b *testing.B) {
	prices := benchmarkPrices(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitOptimized(prices)
	}
}

// benchmarkPrices 用确定性的伪随机模式生成价格数组，保证基准可复现
func benchmarkPrices(n int) []int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = (i*7919 + 104729) % 100000
	}
	return prices
}
