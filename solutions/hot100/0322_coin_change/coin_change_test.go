package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "示例1：基本用例",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "示例2：无法凑成",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "示例3：金额为0",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "正好一枚",
			coins:  []int{1, 2, 5},
			amount: 5,
			want:   1,
		},
		{
			name:   "大面额无法使用",
			coins:  []int{5, 10, 25},
			amount: 1,
			want:   -1,
		},
		{
			name:   "贪心不可行",
			coins:  []int{1, 3, 4},
			amount: 6,
			want:   2,
		},
		{
			name:   "多种组合",
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			want:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("CoinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	amount := 11
	for i := 0; i < b.N; i++ {
		CoinChange(coins, amount)
	}
}
