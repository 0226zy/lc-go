package setmatrixzeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 中心为0的3x3", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"示例2: 首行首列含0", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},

		// 边界：无 0
		{"无0矩阵", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		// 边界：单元素为 0
		{"单元素为0", [][]int{{0}}, [][]int{{0}}},
		// 边界：单元素非 0
		{"单元素非0", [][]int{{5}}, [][]int{{5}}},
		// 边界：单行含 0
		{"单行含0", [][]int{{1, 0, 3}}, [][]int{{0, 0, 0}}},
		// 边界：单列含 0
		{"单列含0", [][]int{{1}, {0}, {3}}, [][]int{{0}, {0}, {0}}},
		// 边界：所有元素都是 0
		{"全0矩阵", [][]int{{0, 0}, {0, 0}}, [][]int{{0, 0}, {0, 0}}},
		// 边界：0 位于角落
		{"0在左下角", [][]int{{1, 2, 3}, {4, 5, 6}, {0, 8, 9}}, [][]int{{0, 2, 3}, {0, 5, 6}, {0, 0, 0}}},
		{"0在右上角", [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}, [][]int{{0, 0, 0}, {4, 5, 0}, {7, 8, 0}}},
		// 边界：首行首列同时含 0
		{"首行首列均有0", [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}, [][]int{{0, 0, 0}, {0, 4, 5}, {0, 7, 8}}},
		// 边界：含负数
		{"含负数与0", [][]int{{-1, 2, 0}, {4, -5, 6}, {7, 8, 9}}, [][]int{{0, 0, 0}, {4, -5, 0}, {7, 8, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			SetZeroes(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("SetZeroes(%v) = %v, want %v", tt.matrix, matrix, tt.want)
			}
		})
	}
}

func BenchmarkSetZeroes(b *testing.B) {
	for _, size := range []int{10, 100, 500} {
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 每次重新生成矩阵，避免上一轮置零影响本轮
				SetZeroes(makeMatrixWithZero(size))
			}
		})
	}
}

// makeMatrixWithZero 生成 n x n 矩阵，对角线上放一个 0
func makeMatrixWithZero(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j + 1
		}
		matrix[i][i] = 0
	}
	return matrix
}
