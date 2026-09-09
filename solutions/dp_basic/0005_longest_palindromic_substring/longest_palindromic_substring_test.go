package longestpalindromicsubstring

import "testing"

// 回文答案可能不唯一（如 "babad" 的 "bab"/"aba"），用「长度 + 确实是回文 + 确实是子串」校验
var longestPalindromeCases = []struct {
	name string
	s    string
	want string // 期望答案之一；若为多个则用 wantLen + 回文校验
}{
	{name: "示例1_babad", s: "babad", want: "bab"},
	{name: "示例2_cbbd", s: "cbbd", want: "bb"},
	{name: "单字符", s: "a", want: "a"},
	{name: "两字符不相等", s: "ac", want: "a"},
	{name: "两字符相等", s: "bb", want: "bb"},
	{name: "整串回文", s: "abba", want: "abba"},
	{name: "全相同字符", s: "aaaa", want: "aaaa"},
	{name: "回文在末尾", s: "abcdcba", want: "abcdcba"},
	{name: "偶数长度回文", s: "aabbaa", want: "aabbaa"},
	{name: "混合字符", s: "abcbe", want: "bcb"},
}

// isPalindrome 校验字符串是否为回文
func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

// contains 校验 sub 是否为 s 的子串
func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func checkResult(t *testing.T, fn func(string) string, fnName, s, want string) {
	t.Helper()
	got := fn(s)
	if len(got) != len(want) || !isPalindrome(got) || !contains(s, got) {
		t.Errorf("%s(%q) = %q，期望长度为 %d 的回文子串", fnName, s, got, len(want))
	}
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range longestPalindromeCases {
		t.Run(tt.name, func(t *testing.T) {
			checkResult(t, LongestPalindrome, "LongestPalindrome", tt.s, tt.want)
		})
	}
}

func TestLongestPalindromeAlternative(t *testing.T) {
	for _, tt := range longestPalindromeCases {
		t.Run(tt.name, func(t *testing.T) {
			checkResult(t, LongestPalindromeAlternative, "LongestPalindromeAlternative", tt.s, tt.want)
		})
	}
}

var benchStr = func() string {
	b := make([]byte, 0, 1000)
	for i := 0; i < 1000; i++ {
		b = append(b, byte('a'+i%26))
	}
	return string(b)
}()

func BenchmarkLongestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestPalindrome(benchStr)
	}
}

func BenchmarkLongestPalindromeAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestPalindromeAlternative(benchStr)
	}
}
