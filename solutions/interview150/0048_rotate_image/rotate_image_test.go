package rotateimage

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3方阵", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{"示例2: 4x4方阵", [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},

		// 边界：单元素矩阵
		{"单元素", [][]int{{1}}, [][]int{{1}}},
		// 边界：2x2 矩阵
		{"2x2矩阵", [][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		// 边界：含负数与极值
		{"含负数与极值", [][]int{{1, -2, math.MinInt}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, -2}, {9, 6, math.MinInt}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			Rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("Rotate(%v) = %v, want %v", tt.matrix, matrix, tt.want)
			}
		})
	}
}

// TestRotateCopy 辅助数组解法结果应与原地解法一致
func TestRotateCopy(t *testing.T) {
	original := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	matrix := make([][]int, len(original))
	for i := range original {
		matrix[i] = make([]int, len(original[i]))
		copy(matrix[i], original[i])
	}
	RotateCopy(matrix)
	if !reflect.DeepEqual(matrix, want) {
		t.Errorf("RotateCopy(%v) = %v, want %v", original, matrix, want)
	}
}

func BenchmarkRotate(b *testing.B) {
	for _, n := range []int{10, 100, 1000} {
		b.Run(fmt.Sprintf("原地n=%d", n), func(b *testing.B) {
			matrix := makeSquareMatrix(n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Rotate(matrix)
			}
		})
		b.Run(fmt.Sprintf("复制n=%d", n), func(b *testing.B) {
			matrix := makeSquareMatrix(n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				RotateCopy(matrix)
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
