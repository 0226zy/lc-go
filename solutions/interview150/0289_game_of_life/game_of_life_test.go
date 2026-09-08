package gameoflife

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 4x3棋盘", [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},
		{"示例2: 2x2棋盘", [][]int{{1, 1}, {1, 0}}, [][]int{{1, 1}, {1, 1}}},

		// 边界：全死棋盘保持不变
		{"全死棋盘", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		// 边界：全活棋盘（角活2死，边活3活，中心活4死）
		{"全活3x3", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		// 边界：单活细胞因孤立死亡
		{"单活细胞", [][]int{{1}}, [][]int{{0}}},
		// 边界：单死细胞保持死亡（无邻居）
		{"单死细胞", [][]int{{0}}, [][]int{{0}}},
		// 边界：三个活细胞排成一列，中间复活两端
		{"竖直三连", [][]int{{1}, {1}, {1}}, [][]int{{0}, {1}, {0}}},
		// 边界：L 形三角（每个活细胞都有 2 个活邻居，稳态不变）
		{"L形稳态", [][]int{{1, 1}, {1, 0}}, [][]int{{1, 1}, {1, 1}}},
		// 边界：经典方块振荡器（2x2 全活为静物）
		{"方块静物", [][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			[][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := make([][]int, len(tt.board))
			for i := range tt.board {
				board[i] = make([]int, len(tt.board[i]))
				copy(board[i], tt.board[i])
			}
			GameOfLife(board)
			if !reflect.DeepEqual(board, tt.want) {
				t.Errorf("GameOfLife(%v) = %v, want %v", tt.board, board, tt.want)
			}
		})
	}
}

// TestGameOfLifeCopy 辅助数组解法结果应与原地解法一致
func TestGameOfLifeCopy(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	want := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}

	copyBoard := make([][]int, len(board))
	for i := range board {
		copyBoard[i] = make([]int, len(board[i]))
		copy(copyBoard[i], board[i])
	}
	GameOfLifeCopy(copyBoard)
	if !reflect.DeepEqual(copyBoard, want) {
		t.Errorf("GameOfLifeCopy(%v) = %v, want %v", board, copyBoard, want)
	}
}

func BenchmarkGameOfLife(b *testing.B) {
	for _, size := range []int{25, 100, 400} {
		b.Run(fmt.Sprintf("原地n=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GameOfLife(makeRandomBoard(size))
			}
		})
		b.Run(fmt.Sprintf("复制n=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GameOfLifeCopy(makeRandomBoard(size))
			}
		})
	}
}

// makeRandomBoard 生成 size x size 的确定性伪随机棋盘（活细胞约占一半）
func makeRandomBoard(size int) [][]int {
	board := make([][]int, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
		for j := 0; j < size; j++ {
			// 用简单的哈希函数生成确定性伪随机值
			board[i][j] = (i*31 + j*17 + i*j) % 2
		}
	}
	return board
}
