package surroundedregions

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		// LeetCode 官方示例
		{
			"示例1: 内部O被翻转",
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			"示例2: 单个X不变",
			[][]byte{{'X'}},
			[][]byte{{'X'}},
		},

		// 边界：空矩阵
		{"空矩阵", [][]byte{}, [][]byte{}},

		// 边界：单行/单列
		{"单行全O", [][]byte{{'O', 'O', 'O'}}, [][]byte{{'O', 'O', 'O'}}},
		{"单列全O", [][]byte{{'O'}, {'O'}, {'O'}}, [][]byte{{'O'}, {'O'}, {'O'}}},

		// 边界：2x2 全 O（都在边界上，全部保留）
		{
			"2x2全O",
			[][]byte{{'O', 'O'}, {'O', 'O'}},
			[][]byte{{'O', 'O'}, {'O', 'O'}},
		},

		// 边界：全部被围绕
		{
			"中心O被围绕",
			[][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},

		// 边界：角落的 O 保留
		{
			"角落O保留",
			[][]byte{
				{'O', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			[][]byte{
				{'O', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},

		// 与边界 O 相连的 O 保留
		{
			"边界连通区域保留",
			[][]byte{
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},

		// 全 X
		{
			"全X",
			[][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
			[][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solve(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("Solve() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}

// makeBoard 生成 rows x cols 的棋盘并按 checker 模式填充
// fillCenter 为 true 时中心区域填 'O'、边界填 'X'，否则全部填 'O'
func makeBoard(rows, cols int, centerX bool) [][]byte {
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			if centerX && i > 0 && i < rows-1 && j > 0 && j < cols-1 {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
	return board
}

func BenchmarkSolve(b *testing.B) {
	benchmarks := []struct {
		name string
		rows int
		cols int
	}{
		{"50x50", 50, 50},
		{"100x100", 100, 100},
		{"200x200", 200, 200},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				board := makeBoard(bm.rows, bm.cols, false)
				Solve(board)
			}
		})
	}
}
