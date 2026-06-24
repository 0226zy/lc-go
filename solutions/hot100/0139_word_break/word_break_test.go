package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{"官方示例1", "leetcode", []string{"leet", "code"}, true},
		{"官方示例2", "applepenapple", []string{"apple", "pen"}, true},
		{"官方示例3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"单字符匹配", "a", []string{"a"}, true},
		{"单字符不匹配", "a", []string{"b"}, false},
		{"空字典无法匹配", "a", []string{}, false},
		{"完整重复使用", "aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"长串重复", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, true},
		{"无法拆分的边界", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa"}, false},
		{"单词等于整个字符串", "abcd", []string{"abcd"}, true},
		{"需要精准组合", "cars", []string{"car", "ca", "rs"}, true},
		{"字典有多余单词", "ab", []string{"a", "b", "c", "d"}, true},
		{"前缀不匹配", "deadbeef", []string{"dead", "bee"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreak(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func TestWordBreakBFS(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{"官方示例1", "leetcode", []string{"leet", "code"}, true},
		{"官方示例2", "applepenapple", []string{"apple", "pen"}, true},
		{"官方示例3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"单字符匹配", "a", []string{"a"}, true},
		{"单字符不匹配", "a", []string{"b"}, false},
		{"完整重复使用", "aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"需要精准组合", "cars", []string{"car", "ca", "rs"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreakBFS(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreakBFS(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func TestWordBreakMemo(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{"官方示例1", "leetcode", []string{"leet", "code"}, true},
		{"官方示例2", "applepenapple", []string{"apple", "pen"}, true},
		{"官方示例3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"单字符匹配", "a", []string{"a"}, true},
		{"无法拆分的边界", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa"}, false},
		{"需要精准组合", "cars", []string{"car", "ca", "rs"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreakMemo(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreakMemo(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func TestWordBreakTrie(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{"官方示例1", "leetcode", []string{"leet", "code"}, true},
		{"官方示例2", "applepenapple", []string{"apple", "pen"}, true},
		{"官方示例3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"单字符匹配", "a", []string{"a"}, true},
		{"完整重复使用", "aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"无法拆分的边界", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa"}, false},
		{"需要精准组合", "cars", []string{"car", "ca", "rs"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreakTrie(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreakTrie(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

// TestConsistency 验证四种实现结果一致
func TestConsistency(t *testing.T) {
	cases := []struct {
		s        string
		wordDict []string
	}{
		{"leetcode", []string{"leet", "code"}},
		{"applepenapple", []string{"apple", "pen"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa"}},
		{"cars", []string{"car", "ca", "rs"}},
		{"aaaaaaa", []string{"aaaa", "aaa"}},
		{"a", []string{"a"}},
		{"a", []string{"b"}},
	}
	for _, c := range cases {
		a := WordBreak(c.s, c.wordDict)
		b := WordBreakBFS(c.s, c.wordDict)
		cc := WordBreakMemo(c.s, c.wordDict)
		d := WordBreakTrie(c.s, c.wordDict)
		if a != b || b != cc || cc != d {
			t.Errorf("四种实现结果不一致: s=%q, DP=%v, BFS=%v, Memo=%v, Trie=%v",
				c.s, a, b, cc, d)
		}
	}
}

func BenchmarkWordBreak(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i < b.N; i++ {
		WordBreak(s, dict)
	}
}

func BenchmarkWordBreakTrie(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i < b.N; i++ {
		WordBreakTrie(s, dict)
	}
}
