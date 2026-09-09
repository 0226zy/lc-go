package wordbreak

import (
	"strings"
	"testing"
)

var wordBreakCases = []struct {
	name     string
	s        string
	wordDict []string
	want     bool
}{
	// LeetCode 官方示例
	{name: "示例1：两个单词拼接", s: "leetcode", wordDict: []string{"leet", "code"}, want: true},
	{name: "示例2：单词重复使用", s: "applepenapple", wordDict: []string{"apple", "pen"}, want: true},
	{name: "示例3：无法拼接", s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, want: false},

	// 边界：单字符
	{name: "单字符且在字典中", s: "a", wordDict: []string{"a"}, want: true},
	{name: "单字符不在字典中", s: "a", wordDict: []string{"b"}, want: false},

	// 边界：整个字符串恰好是一个字典单词
	{name: "整个字符串就是一个单词", s: "apple", wordDict: []string{"apple", "pen"}, want: true},

	// 边界：字典单词可重复使用
	{name: "单词重复使用拼出长串", s: "aaaaaaa", wordDict: []string{"aa", "aaa"}, want: true},
	{name: "只能重复仍拼不出", s: "aaaa", wordDict: []string{"aaa"}, want: false},
	{name: "长单词拼不出", s: "aaaaaaa", wordDict: []string{"aaaa"}, want: false},

	// 易错：贪心失败，需要枚举最后一段
	{name: "贪心选长词失败但长加短可拆", s: "cars", wordDict: []string{"car", "ca", "rs"}, want: true},
	{name: "多种切分路径", s: "abcd", wordDict: []string{"a", "abc", "b", "cd"}, want: true},
	{name: "看似可拆实则不可拆", s: "aaaaa", wordDict: []string{"aab", "baa", "ab"}, want: false},
	{name: "贪心选错会失败但DP正确", s: "ccbb", wordDict: []string{"cb", "bc"}, want: false},

	// 边界：字典单词比串长不影响结果
	{name: "字典单词比串长不影响结果", s: "ab", wordDict: []string{"abc", "ab"}, want: true},
	{name: "字典单词全部比串长", s: "ab", wordDict: []string{"abc", "abcd"}, want: false},

	// 压力场景：长度300的全a串（约束上限）
	{name: "压力：300个a可拆", s: strings.Repeat("a", 300), wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, want: true},
	{name: "压力：300个a加结尾b不可拆", s: strings.Repeat("a", 299) + "b", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, want: false},
}

func TestWordBreak(t *testing.T) {
	for _, tt := range wordBreakCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreak(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func BenchmarkWordBreak(b *testing.B) {
	s := "leetcodeleetcodeleetcodeleetcode"
	dict := []string{"leet", "code", "leetcode", "lee", "tc"}
	for i := 0; i < b.N; i++ {
		WordBreak(s, dict)
	}
}
