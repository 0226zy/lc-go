package paintfence

import "testing"

func TestNumWays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: n=3 k=2", 3, 2, 6},
		{"示例2: n=1 k=1", 1, 1, 1},
		{"示例3: n=7 k=2", 7, 2, 42},

		// 边界：n 为 0 或 1
		{"没有木柱", 0, 3, 0},
		{"单木柱多颜色", 1, 5, 5},

		// 边界：n=2 时无约束
		{"两根木柱两种颜色", 2, 2, 4},
		{"两根木柱三种颜色", 2, 3, 9},

		// 边界：k=1 时超过两根即无解
		{"一种颜色一根木柱", 1, 1, 1},
		{"一种颜色两根木柱", 2, 1, 1},
		{"一种颜色三根木柱无解", 3, 1, 0},
		{"一种颜色十根木柱无解", 10, 1, 0},

		// 典型场景
		{"n=4 k=3", 4, 3, 66},
		{"n=3 k=3", 3, 3, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumWays(tt.n, tt.k); got != tt.want {
				t.Errorf("NumWays(%d, %d) = %d, want %d", tt.n, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkNumWays(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"n=10,k=3", 10, 3},
		{"n=50,k=100", 50, 100},
		{"n=50,k=100000", 50, 100000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NumWays(bm.n, bm.k)
			}
		})
	}
}
