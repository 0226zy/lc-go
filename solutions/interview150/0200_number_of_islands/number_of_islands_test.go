package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 三座岛",
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			"示例2: 两座岛",
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},

		// 边界：空网格
		{"空网格", [][]byte{}, 0},
		{"全为水", [][]byte{{'0', '0'}, {'0', '0'}}, 0},
		{"全为陆地", [][]byte{{'1', '1'}, {'1', '1'}}, 1},
		{"单格陆地", [][]byte{{'1'}}, 1},
		{"单格水", [][]byte{{'0'}}, 0},

		// 边界：对角线不算相连
		{
			"对角不相连",
			[][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			2,
		},
		{
			"斜向复杂形状",
			[][]byte{
				{'1', '1', '0'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			3,
		},

		// 长条岛屿
		{
			"横条岛屿",
			[][]byte{{'1', '1', '1', '1', '1'}},
			1,
		},
		{
			"竖条岛屿",
			[][]byte{{'1'}, {'1'}, {'1'}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制网格，因为主解会修改网格
			grid := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]byte, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			if got := NumIslands(grid); got != tt.want {
				t.Errorf("NumIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNumIslandsUF(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			"示例1: 一座岛",
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			"示例2: 三座岛",
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{"空网格", [][]byte{}, 0},
		{"全为水", [][]byte{{'0', '0'}, {'0', '0'}}, 0},
		{"全为陆地", [][]byte{{'1', '1'}, {'1', '1'}}, 1},
		{"单格陆地", [][]byte{{'1'}}, 1},
		{
			"对角不相连",
			[][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslandsUF(tt.grid); got != tt.want {
				t.Errorf("NumIslandsUF() = %d, want %d", got, tt.want)
			}
		})
	}
}

// makeGrid 生成 rows x cols 的网格，全部填充 fill 字符
func makeGrid(rows, cols int, fill byte) [][]byte {
	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
		for j := range grid[i] {
			grid[i][j] = fill
		}
	}
	return grid
}

func BenchmarkNumIslands(b *testing.B) {
	benchmarks := []struct {
		name string
		rows int
		cols int
		fill byte
	}{
		{"10x10全水", 10, 10, '0'},
		{"50x50全陆地", 50, 50, '1'},
		{"100x100全陆地", 100, 100, '1'},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NumIslands(makeGrid(bm.rows, bm.cols, bm.fill))
			}
		})
	}
}

func BenchmarkNumIslandsUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := makeGrid(100, 100, '1')
		NumIslandsUF(grid)
	}
}
