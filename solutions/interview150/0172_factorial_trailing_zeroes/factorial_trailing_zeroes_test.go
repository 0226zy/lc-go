package factorialtrailingzeroes

import "testing"

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 3! = 6 无尾随零", 3, 0},
		{"示例2: 5! = 120 一个尾随零", 5, 1},
		{"示例3: 0! = 1 无尾随零", 0, 0},

		// 边界：小于 5
		{"1 无尾随零", 1, 0},
		{"4 无尾随零", 4, 0},

		// 含 25、125 等多因子 5
		{"10! 两个尾随零", 10, 2},
		{"25! 六个尾随零", 25, 6},
		{"100! 二十四个尾随零", 100, 24},
		{"125! 三十一个尾随零", 125, 31},

		// 题目上界附近
		{"1000!", 1000, 249},
		{"10000!", 10000, 2499},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrailingZeroes(tt.n); got != tt.want {
				t.Errorf("TrailingZeroes(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkTrailingZeroes(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=5", 5},
		{"n=100", 100},
		{"n=10000", 10000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TrailingZeroes(bm.n)
			}
		})
	}
}
