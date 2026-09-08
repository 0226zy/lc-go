package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: Hello World", "Hello World", 5},
		{"示例2: 飞航 moon", "   fly me   to   the moon  ", 4},
		{"示例3: joyboy", "luffy is still joyboy", 6},

		// 边界：单个单词
		{"单个单词", "word", 4},
		{"单个字母", "a", 1},

		// 边界：尾部带空格
		{"尾部带一个空格", "hello ", 5},
		{"尾部带多个空格", "hello    ", 5},

		// 边界：单词之间多个空格
		{"多个空格分隔", "a   b    c", 1},
		{"空格开头", "  hello", 5},

		// 边界：超长单词
		{"超长单词", "internationalization", 20},

		// 边界：含大写与符号
		{"大小写混合", "GoLang  Rocks ", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("LengthOfLastWord(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestLengthOfLastWordTrim(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1: Hello World", "Hello World", 5},
		{"示例2: 飞航 moon", "   fly me   to   the moon  ", 4},
		{"示例3: joyboy", "luffy is still joyboy", 6},
		{"无空格", "word", 4},
		{"全带空格", "   word   ", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWordTrim(tt.s); got != tt.want {
				t.Errorf("LengthOfLastWordTrim(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	s := "   fly me   to   the moon  "
	b.Run("双指针", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LengthOfLastWord(s)
		}
	})
	b.Run("标准库Trim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LengthOfLastWordTrim(s)
		}
	})
}
