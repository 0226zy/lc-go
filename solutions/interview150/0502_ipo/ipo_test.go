package ipo

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		// LeetCode 官方示例
		{"示例1: k=2 w=0", 2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4},
		{"示例2: k=3 w=0", 3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6},

		// 边界：只有一个项目且资本足够
		{"单项目可做", 1, 5, []int{10}, []int{3}, 15},
		// 边界：只有一个项目但资本不足
		{"单项目不可做", 1, 1, []int{10}, []int{3}, 1},
		// 边界：k 大于项目数
		{"k大于项目数", 10, 0, []int{1, 2}, []int{0, 1}, 3},
		// 边界：初始资本足够做所有项目，k 限制次数
		{"资本充足选前k大利润", 2, 10, []int{1, 2, 3, 4}, []int{0, 1, 2, 3}, 17},
		// 边界：利润为 0
		{"利润为0", 2, 0, []int{0, 0}, []int{0, 0}, 0},
		// 边界：项目按利润降序给出
		// 先做 capital=0、profit=5 → w=5，再在剩余可做项目中选利润最大的两笔 4、3 → w=12
		{"利润降序给出", 3, 0, []int{5, 4, 3, 2, 1}, []int{0, 1, 2, 3, 4}, 12},
		// 边界：靠前期利润解锁后期高资本项目
		{"滚动解锁", 3, 0, []int{2, 3, 100}, []int{0, 1, 4}, 105},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaximizedCapital(tt.k, tt.w, tt.profits, tt.capital); got != tt.want {
				t.Errorf("FindMaximizedCapital(%d, %d, %v, %v) = %d, want %d",
					tt.k, tt.w, tt.profits, tt.capital, got, tt.want)
			}
		})
	}
}

func BenchmarkFindMaximizedCapital(b *testing.B) {
	// 构造 n 个项目，资本和利润均为伪随机序列
	n := 100000
	profits := make([]int, n)
	capital := make([]int, n)
	for i := 0; i < n; i++ {
		profits[i] = (i * 7919) % 10000
		capital[i] = (i * 104729) % 100000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMaximizedCapital(50000, 0, profits, capital)
	}
}
