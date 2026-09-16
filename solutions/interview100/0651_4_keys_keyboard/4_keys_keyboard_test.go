package fourkeyskeyboard

import "testing"

func TestMaxA(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: n=3 全按A", 3, 3},
		{"示例2: n=7 复制粘贴翻倍", 7, 9},

		// 边界：按键次数太少，复制粘贴没有意义，只能一直按 A
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=4", 4, 4},
		{"n=5", 5, 5},
		{"n=6", 6, 6},

		// 典型场景：复制粘贴开始优于纯按 A
		{"n=8", 8, 12},
		{"n=9", 9, 16},
		{"n=10", 10, 20},
		{"n=11", 11, 27},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxA(tt.n); got != tt.want {
				t.Errorf("MaxA(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxA(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=10", 10},
		{"n=25", 25},
		{"n=50", 50},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxA(bm.n)
			}
		})
	}
}
