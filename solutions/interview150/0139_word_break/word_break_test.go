package wordbreak

import (
	"strings"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		// LeetCode 官方示例
		{"示例1: leetcode可由leet+code拆分", "leetcode", []string{"leet", "code"}, true},
		{"示例2: applepenapple可重复使用apple", "applepenapple", []string{"apple", "pen"}, true},
		{"示例3: catsandog无法拆分", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},

		// 边界：单字符
		{"单字符且在字典中", "a", []string{"a"}, true},
		{"单字符不在字典中", "a", []string{"b"}, false},

		// 边界：整个字符串恰好是一个字典单词
		{"整个串就是一个单词", "apple", []string{"apple", "pen"}, true},

		// 边界：字典单词可重复使用
		{"单词重复使用拼出长串", "aaaaaaa", []string{"aa", "aaa"}, true},
		{"只能重复仍拼不出", "aaaa", []string{"aaa"}, false},

		// 易错：贪心失败，需要回溯式枚举最后一段
		{"贪心选长词失败但长+短可拆", "cars", []string{"car", "ca", "rs"}, true},
		{"前缀可拆但整体不可拆", "abcd", []string{"ab", "abc", "cd", "a"}, true},
		{"看似可拆实则不可拆", "aaaaa", []string{"aab", "baa", "ab"}, false},

		// 边界：长度不超过最长单词的剪枝场景
		{"字典单词比串长不影响结果", "ab", []string{"abc", "ab"}, true},
		{"字典单词全部比串长", "ab", []string{"abc", "abcd"}, false},

		// 压力场景：长度300的全a串，字典含各种长度的a
		{"压力: 300个a且字典为a到a^20", strings.Repeat("a", 300), []string{
			"a", "aa", "aaa", "aaaa", "aaaaa",
		}, true},
		{"压力: 300个a加结尾b不可拆", strings.Repeat("a", 299) + "b", []string{
			"a", "aa", "aaa", "aaaa", "aaaaa",
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreak(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func BenchmarkWordBreak(b *testing.B) {
	benchmarks := []struct {
		name     string
		s        string
		wordDict []string
	}{
		{"官方示例leetcode", "leetcode", []string{"leet", "code"}},
		{"官方不可拆catsandog", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
		{"300个a可拆", strings.Repeat("a", 300), []string{"a", "aa", "aaa", "aaaa", "aaaaa"}},
		{"300字符不可拆", strings.Repeat("a", 299) + "b", []string{"a", "aa", "aaa", "aaaa", "aaaaa"}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordBreak(bm.s, bm.wordDict)
			}
		})
	}
}
