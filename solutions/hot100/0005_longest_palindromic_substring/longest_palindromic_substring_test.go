package longestpalindromicsubstring

import (
	"strings"
	"testing"
)

func substringOf(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	return strings.Contains(s, sub)
}

func assertPalindrome(t *testing.T, p string) {
	t.Helper()
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		if p[i] != p[j] {
			t.Fatalf("%q is not a palindrome", p)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int
	}{
		{"示例1_babad", "babad", 3},
		{"示例2_cbbd", "cbbd", 2},
		{"示例3_单字符", "a", 1},
		{"示例4_ac", "ac", 1},
		{"全相同", "aaa", 3},
		{"偶长度回文", "abba", 4},
		{"奇长度回文", "racecar", 7},
		{"含数字", "ab2ba", 5},
		{"首尾相同长串", "abcdefedcba", 11},
		{"纯a_10", strings.Repeat("a", 10), 10},
		{"奇数中心cdc", "abcdc", 3},
		{"整串回文", "abcba", 5},
		{"多峰", "forgeeksskeegfor", 10}, // geeksskeeg
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindrome(tt.s)
			if len(got) != tt.wantLen {
				t.Errorf("LongestPalindrome(%q) len=%d want len=%d, got %q", tt.s, len(got), tt.wantLen, got)
			}
			if len(tt.s) > 0 && len(got) == 0 {
				t.Fatalf("LongestPalindrome(%q) empty", tt.s)
			}
			assertPalindrome(t, got)
			if !substringOf(tt.s, got) {
				t.Errorf("LongestPalindrome(%q) = %q, not a substring of s", tt.s, got)
			}
		})
	}
}

func TestLongestPalindromeDP(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int
	}{
		{"示例1_babad", "babad", 3},
		{"示例2_cbbd", "cbbd", 2},
		{"示例3_单字符", "a", 1},
		{"示例4_ac", "ac", 1},
		{"全相同", "aaa", 3},
		{"偶长度回文", "abba", 4},
		{"奇长度回文", "racecar", 7},
		{"含数字", "ab2ba", 5},
		{"首尾相同长串", "abcdefedcba", 11},
		{"纯a_10", strings.Repeat("a", 10), 10},
		{"奇数中心cdc", "abcdc", 3},
		{"整串回文", "abcba", 5},
		{"多峰", "forgeeksskeegfor", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestPalindromeDP(tt.s)
			if len(got) != tt.wantLen {
				t.Errorf("LongestPalindromeDP(%q) len=%d want len=%d, got %q", tt.s, len(got), tt.wantLen, got)
			}
			if len(tt.s) > 0 && len(got) == 0 {
				t.Fatalf("LongestPalindromeDP(%q) empty", tt.s)
			}
			assertPalindrome(t, got)
			if !substringOf(tt.s, got) {
				t.Errorf("LongestPalindromeDP(%q) = %q, not a substring of s", tt.s, got)
			}
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"len=10", strings.Repeat("abcdefghij", 1)},
		{"len=100", strings.Repeat("abcdefghij", 10)},
		{"len=500", strings.Repeat("ab", 250)},
	}
	for _, bm := range benchmarks {
		b.Run("中心扩展_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindrome(bm.s)
			}
		})
		b.Run("DP_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindromeDP(bm.s)
			}
		})
	}
}
