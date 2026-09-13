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
			"示例1: 内部区域被翻转边界O保留",
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

		// 边界情况
		{"空矩阵", [][]byte{}, [][]byte{}},
		{
			"单行全O均在边界保留",
			[][]byte{{'O', 'O', 'O'}},
			[][]byte{{'O', 'O', 'O'}},
		},
		{
			"单列全O均在边界保留",
			[][]byte{{'O'}, {'O'}, {'O'}},
			[][]byte{{'O'}, {'O'}, {'O'}},
		},
		{
			"2x2全O全部在边界",
			[][]byte{{'O', 'O'}, {'O', 'O'}},
			[][]byte{{'O', 'O'}, {'O', 'O'}},
		},
		{
			"中心O被完全围绕",
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
		{
			"与边界连通的内部O保留",
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
		{
			"全X矩阵不变",
			[][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
			[][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
		},
		{
			"大块内部区域全部翻转",
			[][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
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

// makeAllO 生成 rows x cols 的全 'O' 棋盘，用于压测最坏递归深度
func makeAllO(rows, cols int) [][]byte {
	board := make([][]byte, rows)
	for i := range board {
		board[i] = make([]byte, cols)
		for j := range board[i] {
			board[i][j] = 'O'
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
		{"50x50全O", 50, 50},
		{"100x100全O", 100, 100},
		{"200x200全O", 200, 200},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Solve(makeAllO(bm.rows, bm.cols))
			}
		})
	}
}
