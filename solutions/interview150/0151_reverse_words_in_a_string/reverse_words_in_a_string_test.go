package reversewordsinastring

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: the sky is blue", "the sky is blue", "blue is sky the"},
		{"示例2: 多空格", "  hello world  ", "world hello"},
		{"示例3: 单词间多空格", "a good   example", "example good a"},

		// 边界：单个单词
		{"单个单词", "word", "word"},
		{"单个字母", "a", "a"},

		// 边界：首尾带空格
		{"前导空格", "   hello", "hello"},
		{"尾随空格", "hello   ", "hello"},
		{"首尾都有空格", "  hello world  ", "world hello"},

		// 边界：单词间多空格
		{"单词间多个空格", "a  b   c", "c b a"},
		{"全空格", "     ", ""},

		// 边界：两个单词
		{"两个单词", "hello world", "world hello"},

		// 边界：含大写与数字
		{"大小写数字混合", "Go 1.22 Rocks", "Rocks 1.22 Go"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.s); got != tt.want {
				t.Errorf("ReverseWords(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkReverseWords(b *testing.B) {
	s := "  the   quick  brown  fox  jumps  over  the  lazy  dog  "
	b.Run("切片重建", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ReverseWords(s)
		}
	})
}
