package factorcombinations

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestGetFactors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: n=1", 1, [][]int{}},
		{"示例2: n=12", 12, [][]int{{2, 6}, {2, 2, 3}, {3, 4}}},
		{"示例3: n=37 质数", 37, [][]int{}},

		// 边界：2 和 3 都是质数
		{"n=2 质数", 2, [][]int{}},
		{"n=3 质数", 3, [][]int{}},

		// 边界：最小有组合的合数
		{"n=4", 4, [][]int{{2, 2}}},

		// 典型：多种分解方式
		{"n=8", 8, [][]int{{2, 4}, {2, 2, 2}}},
		{
			"n=32 幂次分解",
			32,
			[][]int{{2, 16}, {2, 2, 8}, {2, 2, 2, 4}, {2, 2, 2, 2, 2}, {2, 4, 4}, {4, 8}},
		},
		{
			"n=24",
			24,
			[][]int{{2, 12}, {2, 2, 6}, {2, 2, 2, 3}, {2, 3, 4}, {3, 8}, {4, 6}},
		},

		// 典型：大质数的倍数
		{"n=13 质数", 13, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFactors(tt.n)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("GetFactors(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkGetFactors(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=12", 12},
		{"n=360 高合数", 360},
		{"n=12960", 12960},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetFactors(bm.n)
			}
		})
	}
}
