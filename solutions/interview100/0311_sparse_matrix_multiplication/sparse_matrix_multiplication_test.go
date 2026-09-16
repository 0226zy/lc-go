package sparsematrixmultiplication

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		mat1 [][]int
		mat2 [][]int
		want [][]int
	}{
		// LeetCode 官方示例
		{
			"示例1: 2x3 与 3x3 稀疏矩阵",
			[][]int{{1, 0, 0}, {-1, 0, 3}},
			[][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			[][]int{{7, 0, 0}, {-7, 0, 3}},
		},
		{
			"示例2: 1x1 零矩阵",
			[][]int{{0}},
			[][]int{{0}},
			[][]int{{0}},
		},

		// 边界：1x1 非零
		{
			"1x1 非零矩阵",
			[][]int{{5}},
			[][]int{{-3}},
			[][]int{{-15}},
		},

		// 边界：全零矩阵
		{
			"全零矩阵相乘",
			[][]int{{0, 0}, {0, 0}},
			[][]int{{1, 2}, {3, 4}},
			[][]int{{0, 0}, {0, 0}},
		},

		// 边界：单位矩阵
		{
			"与单位矩阵相乘",
			[][]int{{1, 0}, {0, 1}},
			[][]int{{5, 6}, {7, 8}},
			[][]int{{5, 6}, {7, 8}},
		},

		// 典型场景：稠密矩阵
		{
			"稠密矩阵相乘",
			[][]int{{1, 2}, {3, 4}},
			[][]int{{5, 6}, {7, 8}},
			[][]int{{19, 22}, {43, 50}},
		},

		// 典型场景：含负数
		{
			"含负数元素",
			[][]int{{-1, 2}, {0, -3}},
			[][]int{{4, -5}, {-6, 7}},
			[][]int{{-16, 19}, {18, -21}},
		},

		// 边界：非方阵
		{
			"1x3 与 3x1（点积）",
			[][]int{{1, 2, 3}},
			[][]int{{4}, {5}, {6}},
			[][]int{{32}},
		},
		{
			"3x1 与 1x3（外积）",
			[][]int{{1}, {2}, {3}},
			[][]int{{4, 5, 6}},
			[][]int{{4, 5, 6}, {8, 10, 12}, {12, 15, 18}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.mat1, tt.mat2)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.mat1, tt.mat2, got, tt.want)
			}
		})
	}
}

func BenchmarkMultiply(b *testing.B) {
	// 构造 100x100 稀疏矩阵（约 5% 非零元素）
	size := 100
	mat1 := make([][]int, size)
	mat2 := make([][]int, size)
	for i := 0; i < size; i++ {
		mat1[i] = make([]int, size)
		mat2[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if (i*size+j)%20 == 0 {
				mat1[i][j] = (i + j) % 7
				mat2[i][j] = (i*j)%5 + 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Multiply(mat1, mat2)
	}
}
