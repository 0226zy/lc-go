package coinchange

import "testing"

var coinChangeCases = []struct {
	name   string
	coins  []int
	amount int
	want   int
}{
	{name: "官方示例1", coins: []int{1, 2, 5}, amount: 11, want: 3},
	{name: "官方示例2凑不出", coins: []int{2}, amount: 3, want: -1},
	{name: "官方示例3金额0", coins: []int{1}, amount: 0, want: 0},
	{name: "单种硬币整除", coins: []int{1}, amount: 2, want: 2},
	{name: "单种硬币不整除", coins: []int{3}, amount: 7, want: -1},
	{name: "硬币大于金额", coins: []int{5, 10}, amount: 3, want: -1},
	{name: "大面额组合", coins: []int{186, 419, 83, 408}, amount: 6249, want: 20},
	{name: "金额较大", coins: []int{1, 2, 5}, amount: 100, want: 20},
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

func TestCoinChangeAlternative(t *testing.T) {
	for _, tt := range coinChangeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChangeAlternative(tt.coins, tt.amount); got != tt.want {
				t.Errorf("CoinChangeAlternative(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	for i := 0; i < b.N; i++ {
		CoinChange(coins, 1000)
	}
}

func BenchmarkCoinChangeAlternative(b *testing.B) {
	coins := []int{1, 2, 5}
	for i := 0; i < b.N; i++ {
		CoinChangeAlternative(coins, 1000)
	}
}
