package wordsearch

import "testing"

// makeBoard 把字符串切片转成独立的字符网格，避免测试用例之间共享底层数组
func makeBoard(rows []string) [][]byte {
	board := make([][]byte, len(rows))
	for i, row := range rows {
		board[i] = []byte(row)
	}
	return board
}

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board []string
		word  string
		want  bool
	}{
		// LeetCode 官方示例
		{"示例一：ABCCED 在网格中存在",
			[]string{"ABCE", "SFCS", "ADEE"}, "ABCCED", true},
		{"示例二：SEE 在网格中存在",
			[]string{"ABCE", "SFCS", "ADEE"}, "SEE", true},
		{"示例三：ABCB 因格子不可复用而不存在",
			[]string{"ABCE", "SFCS", "ADEE"}, "ABCB", false},

		// 边界情况
		{"单词长度超过格子总数", []string{"AB", "CD"}, "ABCDE", false},
		{"单词首字符不在网格中", []string{"AB", "CD"}, "X", false},
		{"单字符单词且命中", []string{"AB", "CD"}, "A", true},
		{"单字符单词且未命中", []string{"AB", "CD"}, "Z", false},
		{"需要回头走但会被已访问标记挡住", []string{"AB", "CD"}, "ABA", false},
		{"蛇形路径恰好走满整个网格", []string{"ABC", "DEF", "GHI"}, "ABCFEDGHI", true},
		{"某字符在单词中出现次数超过网格总数（预剪枝命中）", []string{"AB", "CD"}, "ABDDC", false},
		{"单行网格横向匹配", []string{"ABCDE"}, "EDC", true},
		{"单列网格纵向匹配", []string{"A", "B", "C", "D"}, "ABC", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(makeBoard(tt.board), tt.word); got != tt.want {
				t.Errorf("Exist(%v, %q) = %v, 期望 %v", tt.board, tt.word, got, tt.want)
			}
		})
	}
}

func BenchmarkExist(b *testing.B) {
	rows := []string{
		"ABCEHJIG", "SFCSLOPQ", "ADEEMNOE", "ADIDEJFM", "VCDIFGGT",
	}
	word := "SLHECCEIDEJFGWTHE"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Exist(makeBoard(rows), word)
	}
}
