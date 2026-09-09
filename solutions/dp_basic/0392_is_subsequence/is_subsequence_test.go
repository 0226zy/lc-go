package issubsequence

import "testing"

var isSubsequenceCases = []struct {
	name string
	s    string
	t    string
	want bool
}{
	{name: "官方示例1：abc是ahbgdc的子序列", s: "abc", t: "ahbgdc", want: true},
	{name: "官方示例2：axc不是ahbgdc的子序列", s: "axc", t: "ahbgdc", want: false},
	{name: "s为空串", s: "", t: "ahbgdc", want: true},
	{name: "t为空串且s非空", s: "a", t: "", want: false},
	{name: "s与t均为空串", s: "", t: "", want: true},
	{name: "s与t完全相同", s: "abc", t: "abc", want: true},
	{name: "s比t长", s: "abcd", t: "abc", want: false},
	{name: "单字符匹配", s: "a", t: "ba", want: true},
	{name: "顺序不对不匹配", s: "aec", t: "abcde", want: false},
	{name: "重复字符", s: "aaaa", t: "aa", want: false},
}

func TestIsSubsequence(t *testing.T) {
	for _, tt := range isSubsequenceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.s, tt.t); got != tt.want {
				t.Errorf("IsSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestIsSubsequenceAlternative(t *testing.T) {
	for _, tt := range isSubsequenceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequenceAlternative(tt.s, tt.t); got != tt.want {
				t.Errorf("IsSubsequenceAlternative(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkIsSubsequence(b *testing.B) {
	s, str := "abc", "ahbgdc"
	for i := 0; i < b.N; i++ {
		IsSubsequence(s, str)
	}
}

func BenchmarkIsSubsequenceAlternative(b *testing.B) {
	s, str := "abc", "ahbgdc"
	for i := 0; i < b.N; i++ {
		IsSubsequenceAlternative(s, str)
	}
}
