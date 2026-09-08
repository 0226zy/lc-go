package wordpattern

import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		// LeetCode 官方示例
		{"示例1: abba 匹配 dog cat cat dog", "abba", "dog cat cat dog", true},
		{"示例2: abba 不匹配 dog cat cat fish", "abba", "dog cat cat fish", false},
		{"示例3: aaaa 不匹配 dog cat cat dog", "aaaa", "dog cat cat dog", false},

		// 边界：长度（单词数）不一致
		{"单词数少于字符数", "aaa", "aa aa aa aa", false},
		{"单词数多于字符数", "aaa", "aa aa", false},

		// 边界：空串
		{"pattern 为空且 s 为空", "", "", true},

		// 边界：单字符单单词
		{"单字符单单词", "a", "dog", true},

		// 边界：一对多冲突（同一字符映射到不同单词）
		{"同一字符映射到不同单词", "ab", "dog dog", false},

		// 边界：多对一冲突（不同字符映射到同一单词）
		{"不同字符映射到同一单词", "aa", "dog cat", false},

		// 边界：对称重复模式
		{"abab 模式", "abab", "dog cat dog cat", true},
		{"abcabc 模式", "abcabc", "one two three one two three", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordPattern(tt.pattern, tt.s); got != tt.want {
				t.Errorf("WordPattern(%q, %q) = %v, want %v", tt.pattern, tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkWordPattern(b *testing.B) {
	benchmarks := []struct {
		name    string
		pattern string
		s       string
	}{
		{"短串", "abba", "dog cat cat dog"},
		{"长串匹配", generatePattern("abc", 400), generateWords("one two three", 400)},
		{"长串不匹配", generatePattern("aaa", 400), generateWords("one two three", 400)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordPattern(bm.pattern, bm.s)
			}
		})
	}
}

// generatePattern 将 pattern 重复 n 次拼接成字符串
func generatePattern(pattern string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += pattern
	}
	return result
}

// generateWords 将 words（以空格分隔）重复 n 次并用空格连接
func generateWords(words string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		if result != "" {
			result += " "
		}
		result += words
	}
	return result
}
