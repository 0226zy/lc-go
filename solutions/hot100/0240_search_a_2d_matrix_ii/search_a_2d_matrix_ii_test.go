package searcha2dmatrixii

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "示例1-找到目标",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
		{
			name: "示例2-未找到目标",
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
		{
			name:   "空矩阵",
			matrix: [][]int{},
			target: 1,
			want:   false,
		},
		{
			name:   "空行矩阵",
			matrix: [][]int{{}},
			target: 1,
			want:   false,
		},
		{
			name:   "单元素-找到",
			matrix: [][]int{{5}},
			target: 5,
			want:   true,
		},
		{
			name:   "单元素-未找到",
			matrix: [][]int{{5}},
			target: 3,
			want:   false,
		},
		{
			name:   "单行矩阵-找到",
			matrix: [][]int{{1, 3, 5, 7, 9}},
			target: 7,
			want:   true,
		},
		{
			name:   "单行矩阵-未找到",
			matrix: [][]int{{1, 3, 5, 7, 9}},
			target: 6,
			want:   false,
		},
		{
			name:   "单列矩阵-找到",
			matrix: [][]int{{1}, {3}, {5}, {7}, {9}},
			target: 7,
			want:   true,
		},
		{
			name:   "单列矩阵-未找到",
			matrix: [][]int{{1}, {3}, {5}, {7}, {9}},
			target: 6,
			want:   false,
		},
		{
			name: "target为左上角",
			matrix: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			target: 1,
			want:   true,
		},
		{
			name: "target为右下角",
			matrix: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			target: 9,
			want:   true,
		},
		{
			name: "target小于最小值",
			matrix: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			target: 0,
			want:   false,
		},
		{
			name: "target大于最大值",
			matrix: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			target: 10,
			want:   false,
		},
		{
			name: "包含负数",
			matrix: [][]int{
				{-5, -3, -1},
				{-4, -2, 0},
				{-3, -1, 1},
			},
			target: -2,
			want:   true,
		},
		{
			name: "包含负数-未找到",
			matrix: [][]int{
				{-5, -3, -1},
				{-4, -2, 0},
				{-3, -1, 1},
			},
			target: -6,
			want:   false,
		},
		{
			name: "target在边界行",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			target: 6,
			want:   true,
		},
		{
			name: "target在边界列",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			target: 4,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 20
	for i := 0; i < b.N; i++ {
		_ = SearchMatrix(matrix, target)
	}
}
