package rottingoranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "示例1：4分钟全部腐烂",
			grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			want: 4,
		},
		{
			name: "示例2：无法全部腐烂",
			grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			want: -1,
		},
		{
			name: "示例3：初始无新鲜橘子",
			grid: [][]int{{0, 2}},
			want: 0,
		},
		{
			name: "边界：全为空",
			grid: [][]int{{0}},
			want: 0,
		},
		{
			name: "全为新鲜橘子",
			grid: [][]int{{1}},
			want: -1,
		},
		{
			name: "单格腐烂橘子",
			grid: [][]int{{2}},
			want: 0,
		},
		{
			name: "1分钟腐化",
			grid: [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrangesRotting(tt.grid)
			if got != tt.want {
				t.Errorf("OrangesRotting(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func BenchmarkOrangesRotting(b *testing.B) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	for i := 0; i < b.N; i++ {
		OrangesRotting(grid)
	}
}
