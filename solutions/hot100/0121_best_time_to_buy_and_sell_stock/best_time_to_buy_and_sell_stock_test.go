package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "示例1：有利润",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "示例2：无利润",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "边界：单天",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "边界：两天有利润",
			prices: []int{1, 5},
			want:   4,
		},
		{
			name:   "边界：两天无利润",
			prices: []int{5, 1},
			want:   0,
		},
		{
			name:   "先跌后涨",
			prices: []int{3, 2, 1, 5, 6, 2},
			want:   5,
		},
		{
			name:   "全部相等",
			prices: []int{3, 3, 3, 3},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}
