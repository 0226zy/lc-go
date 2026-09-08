package wordsearch

import "testing"

// buildBoard 把字符串切片转换为字符网格的拷贝（避免用例间共享底层数组）
func buildBoard(rows []string) [][]byte {
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
		{"示例1: ABCCED 存在",
			[]string{"ABCE", "SFCS", "ADEE"}, "ABCCED", true},
		{"示例2: SEE 存在",
			[]string{"ABCE", "SFCS", "ADEE"}, "SEE", true},
		{"示例3: ABCB 不存在（不能重复使用格子）",
			[]string{"ABCE", "SFCS", "ADEE"}, "ABCB", false},

		// 边界：单词比网格总格子数长
		{"单词过长", []string{"AB", "CD"}, "ABCDE", false},
		// 边界：单词首字符不在网格中
		{"首字符不存在", []string{"AB", "CD"}, "X", false},
		// 边界：单格单词
		{"单格单词存在", []string{"AB", "CD"}, "A", true},
		{"单格单词不存在", []string{"AB", "CD"}, "Z", false},
		// 边界：需要折返但允许另起路径
		{"绕路匹配", []string{"ABCE", "SFES", "ADEE"}, "ABCESEEEFS", true},
		// 边界：重复字符陷阱（ABCB 类）
		{"ABA 不匹配", []string{"AB", "CD"}, "ABA", false},
		// 边界：蛇形路径恰好覆盖全网格
		{"蛇形覆盖全网格", []string{"ABC", "DEF", "GHI"}, "ABCFEDGHI", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(buildBoard(tt.board), tt.word); got != tt.want {
				t.Errorf("Exist(%v, %q) = %v, want %v", tt.board, tt.word, got, tt.want)
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
		Exist(buildBoard(rows), word)
	}
}
