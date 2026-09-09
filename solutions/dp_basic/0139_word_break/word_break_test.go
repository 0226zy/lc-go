package wordbreak

import "testing"

var wordBreakCases = []struct {
	name     string
	s        string
	wordDict []string
	want     bool
}{
	{name: "示例1：两个单词拼接", s: "leetcode", wordDict: []string{"leet", "code"}, want: true},
	{name: "示例2：单词重复使用", s: "applepenapple", wordDict: []string{"apple", "pen"}, want: true},
	{name: "示例3：无法拼接", s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, want: false},
	{name: "整个字符串就是一个单词", s: "apple", wordDict: []string{"apple"}, want: true},
	{name: "单字符重复拼接", s: "aaaaaaa", wordDict: []string{"a"}, want: true},
	{name: "长单词拼不出", s: "aaaaaaa", wordDict: []string{"aaaa"}, want: false},
	{name: "3a加4a正好拼7个a", s: "aaaaaaa", wordDict: []string{"aaa", "aaaa"}, want: true},
	{name: "多种切分路径", s: "abcd", wordDict: []string{"a", "abc", "b", "cd"}, want: true},
	{name: "字典单词比剩余串长", s: "cars", wordDict: []string{"car", "ca", "rs"}, want: true},
	{name: "贪心选错会失败但DP正确", s: "ccbb", wordDict: []string{"cb", "bc"}, want: false},
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
