package nqueens

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// normalize 将每个棋盘拼成单个字符串后再整体排序，用于解的无序比较
func normalize(boards [][]string) []string {
	list := make([]string, len(boards))
	for i, board := range boards {
		list[i] = strings.Join(board, "\n")
	}
	sort.Strings(list)
	return list
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "示例1：n=4 两个解",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "示例2：n=1 只有一个解",
			n:    1,
			want: [][]string{{"Q"}},
		},
		{
			name: "边界：n=2 无解",
			n:    2,
			want: [][]string{},
		},
		{
			name: "边界：n=3 无解",
			n:    3,
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveNQueens(tt.n)
			gotNorm := normalize(got)
			wantNorm := normalize(tt.want)
			if len(gotNorm) == 0 && len(wantNorm) == 0 {
				return
			}
			if !reflect.DeepEqual(gotNorm, wantNorm) {
				t.Errorf("SolveNQueens(%d) = %v, 期望 %v", tt.n, got, tt.want)
			}
		})
	}

	// 解数量校验（经典结论）
	counts := []struct {
		name string
		n    int
		want int
	}{
		{"n=5 共 10 个解", 5, 10},
		{"n=6 共 4 个解", 6, 4},
		{"n=7 共 40 个解", 7, 40},
		{"n=8 经典八皇后共 92 个解", 8, 92},
		{"n=9 共 352 个解", 9, 352},
	}
	for _, tc := range counts {
		t.Run(tc.name, func(t *testing.T) {
			if got := len(SolveNQueens(tc.n)); got != tc.want {
				t.Errorf("SolveNQueens(%d) 解数量 = %d, 期望 %d", tc.n, got, tc.want)
			}
		})
	}

	// 解的合法性校验：每个解必须满足行、列、对角线互不攻击
	t.Run("所有解的合法性校验", func(t *testing.T) {
		for _, board := range SolveNQueens(5) {
			n := len(board)
			// 每行恰好一个 Q
			for r, row := range board {
				cnt := 0
				for _, ch := range row {
					if ch == 'Q' {
						cnt++
					}
				}
				if cnt != 1 {
					t.Errorf("第 %d 行应有 1 个 Q，实际 %d 个", r, cnt)
				}
			}
			// 每列恰好一个 Q，且对角线无冲突
			for c := 0; c < n; c++ {
				cnt := 0
				for r := 0; r < n; r++ {
					if board[r][c] == 'Q' {
						cnt++
					}
				}
				if cnt != 1 {
					t.Errorf("第 %d 列应有 1 个 Q，实际 %d 个", c, cnt)
				}
			}
			for r := 0; r < n; r++ {
				for c := 0; c < n; c++ {
					if board[r][c] != 'Q' {
						continue
					}
					for i, j := r+1, c+1; i < n && j < n; i, j = i+1, j+1 {
						if board[i][j] == 'Q' {
							t.Errorf("主对角线冲突：(%d,%d) 与 (%d,%d)", r, c, i, j)
						}
					}
					for i, j := r+1, c-1; i < n && j >= 0; i, j = i+1, j-1 {
						if board[i][j] == 'Q' {
							t.Errorf("副对角线冲突：(%d,%d) 与 (%d,%d)", r, c, i, j)
						}
					}
				}
			}
		}
	})
}

func BenchmarkSolveNQueens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveNQueens(8)
	}
}
