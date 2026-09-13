package targetsum

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: nums=[1,1,1,1,1],target=3", []int{1, 1, 1, 1, 1}, 3, 5},
		{"示例2: nums=[1],target=1", []int{1}, 1, 1},

		// 边界：单个数取不到 target
		{"单数取不到target", []int{1}, 2, 0},
		// 边界：单数取负号
		{"单数取负号成立", []int{2}, -2, 1},
		// 边界：包含 0，0 的两种符号都能成立，方案数翻倍
		{"包含0方案翻倍", []int{0, 1}, 1, 2},
		{"全为0", []int{0, 0, 0}, 0, 8},
		// 边界：target 绝对值超过总和，无解
		{"target超出总和", []int{1, 2, 3}, 100, 0},
		// 边界：全部取加号恰好等于 target
		{"全部取加号", []int{1, 2, 3}, 6, 1},
		// 边界：全部取负号恰好等于 target
		{"全部取负号", []int{1, 2, 3}, -6, 1},
		// 一般情况
		{"一般情况", []int{1, 2, 1}, 0, 2},
		{"大target为0", []int{1, 1, 1, 1}, 0, 6},
		{"接近上界规模", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0, 252},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSumWays(tt.nums, tt.target); got != tt.want {
				t.Errorf("FindTargetSumWays(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkFindTargetSumWays(b *testing.B) {
	benchmarks := []struct {
		name   string
		nums   []int
		target int
	}{
		{"n=10", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3},
		{"n=15", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3},
		{"n=20", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindTargetSumWays(bm.nums, bm.target)
			}
		})
	}
}
