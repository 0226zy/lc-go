package reversewordsinastringii

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: the sky is blue", "the sky is blue", "blue is sky the"},
		{"示例2: 单字符", "a", "a"},

		// 边界：单单词
		{"单个长单词", "hello", "hello"},

		// 典型场景
		{"两个单词", "hello world", "world hello"},
		{"首尾单词长度不同", "a bcd", "bcd a"},
		{"含数字与大小写", "Go is NO1", "NO1 is Go"},
		{"多个单词", "one two three four", "four three two one"},
		{"重复单词", "ab ab ab", "ab ab ab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := []byte(tt.s)
			ReverseWords(s)
			if got := string(s); got != tt.want {
				t.Errorf("ReverseWords(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkReverseWords(b *testing.B) {
	// 构造一段长文本：10000 个单词
	word := []byte("leetcode ")
	long := make([]byte, 0, 90000)
	for len(long)+len(word) <= 90000 {
		long = append(long, word...)
	}

	b.Run("90000字符", func(b *testing.B) {
		s := make([]byte, len(long))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(s, long) // 每轮还原输入，保证原地修改不影响计时公平性
			b.StartTimer()
			ReverseWords(s)
		}
	})
}
