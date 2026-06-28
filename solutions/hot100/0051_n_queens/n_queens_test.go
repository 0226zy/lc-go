package nqueens

import (
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantLen int
	}{
		{
			name:    "示例1：4皇后",
			n:       4,
			wantLen: 2,
		},
		{
			name:    "示例2：1皇后",
			n:       1,
			wantLen: 1,
		},
		{
			name:    "2皇后无解",
			n:       2,
			wantLen: 0,
		},
		{
			name:    "3皇后无解",
			n:       3,
			wantLen: 0,
		},
		{
			name:    "5皇后",
			n:       5,
			wantLen: 10,
		},
		{
			name:    "6皇后",
			n:       6,
			wantLen: 4,
		},
		{
			name:    "7皇后",
			n:       7,
			wantLen: 40,
		},
		{
			name:    "8皇后（经典）",
			n:       8,
			wantLen: 92,
		},
		{
			name:    "9皇后",
			n:       9,
			wantLen: 352,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveNQueens(tt.n)
			if got == nil {
				t.Skip("未实现")
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("SolveNQueens(%d) 应返回 %d 个解，实际返回 %d 个", tt.n, tt.wantLen, len(got))
			}
			// 验证每个解的有效性
			for _, board := range got {
				if len(board) != tt.n {
					t.Errorf("解的棋盘大小不正确: got %d, want %d", len(board), tt.n)
					continue
				}
				// 检查每行只有一个Q
				for i, row := range board {
					qCount := 0
					for _, ch := range row {
						if ch == 'Q' {
							qCount++
						}
					}
					if qCount != 1 {
						t.Errorf("第 %d 行应有 1 个 Q，实际有 %d 个", i, qCount)
					}
				}
				// 检查每列只有一个Q
				for col := 0; col < tt.n; col++ {
					qCount := 0
					for row := 0; row < tt.n; row++ {
						if board[row][col] == 'Q' {
							qCount++
						}
					}
					if qCount != 1 {
						t.Errorf("第 %d 列应有 1 个 Q，实际有 %d 个", col, qCount)
					}
				}
				// 检查对角线无冲突
				for r := 0; r < tt.n; r++ {
					for c := 0; c < tt.n; c++ {
						if board[r][c] != 'Q' {
							continue
						}
						// 右下对角线
						for i, j := r+1, c+1; i < tt.n && j < tt.n; i, j = i+1, j+1 {
							if board[i][j] == 'Q' {
								t.Errorf("对角线冲突: (%d,%d) 和 (%d,%d)", r, c, i, j)
							}
						}
						// 左下对角线
						for i, j := r+1, c-1; i < tt.n && j >= 0; i, j = i+1, j-1 {
							if board[i][j] == 'Q' {
								t.Errorf("反对角线冲突: (%d,%d) 和 (%d,%d)", r, c, i, j)
							}
						}
					}
				}
			}
		})
	}
}

func BenchmarkSolveNQueens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveNQueens(8)
	}
}
