package longestpalindromicsubstring

import "testing"

// 回文答案可能不唯一（如 "babad" 的 "bab"/"aba"），用「长度 + 确实是回文 + 确实是子串」属性校验
var longestPalindromeCases = []struct {
	name string
	s    string
	want string // 期望答案之一，用于确定最长长度
}{
	// LeetCode 官方示例
	{name: "示例1_babad", s: "babad", want: "bab"},
	{name: "示例2_cbbd", s: "cbbd", want: "bb"},

	// 边界：单个字符
	{name: "单字符a", s: "a", want: "a"},
	{name: "单字符数字1", s: "1", want: "1"},

	// 边界：两个字符
	{name: "两字符回文aa", s: "aa", want: "aa"},
	{name: "两字符非回文ab", s: "ab", want: "a"},

	// 整个字符串本身就是回文
	{name: "整体是奇数长度回文abcba", s: "abcba", want: "abcba"},
	{name: "整体是偶数长度回文abba", s: "abba", want: "abba"},

	// 最长回文在开头或结尾
	{name: "回文在开头_aab", s: "aab", want: "aa"},
	{name: "回文在结尾_baa", s: "baa", want: "aa"},

	// 偶数长度回文需要双中心扩展
	{name: "偶数长度回文abccba", s: "abccba", want: "abccba"},
	{name: "偶数长度回文嵌套xaabbaay", s: "xaabbaay", want: "aabbaa"},

	// 全部字符相同
	{name: "全部相同字符aaaaa", s: "aaaaa", want: "aaaaa"},

	// 无长度大于1的回文时返回单个字符
	{name: "完全不同的字符abcdef", s: "abcdef", want: "a"},

	// 数字与字母混合
	{name: "数字回文a121b", s: "a121b", want: "121"},
	{name: "字母数字混合回文ab123321x", s: "ab123321x", want: "123321"},

	// 长度恰为约束上界 1000
	{name: "1000个相同字符", s: repeatString("z", 1000), want: repeatString("z", 1000)},
	// 两侧为无回文的周期串，最长回文 "abba" 位于正中间
	{name: "1000个字符且回文居中", s: patternRepeat("abc", 498) + "abba" + patternRepeat("def", 498), want: "abba"},
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

// isSubstr 校验 sub 是否为 s 的子串
func isSubstr(s, sub string) bool {
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
	if len(got) != len(want) || !isPalindrome(got) || !isSubstr(s, got) {
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

// repeatString 构造由 ch 重复 n 次组成的字符串
func repeatString(ch string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = ch[0]
	}
	return string(b)
}

// patternRepeat 构造 pattern 循环重复、总长度为 n 的字符串
func patternRepeat(pattern string, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = pattern[i%len(pattern)]
	}
	return string(b)
}

var benchStrs = []struct {
	name string
	s    string
}{
	{"短串babad", "babad"},
	{"1000个相同字符", repeatString("a", 1000)},
	{"1000个字符无回文", patternRepeat("abc", 1000)},
}

func BenchmarkLongestPalindrome(b *testing.B) {
	for _, bm := range benchStrs {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindrome(bm.s)
			}
		})
	}
}

func BenchmarkLongestPalindromeAlternative(b *testing.B) {
	for _, bm := range benchStrs {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindromeAlternative(bm.s)
			}
		})
	}
}
