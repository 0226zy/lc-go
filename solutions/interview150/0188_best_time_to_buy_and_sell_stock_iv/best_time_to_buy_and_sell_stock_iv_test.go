package besttimetobuyandsellstockiv

import "testing"

var maxProfitCases = []struct {
	name   string
	k      int
	prices []int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1：k=2 两段上涨", k: 2, prices: []int{2, 4, 1}, want: 2},
	{name: "示例2：k=2 经典", k: 2, prices: []int{3, 2, 6, 5, 0, 3}, want: 7},

	// 边界：k=0 无法交易
	{name: "k=0 无法交易", k: 0, prices: []int{1, 2, 3}, want: 0},

	// 边界：单日 / 空数组
	{name: "单日无法交易", k: 1, prices: []int{5}, want: 0},
	{name: "空价格数组", k: 2, prices: []int{}, want: 0},

	// 边界：k=1 退化为一次交易
	{name: "k=1 一次交易", k: 1, prices: []int{3, 2, 6, 5, 0, 3}, want: 4},
	{name: "k=1 持续上涨", k: 1, prices: []int{1, 2, 3, 4, 5}, want: 4},

	// 一路下跌
	{name: "一路下跌利润为0", k: 3, prices: []int{7, 6, 4, 3, 1}, want: 0},

	// 价格全部相同
	{name: "价格全部相同", k: 2, prices: []int{5, 5, 5, 5}, want: 0},

	// k 足够大时退化为不限次数
	{name: "k很大吃满所有上坡", k: 100, prices: []int{1, 2, 1, 2, 1, 2}, want: 3},
	{name: "k>=n/2 持续上涨", k: 10, prices: []int{1, 2, 3, 4, 5}, want: 4},

	// 多段上涨，k 限制交易次数
	{name: "三段上涨但k=2", k: 2, prices: []int{1, 5, 2, 8, 3, 10}, want: 14},
	{name: "三段上涨k=3", k: 3, prices: []int{1, 5, 2, 8, 3, 10}, want: 17},

	// 含 0 价格
	{name: "含0价格", k: 2, prices: []int{0, 5, 0, 5}, want: 10},

	// 极大值
	{name: "极大值两笔", k: 2, prices: []int{0, 100000, 0, 100000}, want: 200000},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range maxProfitCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.k, tt.prices); got != tt.want {
				t.Errorf("MaxProfit(%d, %v) = %d, want %d", tt.k, tt.prices, got, tt.want)
			}
		})
	}
}

func TestMaxProfitStress(t *testing.T) {
	n := 100000
	up := make([]int, n)
	for i := range up {
		up[i] = i
	}
	t.Run("压力：10万持续上涨 k=2", func(t *testing.T) {
		if got := MaxProfit(2, up); got != n-1 {
			t.Errorf("MaxProfit(2, 持续上涨) = %d, want %d", got, n-1)
		}
	})
	t.Run("压力：10万持续上涨 k很大", func(t *testing.T) {
		if got := MaxProfit(n, up); got != n-1 {
			t.Errorf("MaxProfit(n, 持续上涨) = %d, want %d", got, n-1)
		}
	})

	zigzag := make([]int, n)
	for i := range zigzag {
		zigzag[i] = i % 2
	}
	t.Run("压力：10万0/1交替 k=2", func(t *testing.T) {
		if got := MaxProfit(2, zigzag); got != 2 {
			t.Errorf("MaxProfit(2, 0/1交替) = %d, want 2", got)
		}
	})
	t.Run("压力：10万0/1交替 k很大", func(t *testing.T) {
		want := n / 2 // 每一对 0->1 赚 1
		if got := MaxProfit(n, zigzag); got != want {
			t.Errorf("MaxProfit(n, 0/1交替) = %d, want %d", got, want)
		}
	})
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := make([]int, 10000)
	for i := range prices {
		prices[i] = (i*7919 + 104729) % 100000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfit(100, prices)
	}
}

func BenchmarkMaxProfitUnlimitedPath(b *testing.B) {
	prices := make([]int, 100000)
	for i := range prices {
		prices[i] = (i*7919 + 104729) % 100000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfit(len(prices), prices) // 走不限次数分支
	}
}
