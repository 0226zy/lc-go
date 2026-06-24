package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "示例1：1个岛屿",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "示例2：3个岛屿",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "边界：全为水",
			grid: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			want: 0,
		},
		{
			name: "边界：全为陆地",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "边界：1x1陆地",
			grid: [][]byte{{'1'}},
			want: 1,
		},
		{
			name: "边界：1x1水",
			grid: [][]byte{{'0'}},
			want: 0,
		},
		{
			name: "对角线不算连接",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumIslands(tt.grid)
			if got != tt.want {
				t.Errorf("NumIslands(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestNumIslandsBFS(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "示例1：1个岛屿",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "示例2：3个岛屿",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "边界：全为水",
			grid: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			want: 0,
		},
		{
			name: "边界：全为陆地",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "边界：1x1陆地",
			grid: [][]byte{{'1'}},
			want: 1,
		},
		{
			name: "边界：1x1水",
			grid: [][]byte{{'0'}},
			want: 0,
		},
		{
			name: "对角线不算连接",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// BFS 原地修改 grid，需要深拷贝
			gridCopy := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]byte, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}
			got := NumIslandsBFS(gridCopy)
			if got != tt.want {
				t.Errorf("NumIslandsBFS(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func BenchmarkNumIslands(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		NumIslands(grid)
	}
}

func BenchmarkNumIslandsBFS(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		// BFS 原地修改，需要拷贝
		gridCopy := make([][]byte, len(grid))
		for j := range grid {
			gridCopy[j] = make([]byte, len(grid[j]))
			copy(gridCopy[j], grid[j])
		}
		NumIslandsBFS(gridCopy)
	}
}
