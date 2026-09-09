package longestcommonsubsequence

import (
	"strings"
	"testing"
)

var lcsCases = []struct {
	name  string
	text1 string
	text2 string
	want  int
}{
	{name: "示例1：abcde与ace", text1: "abcde", text2: "ace", want: 3},
	{name: "示例2：完全相同", text1: "abc", text2: "abc", want: 3},
	{name: "示例3：无公共子序列", text1: "abc", text2: "def", want: 0},
	{name: "单字符相等", text1: "a", text2: "a", want: 1},
	{name: "单字符不等", text1: "a", text2: "b", want: 0},
	{name: "一个是另一个的子序列", text1: "abc", text2: "ahbgdc", want: 3},
	{name: "顺序不同", text1: "abcde", text2: "edcba", want: 1},
	{name: "重复字符", text1: "aabaa", text2: "ababa", want: 4},
	{name: "官方经典样例", text1: "oxcpqrsvwf", text2: "shmtulqrypy", want: 2},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tt := range lcsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("LongestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubsequenceOptimized(t *testing.T) {
	for _, tt := range lcsCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequenceOptimized(tt.text1, tt.text2); got != tt.want {
				t.Errorf("LongestCommonSubsequenceOptimized(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	text1 := strings.Repeat("abcde", 200)
	text2 := strings.Repeat("acebd", 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequence(text1, text2)
	}
}

func BenchmarkLongestCommonSubsequenceOptimized(b *testing.B) {
	text1 := strings.Repeat("abcde", 200)
	text2 := strings.Repeat("acebd", 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequenceOptimized(text1, text2)
	}
}
