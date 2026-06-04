package spiralmatrix

import (
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
		{"示例1: 3x3矩阵", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"示例2: 3x4矩阵", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},

		// 边界：单元素矩阵
		{"单元素矩阵", [][]int{{1}}, []int{1}},
		{"单行矩阵", [][]int{{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"单列矩阵", [][]int{{1}, {2}, {3}, {4}}, []int{1, 2, 3, 4}},

		// 边界：全为负数
		{"全为负数", [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}}, []int{-1, -2, -3, -6, -9, -8, -7, -4, -5}},

		// 边界：含0
		{"含0", [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}, []int{0, 1, 2, 5, 8, 7, 6, 3, 4}},

		// 边界：2x2矩阵
		{"2x2矩阵", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},

		// 边界：1x1矩阵
		{"1x1矩阵", [][]int{{5}}, []int{5}},

		// 较大矩阵
		{"4x4矩阵", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
		{"5x3矩阵", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}}, []int{1, 2, 3, 6, 9, 12, 15, 14, 13, 10, 7, 4, 5, 8, 11}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}

func BenchmarkSpiralOrder(b *testing.B) {
	benchmarks := []struct {
		name   string
		matrix [][]int
	}{
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{"3x4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
		{"10x10", generateMatrix(10)},
		{"100x100", generateMatrix(100)},
		{"1000x1000", generateMatrix(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SpiralOrder(bm.matrix)
			}
		})
	}
}

// generateMatrix 生成大小为 n x n 的矩阵，填充 1 到 n*n
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j + 1
		}
	}
	return matrix
}
