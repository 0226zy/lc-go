package nqueensii

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: n=4 有两种解法", 4, 2},
		{"示例2: n=1 有一种解法", 1, 1},

		// 边界与已知方案数：1, 0, 0, 2, 10, 4, 40, 92, 352
		{"边界: n=2 无解", 2, 0},
		{"边界: n=3 无解", 3, 0},
		{"已知解: n=5 有10种解法", 5, 10},
		{"已知解: n=6 有4种解法", 6, 4},
		{"已知解: n=7 有40种解法", 7, 40},
		{"已知解: n=8 有92种解法", 8, 92},
		{"已知解: n=9 有352种解法", 9, 352},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalNQueens(tt.n); got != tt.want {
				t.Errorf("TotalNQueens(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkTotalNQueens(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=7", 7},
		{"n=8", 8},
		{"n=9", 9},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TotalNQueens(bm.n)
			}
		})
	}
}
