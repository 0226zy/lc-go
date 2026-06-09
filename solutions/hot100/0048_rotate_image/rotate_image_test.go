package rotateimage

import (
	"reflect"
	"testing"
)

func copyMatrix(m [][]int) [][]int {
	cp := make([][]int, len(m))
	for i := range m {
		cp[i] = make([]int, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name:   "1x1矩阵",
			input:  [][]int{{1}},
			expect: [][]int{{1}},
		},
		{
			name:   "2x2矩阵",
			input:  [][]int{{1, 2}, {3, 4}},
			expect: [][]int{{3, 1}, {4, 2}},
		},
		{
			name:   "官方示例1-3x3",
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "官方示例2-4x4",
			input:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expect: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			name:   "5x5矩阵",
			input:  [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			expect: [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
		{
			name:   "含负数",
			input:  [][]int{{-1, -2}, {-3, -4}},
			expect: [][]int{{-3, -1}, {-4, -2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := copyMatrix(tt.input)
			Rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.expect) {
				t.Errorf("Rotate() = %v, want %v", matrix, tt.expect)
			}
		})
	}
}

func TestRotateLayer(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name:   "1x1矩阵",
			input:  [][]int{{1}},
			expect: [][]int{{1}},
		},
		{
			name:   "2x2矩阵",
			input:  [][]int{{1, 2}, {3, 4}},
			expect: [][]int{{3, 1}, {4, 2}},
		},
		{
			name:   "官方示例1-3x3",
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "官方示例2-4x4",
			input:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expect: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			name:   "5x5矩阵",
			input:  [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			expect: [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
		{
			name:   "含负数",
			input:  [][]int{{-1, -2}, {-3, -4}},
			expect: [][]int{{-3, -1}, {-4, -2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := copyMatrix(tt.input)
			RotateLayer(matrix)
			if !reflect.DeepEqual(matrix, tt.expect) {
				t.Errorf("RotateLayer() = %v, want %v", matrix, tt.expect)
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	input := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := copyMatrix(input)
		b.StartTimer()
		Rotate(m)
	}
}

func BenchmarkRotateLayer(b *testing.B) {
	input := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := copyMatrix(input)
		b.StartTimer()
		RotateLayer(m)
	}
}
