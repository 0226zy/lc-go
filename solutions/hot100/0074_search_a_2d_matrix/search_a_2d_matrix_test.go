package searcha2dmatrix

import "testing"

func TestSearchMatrix_OfficialExamples(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "官方示例1-目标存在",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "官方示例2-目标不存在",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchMatrix_EdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "单元素矩阵-命中",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
		{
			name:   "单元素矩阵-未命中",
			matrix: [][]int{{1}},
			target: 2,
			want:   false,
		},
		{
			name:   "目标小于最小值",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 0,
			want:   false,
		},
		{
			name:   "目标大于最大值",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 100,
			want:   false,
		},
		{
			name:   "目标为行首元素",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 23,
			want:   true,
		},
		{
			name:   "目标为行尾元素",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 20,
			want:   true,
		},
		{
			name:   "含负数",
			matrix: [][]int{{-10, -5, 0, 5}, {10, 15, 20, 25}},
			target: -5,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchMatrix_AllSolutions(t *testing.T) {
	solutions := []func([][]int, int) bool{
		SearchMatrix,
		SearchMatrixRowCol,
	}

	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "官方示例1",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "官方示例2",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "单元素命中",
			matrix: [][]int{{5}},
			target: 5,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				got := fn(tt.matrix, tt.target)
				if got != tt.want {
					t.Errorf("solution[%d](%v, %d) = %v, want %v", i, tt.matrix, tt.target, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkSearchMatrix_Main(b *testing.B) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 34
	for i := 0; i < b.N; i++ {
		SearchMatrix(matrix, target)
	}
}

func BenchmarkSearchMatrix_RowCol(b *testing.B) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 34
	for i := 0; i < b.N; i++ {
		SearchMatrixRowCol(matrix, target)
	}
}
