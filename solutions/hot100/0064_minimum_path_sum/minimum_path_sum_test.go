package minimumpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "示例1：3x3网格",
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			name: "示例2：2x3网格",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
		{
			name: "边界：1x1网格",
			grid: [][]int{{5}},
			want: 5,
		},
		{
			name: "边界：单行",
			grid: [][]int{{1, 2, 3, 4}},
			want: 10,
		},
		{
			name: "边界：单列",
			grid: [][]int{{1}, {2}, {3}, {4}},
			want: 10,
		},
		{
			name: "正方形最小网格",
			grid: [][]int{{1, 2}, {1, 1}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinPathSum(tt.grid)
			if got != tt.want {
				t.Errorf("MinPathSum(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	for i := 0; i < b.N; i++ {
		MinPathSum(grid)
	}
}
