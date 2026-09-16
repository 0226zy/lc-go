package addboldtaginstring

import (
	"strings"
	"testing"
)

func TestAddBoldTag(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  string
	}{
		// LeetCode 官方示例
		{"示例1: 两段独立加粗", "abcxyz123", []string{"abc", "123"}, "<b>abc</b>xyz<b>123</b>"},
		{"示例2: 相邻区间合并", "aaabbcc", []string{"aaa", "aab", "bc"}, "<b>aaabbc</b>c"},

		// 边界：words 为空，原样返回
		{"空单词表", "abc", []string{}, "abc"},

		// 边界：没有任何匹配
		{"无匹配", "abcdef", []string{"xyz"}, "abcdef"},

		// 边界：整串加粗
		{"整串匹配", "abc", []string{"abc"}, "<b>abc</b>"},

		// 典型场景：重叠区间合并（"aa" 在 "aaaa" 中出现 3 次，区间连成一片）
		{"重叠出现合并", "aaaa", []string{"aa"}, "<b>aaaa</b>"},

		// 典型场景：同一单词出现两次且首尾相接，区间合并
		{"重复单词相邻合并", "abcabc", []string{"abc"}, "<b>abcabc</b>"},

		// 典型场景：同一单词多次出现且不相邻
		{"重复单词多处匹配", "abcXabc", []string{"abc"}, "<b>abc</b>X<b>abc</b>"},

		// 典型场景：一个单词是另一个的子串
		{"长短单词嵌套", "abcdef", []string{"abc", "abcde"}, "<b>abcde</b>f"},

		// 边界：单个字符单词
		{"单字符单词", "aXa", []string{"a"}, "<b>a</b>X<b>a</b>"},

		// 典型场景：尾部恰好在结尾结束，末尾 </b> 不能丢
		{"结尾闭合标签", "xyzabc", []string{"abc"}, "xyz<b>abc</b>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBoldTag(tt.s, tt.words); got != tt.want {
				t.Errorf("AddBoldTag(%q, %v) = %q, want %q", tt.s, tt.words, got, tt.want)
			}
		})
	}
}

func BenchmarkAddBoldTag(b *testing.B) {
	// 构造长度 1000 的字符串和 100 个单词
	s := strings.Repeat("abcdefghij", 100)
	words := make([]string, 0, 100)
	for i := 0; i < 10; i++ {
		for j := 1; j <= 10; j++ {
			words = append(words, s[i:i+j])
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddBoldTag(s, words)
	}
}
