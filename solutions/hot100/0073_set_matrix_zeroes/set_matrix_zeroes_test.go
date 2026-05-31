package setmatrixzeroes

import (
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
		{"示例1: [[1,1,1],[1,0,1],[1,1,1]]", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"示例2: [[0,1,2,0],[3,4,5,2],[1,3,1,5]]", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},

		// 边界：全为非0
		{"全为非0", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"单元素非0", [][]int{{1}}, [][]int{{1}}},

		// 边界：全为0
		{"全为0", [][]int{{0, 0}, {0, 0}}, [][]int{{0, 0}, {0, 0}}},
		{"单元素为0", [][]int{{0}}, [][]int{{0}}},

		// 边界：0在第一行
		{"0在第一行", [][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{0, 0, 0}, {4, 0, 6}, {7, 0, 9}}},
		{"0在第一列", [][]int{{1, 2, 3}, {0, 5, 6}, {7, 8, 9}}, [][]int{{0, 2, 3}, {0, 0, 0}, {0, 8, 9}}},

		// 边界：多个0
		{"多个0", [][]int{{1, 0, 3}, {0, 5, 6}, {7, 8, 0}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		{"多个0在不同位置", [][]int{{1, 0}, {0, 4}}, [][]int{{0, 0}, {0, 0}}},

		// 边界：大矩阵
		{"大矩阵", [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {9, 10, 11, 12}}, [][]int{{1, 0, 3, 4}, {0, 0, 0, 0}, {9, 0, 11, 12}}},
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
	benchmarks := []struct {
		name   string
		matrix [][]int
	}{
		{"3x3", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
		{"3x4", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}},
		{"10x10", generateMatrix(10)},
		{"100x100", generateMatrix(100)},
		{"1000x1000", generateMatrix(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				matrix := make([][]int, len(bm.matrix))
				for i := range bm.matrix {
					matrix[i] = make([]int, len(bm.matrix[i]))
					copy(matrix[i], bm.matrix[i])
				}
				SetZeroes(matrix)
			}
		})
	}
}

// generateMatrix 生成大小为 n x n 的矩阵，主对角线为 1，其余为 0
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		matrix[i][i] = 1
	}
	return matrix
}
