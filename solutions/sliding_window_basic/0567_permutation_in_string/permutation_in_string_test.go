package permutationinstring

import (
	"strings"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		// LeetCode 官方示例
		{
			"示例1: 包含排列 ba",
			"ab", "eidbaooo",
			true,
		},
		{
			"示例2: 不包含任何排列",
			"ab", "eidboaoo",
			false,
		},

		// 边界：s1 比 s2 长
		{
			"s1 比 s2 长",
			"abc", "ab",
			false,
		},

		// 边界：单字符
		{
			"单字符匹配",
			"a", "a",
			true,
		},
		{
			"单字符不匹配",
			"a", "b",
			false,
		},
		{
			"单字符出现在长串末尾",
			"z", "abcz",
			true,
		},

		// 边界：s1 与 s2 长度相等
		{
			"长度相等且互为排列",
			"abc", "bca",
			true,
		},
		{
			"长度相等但不是排列",
			"abc", "abd",
			false,
		},
		{
			"完全相同的字符串",
			"abc", "abc",
			true,
		},

		// 匹配位置：窗口出现在 s2 开头
		{
			"排列在 s2 开头",
			"ab", "baoooo",
			true,
		},
		// 匹配位置：窗口出现在 s2 末尾
		{
			"排列在 s2 末尾",
			"ab", "oooba",
			true,
		},

		// 边界：s1 全为同一字符
		{
			"s1 全为重复字符且匹配",
			"aaa", "caaab",
			true,
		},
		{
			"s1 全为重复字符且不匹配",
			"aaaa", "aaabaaa",
			false,
		},

		// 一般情况：字符种类多、窗口需多次滑动
		{
			"多次滑动后命中",
			"xyz", "afdgzyxksldfm",
			true,
		},
		{
			"交替字符中计数命中",
			"aab", "ababab",
			true,
		},
		{
			"字符种类相同但计数不匹配",
			"adc", "dcba",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CheckInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func BenchmarkCheckInclusion(b *testing.B) {
	benchmarks := []struct {
		name  string
		s1Len int
		s2Len int
	}{
		{"s1长度10_s2长度100", 10, 100},
		{"s1长度100_s2长度1000", 100, 1000},
		{"s1长度1000_s2长度10000", 1000, 10000},
	}

	for _, bm := range benchmarks {
		s1 := strings.Repeat("a", bm.s1Len)
		s2 := strings.Repeat("b", bm.s2Len-bm.s1Len) + s1 // 排列出现在末尾，最坏情况滑完全程
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CheckInclusion(s1, s2)
			}
		})
	}
}
