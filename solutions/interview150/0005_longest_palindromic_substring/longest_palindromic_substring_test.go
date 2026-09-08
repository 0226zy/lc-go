package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string // 存在多个最长答案时，任意一个均可接受
	}{
		// LeetCode 官方示例
		{"示例1: babad的答案是bab或aba", "babad", []string{"bab", "aba"}},
		{"示例2: cbbd的答案是bb", "cbbd", []string{"bb"}},

		// 边界：单个字符
		{"单字符a", "a", []string{"a"}},
		{"单字符数字1", "1", []string{"1"}},

		// 边界：两个字符
		{"两字符回文aa", "aa", []string{"aa"}},
		{"两字符非回文ab", "ab", []string{"a", "b"}},

		// 整个字符串本身就是回文
		{"整体是奇数长度回文abcba", "abcba", []string{"abcba"}},
		{"整体是偶数长度回文abba", "abba", []string{"abba"}},

		// 最长回文在开头或结尾
		{"回文在开头: aab", "aab", []string{"aa"}},
		{"回文在结尾: baa", "baa", []string{"aa"}},

		// 偶数长度回文需要双中心扩展
		{"偶数长度回文abccba", "abccba", []string{"abccba"}},
		{"偶数长度回文嵌套xaabbaay", "xaabbaay", []string{"aabbaa"}},

		// 全部字符相同
		{"全部相同字符aaaaa", "aaaaa", []string{"aaaaa"}},

		// 无长度大于1的回文时返回单个字符
		{"完全不同的字符abcdef", "abcdef", []string{"a", "b", "c", "d", "e", "f"}},

		// 数字与字母混合
		{"数字回文a121b", "a121b", []string{"121"}},
		{"字母数字混合回文ab123321x", "ab123321x", []string{"123321"}},

		// 长度恰为约束上界 1000
		{"1000个相同字符", repeatString("z", 1000), []string{repeatString("z", 1000)}},
		// 两侧为无回文的周期串，最长回文 "abba" 位于正中间
		{"1000个字符且回文居中", patternRepeat("abc", 498) + "abba" + patternRepeat("def", 498), []string{"abba"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindrome(tt.s)
			if !contains(tt.want, got) {
				t.Errorf("LongestPalindrome(%q) = %q, 期望是 %v 中的任意一个", tt.s, got, tt.want)
			}
		})
	}
}

// contains 判断 got 是否在可接受的答案集合中
func contains(candidates []string, got string) bool {
	for _, c := range candidates {
		if c == got {
			return true
		}
	}
	return false
}

// repeatString 构造由 ch 重复 n 次组成的字符串
func repeatString(ch string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = ch[0]
	}
	return string(b)
}

func BenchmarkLongestPalindrome(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"短串babad", "babad"},
		{"1000个相同字符", repeatString("a", 1000)},
		{"1000个字符无回文", noPalindromeString(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindrome(bm.s)
			}
		})
	}
}

// noPalindromeString 构造长度为 n 且无长度大于 1 回文的字符串（周期 abc）
func noPalindromeString(n int) string {
	return patternRepeat("abc", n)
}

// patternRepeat 构造 pattern 循环重复、总长度为 n 的字符串
func patternRepeat(pattern string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = pattern[i%len(pattern)]
	}
	return string(b)
}
