package spiralmatrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3方阵", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"示例2: 3x4矩阵", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},

		// 边界：空矩阵
		{"空矩阵", [][]int{}, nil},
		// 边界：单行
		{"单行矩阵", [][]int{{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		// 边界：单列
		{"单列矩阵", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		// 边界：单行左折回
		{"单行2x1", [][]int{{1}, {2}}, []int{1, 2}},
		{"单元素", [][]int{{7}}, []int{7}},
		// 边界：宽矩阵（只有两行）
		{"2x3矩阵", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 6, 5, 4}},
		// 边界：高矩阵（只有两列）
		{"3x2矩阵", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 4, 6, 5, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}

func BenchmarkSpiralOrder(b *testing.B) {
	for _, size := range []int{10, 100, 500} {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SpiralOrder(makeSquareMatrix(size))
			}
		})
	}
}

// makeSquareMatrix 生成 n x n 的连续数字矩阵
func makeSquareMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j + 1
		}
	}
	return matrix
}
