package maximumnumberofones

import "testing"

func TestMaximumNumberOfOnes(t *testing.T) {
	tests := []struct {
		name       string
		width      int
		height     int
		sideLength int
		maxOnes    int
		want       int
	}{
		// LeetCode 官方示例
		{"示例1: 3x3 边长2 最多1个1", 3, 3, 2, 1, 4},
		{"示例2: 3x3 边长2 最多2个1", 3, 3, 2, 2, 6},

		// 边界：sideLength 等于矩阵边长，唯一子矩阵即整个矩阵
		{"子矩阵即整个矩阵", 4, 4, 4, 5, 5},
		{"子矩阵即整个矩阵允许填满", 4, 4, 4, 16, 16},

		// 边界：sideLength 为 1，每个格子都是独立子矩阵
		{"边长1每个格子独立", 3, 3, 1, 1, 9},

		// 边界：最小矩阵
		{"1x1 矩阵", 1, 1, 1, 1, 1},

		// 典型场景：非正方形矩阵
		{"宽大于高", 5, 3, 2, 2, 10},
		{"高大于宽", 3, 5, 2, 2, 10},
		{"不能整除", 5, 5, 3, 2, 8},
		{"允许较多1", 4, 5, 2, 4, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumNumberOfOnes(tt.width, tt.height, tt.sideLength, tt.maxOnes); got != tt.want {
				t.Errorf("MaximumNumberOfOnes(%d, %d, %d, %d) = %d, want %d",
					tt.width, tt.height, tt.sideLength, tt.maxOnes, got, tt.want)
			}
		})
	}
}

func BenchmarkMaximumNumberOfOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximumNumberOfOnes(100, 100, 50, 2500)
	}
}
