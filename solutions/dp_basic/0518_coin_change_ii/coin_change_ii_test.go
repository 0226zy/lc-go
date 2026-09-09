package coinchangeii

import "testing"

var changeCases = []struct {
	name   string
	amount int
	coins  []int
	want   int
}{
	{name: "官方示例1：amount=5,coins=[1,2,5]", amount: 5, coins: []int{1, 2, 5}, want: 4},
	{name: "官方示例2：amount=3,coins=[2]", amount: 3, coins: []int{2}, want: 0},
	{name: "官方示例3：amount=10,coins=[10]", amount: 10, coins: []int{10}, want: 1},
	{name: "金额为0只有一种组合", amount: 0, coins: []int{1, 2}, want: 1},
	{name: "硬币面额都大于金额", amount: 3, coins: []int{5, 10}, want: 0},
	{name: "只有面额1", amount: 4, coins: []int{1}, want: 1},
	{name: "三种小面额", amount: 4, coins: []int{1, 2, 3}, want: 4},
	{name: "大金额多种面额", amount: 100, coins: []int{1, 5, 10, 25, 50}, want: 292},
	{name: "较大数据", amount: 500, coins: []int{3, 5, 7, 8, 9, 10, 11}, want: 35502874},
}

func TestChange(t *testing.T) {
	for _, tt := range changeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Change(tt.amount, tt.coins); got != tt.want {
				t.Errorf("Change(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
			}
		})
	}
}

func TestChangeOptimized(t *testing.T) {
	for _, tt := range changeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeOptimized(tt.amount, tt.coins); got != tt.want {
				t.Errorf("ChangeOptimized(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
			}
		})
	}
}

func BenchmarkChange(b *testing.B) {
	coins := []int{1, 5, 10, 25, 50}
	for i := 0; i < b.N; i++ {
		Change(100, coins)
	}
}

func BenchmarkChangeOptimized(b *testing.B) {
	coins := []int{1, 5, 10, 25, 50}
	for i := 0; i < b.N; i++ {
		ChangeOptimized(100, coins)
	}
}
