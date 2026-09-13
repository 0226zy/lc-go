package numberofislands

import "testing"

// copyGrid 复制网格，因为淹岛法会修改输入，避免用例之间互相污染
func copyGrid(grid [][]byte) [][]byte {
	copied := make([][]byte, len(grid))
	for i := range grid {
		copied[i] = make([]byte, len(grid[i]))
		copy(copied[i], grid[i])
	}
	return copied
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		// LeetCode 官方示例
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

		// 边界情况
		{"空网格", [][]byte{}, 0},
		{"单格水", [][]byte{{'0'}}, 0},
		{"单格陆地", [][]byte{{'1'}}, 1},
		{"全为水", [][]byte{{'0', '0'}, {'0', '0'}}, 0},
		{"全为陆地", [][]byte{{'1', '1'}, {'1', '1'}}, 1},
		{"横条岛屿", [][]byte{{'1', '1', '1', '1', '1'}}, 1},
		{"竖条岛屿", [][]byte{{'1'}, {'1'}, {'1'}}, 1},

		// 对角线不相连
		{
			"对角不相连",
			[][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			2,
		},
		{
			"复杂交错形状",
			[][]byte{
				{'1', '1', '0'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			3,
		},
		{
			"蛇形单岛",
			[][]byte{
				{'1', '1', '1'},
				{'0', '0', '1'},
				{'1', '1', '1'},
			},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands(copyGrid(tt.grid)); got != tt.want {
				t.Errorf("NumIslands() = %d, want %d", got, tt.want)
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
		{"300x300全陆地", 300, 300, '1'},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NumIslands(makeGrid(bm.rows, bm.cols, bm.fill))
			}
		})
	}
}
