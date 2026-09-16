package singlerowkeyboard

import "testing"

func TestCalculateTime(t *testing.T) {
	tests := []struct {
		name     string
		keyboard string
		word     string
		want     int
	}{
		// LeetCode 官方示例
		{"示例1: 标准键盘输入cba", "abcdefghijklmnopqrstuvwxyz", "cba", 4},
		{"示例2: 乱序键盘输入leetcode", "pqrstuvwxyzabcdefghijklmno", "leetcode", 73},

		// 边界：单字符
		{"单字符在下标0", "abcdefghijklmnopqrstuvwxyz", "a", 0},
		{"单字符在最远端", "abcdefghijklmnopqrstuvwxyz", "z", 25},

		// 边界：同一字符重复输入，手指不动
		{"重复同一字符", "abcdefghijklmnopqrstuvwxyz", "aaaa", 0},

		// 相邻字符
		{"相邻两字符", "abcdefghijklmnopqrstuvwxyz", "ab", 1},

		// 往返移动
		{"两端往返", "abcdefghijklmnopqrstuvwxyz", "azaz", 75},

		// 乱序键盘：'a' 在下标 25、'b' 在下标 24，耗时 25 + 1
		{"逆序键盘输入ab", "zyxwvutsrqponmlkjihgfedcba", "ab", 26},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTime(tt.keyboard, tt.word); got != tt.want {
				t.Errorf("CalculateTime(%q, %q) = %d, want %d", tt.keyboard, tt.word, got, tt.want)
			}
		})
	}
}

func BenchmarkCalculateTime(b *testing.B) {
	keyboard := "pqrstuvwxyzabcdefghijklmno"
	// 构造长度 10000 的长单词：leetcode 循环
	word := make([]byte, 0, 10000)
	for len(word) < 10000 {
		word = append(word, "leetcode"...)
	}
	word = word[:10000]

	b.Run("word长度10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CalculateTime(keyboard, string(word))
		}
	})
}
