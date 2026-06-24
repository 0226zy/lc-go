package largestrectangleinhistogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "示例1：基本用例",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "示例2：两个柱子",
			heights: []int{2, 4},
			want:    4,
		},
		{
			name:    "边界：单柱子",
			heights: []int{5},
			want:    5,
		},
		{
			name:    "全部递增",
			heights: []int{1, 2, 3, 4, 5},
			want:    9,
		},
		{
			name:    "全部递减",
			heights: []int{5, 4, 3, 2, 1},
			want:    9,
		},
		{
			name:    "全部等高",
			heights: []int{3, 3, 3, 3},
			want:    12,
		},
		{
			name:    "包含0",
			heights: []int{0, 2, 0},
			want:    2,
		},
		{
			name:    "全为0",
			heights: []int{0, 0, 0},
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestRectangleArea(tt.heights)
			if got != tt.want {
				t.Errorf("LargestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}

func BenchmarkLargestRectangleArea(b *testing.B) {
	heights := []int{2, 1, 5, 6, 2, 3}
	for i := 0; i < b.N; i++ {
		LargestRectangleArea(heights)
	}
}
