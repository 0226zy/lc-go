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
			name: "双源2分钟腐化",
			grid: [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}},
			want: 2,
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

func TestOrangesRottingBFS(t *testing.T) {
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
			name: "双源2分钟腐化",
			grid: [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}},
			want: 2,
		},
	}

	// 深拷贝辅助：BFS 原地修改 grid
	cloneGrid := func(src [][]int) [][]int {
		dst := make([][]int, len(src))
		for i := range src {
			dst[i] = make([]int, len(src[i]))
			copy(dst[i], src[i])
		}
		return dst
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrangesRottingBFS(cloneGrid(tt.grid))
			if got != tt.want {
				t.Errorf("OrangesRottingBFS(%v) = %d, want %d", tt.grid, got, tt.want)
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

func BenchmarkOrangesRottingBFS(b *testing.B) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	for i := 0; i < b.N; i++ {
		// BFS 原地修改，需要拷贝
		gridCopy := make([][]int, len(grid))
		for j := range grid {
			gridCopy[j] = make([]int, len(grid[j]))
			copy(gridCopy[j], grid[j])
		}
		OrangesRottingBFS(gridCopy)
	}
}
