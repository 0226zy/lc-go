package nqueensii

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: n=4", 4, 2},
		{"示例2: n=1", 1, 1},

		// 经典已知方案数：1,0,0,2,10,4,40,92,352
		{"n=2 无解", 2, 0},
		{"n=3 无解", 3, 0},
		{"n=5 有10解", 5, 10},
		{"n=6 有4解", 6, 4},
		{"n=7 有40解", 7, 40},
		{"n=8 有92解", 8, 92},
		{"n=9 有352解", 9, 352},
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
