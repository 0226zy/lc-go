package longestpalindromicsubsequence

import "testing"

var longestPalindromeSubseqCases = []struct {
	name string
	s    string
	want int
}{
	{name: "官方示例1：bbbab", s: "bbbab", want: 4},
	{name: "官方示例2：cbbd", s: "cbbd", want: 2},
	{name: "单个字符", s: "a", want: 1},
	{name: "两个不同字符", s: "ab", want: 1},
	{name: "两个相同字符", s: "aa", want: 2},
	{name: "整个串都是同一字符", s: "aaaa", want: 4},
	{name: "互不相同的字符", s: "abcdef", want: 1},
	{name: "整个串本身就是回文", s: "racecar", want: 7},
	{name: "回文分散在串中", s: "character", want: 5},
	{name: "回文不相邻", s: "agbdba", want: 5},
}

func TestLongestPalindromeSubseq(t *testing.T) {
	for _, tt := range longestPalindromeSubseqCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeSubseq(tt.s); got != tt.want {
				t.Errorf("LongestPalindromeSubseq(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeSubseqOptimized(t *testing.T) {
	for _, tt := range longestPalindromeSubseqCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeSubseqOptimized(tt.s); got != tt.want {
				t.Errorf("LongestPalindromeSubseqOptimized(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestPalindromeSubseq(b *testing.B) {
	s := "bbbabagbdba"
	for i := 0; i < b.N; i++ {
		LongestPalindromeSubseq(s)
	}
}

func BenchmarkLongestPalindromeSubseqOptimized(b *testing.B) {
	s := "bbbabagbdba"
	for i := 0; i < b.N; i++ {
		LongestPalindromeSubseqOptimized(s)
	}
}
