package coinchange

import "testing"

var coinChangeCases = []struct {
	name   string
	coins  []int
	amount int
	want   int
}{
	// LeetCode 官方示例
	{name: "示例1：coins=[1,2,5] amount=11", coins: []int{1, 2, 5}, amount: 11, want: 3},
	{name: "示例2：无法凑成", coins: []int{2}, amount: 3, want: -1},
	{name: "示例3：amount=0", coins: []int{1}, amount: 0, want: 0},

	// 边界
	{name: "单币恰好", coins: []int{5}, amount: 5, want: 1},
	{name: "单币倍数", coins: []int{5}, amount: 15, want: 3},
	{name: "单币无法整除", coins: []int{5}, amount: 7, want: -1},
	{name: "含1总能凑成", coins: []int{1, 3, 4}, amount: 6, want: 2},
	{name: "贪心非最优", coins: []int{1, 3, 4}, amount: 6, want: 2}, // 4+1+1=3 劣于 3+3
	{name: "大面额优先仍非最优", coins: []int{186, 419, 83, 408}, amount: 6249, want: 20},
	{name: "多种面额", coins: []int{1, 2, 5, 10}, amount: 18, want: 4},
	{name: "amount=1", coins: []int{1, 2}, amount: 1, want: 1},
}

func TestCoinChange(t *testing.T) {
	for _, tt := range coinChangeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("CoinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5, 10, 25}
	amount := 9999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CoinChange(coins, amount)
	}
}
