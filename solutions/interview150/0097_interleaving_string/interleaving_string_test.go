package interleavingstring

import (
	"strings"
	"testing"
)

var isInterleaveCases = []struct {
	name string
	s1   string
	s2   string
	s3   string
	want bool
}{
	// LeetCode 官方示例
	{name: "示例1：aabcc与dbbca交错成aadbbcbcac", s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac", want: true},
	{name: "示例2：aabcc与dbbca无法交错成aadbbbaccc", s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc", want: false},
	{name: "示例3：三个空串", s1: "", s2: "", s3: "", want: true},

	// 边界：空字符串参与
	{name: "s1为空s2等于s3", s1: "", s2: "abc", s3: "abc", want: true},
	{name: "s1为空s2不等于s3", s1: "", s2: "abc", s3: "abd", want: false},
	{name: "s2为空s1等于s3", s1: "abc", s2: "", s3: "abc", want: true},
	{name: "s2为空s1不等于s3", s1: "abc", s2: "", s3: "abd", want: false},
	{name: "s1为空s3为空s2非空", s1: "", s2: "a", s3: "", want: false},

	// 边界：长度不匹配
	{name: "长度不匹配直接失败", s1: "a", s2: "b", s3: "abc", want: false},
	{name: "s3太短", s1: "ab", s2: "cd", s3: "abc", want: false},

	// 边界：单字符
	{name: "单字符交错成功", s1: "a", s2: "b", s3: "ab", want: true},
	{name: "单字符交错交换顺序", s1: "a", s2: "b", s3: "ba", want: true},
	{name: "单字符交错失败", s1: "a", s2: "b", s3: "aa", want: false},

	// 易错：字符顺序必须保持
	{name: "破坏s1内部顺序", s1: "ab", s2: "c", s3: "bac", want: false},
	{name: "破坏s2内部顺序", s1: "c", s2: "ab", s3: "bac", want: false},

	// 易错：相同字符导致多种取法
	{name: "相同字符可交错", s1: "aa", s2: "aa", s3: "aaaa", want: true},
	{name: "相同字符交错失败", s1: "aa", s2: "aa", s3: "aaab", want: false},
	{name: "全相同字符数量不匹配", s1: "aaa", s2: "aa", s3: "aaaaa", want: true},

	// 一般场景
	{name: "交替取字符", s1: "abc", s2: "def", s3: "adbecf", want: true},
	{name: "先取完s1再取s2", s1: "abc", s2: "def", s3: "abcdef", want: true},
	{name: "先取完s2再取s1", s1: "abc", s2: "def", s3: "defabc", want: true},
	{name: "字符对不上", s1: "abc", s2: "def", s3: "adbecx", want: false},
}

func TestIsInterleave(t *testing.T) {
	for _, tt := range isInterleaveCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("IsInterleave(%q, %q, %q) = %v, want %v",
					tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}

func TestIsInterleaveOptimized(t *testing.T) {
	for _, tt := range isInterleaveCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterleaveOptimized(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("IsInterleaveOptimized(%q, %q, %q) = %v, want %v",
					tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}

// TestIsInterleaveLarge 压力场景：最大约束长度（s1、s2 各 100 个字符，s3 共 200 个字符）
func TestIsInterleaveLarge(t *testing.T) {
	s1 := strings.Repeat("a", 100)
	s2 := strings.Repeat("a", 100)

	t.Run("百级长度全相同字符交错成功", func(t *testing.T) {
		s3 := strings.Repeat("a", 200)
		if !IsInterleave(s1, s2, s3) {
			t.Errorf("IsInterleave 期望 true, 得到 false")
		}
		if !IsInterleaveOptimized(s1, s2, s3) {
			t.Errorf("IsInterleaveOptimized 期望 true, 得到 false")
		}
	})

	t.Run("百级长度末位字符不匹配", func(t *testing.T) {
		s3 := strings.Repeat("a", 199) + "b"
		if IsInterleave(s1, s2, s3) {
			t.Errorf("IsInterleave 期望 false, 得到 true")
		}
		if IsInterleaveOptimized(s1, s2, s3) {
			t.Errorf("IsInterleaveOptimized 期望 false, 得到 true")
		}
	})
}

var isInterleaveBenchmarks = []struct {
	name string
	s1   string
	s2   string
	s3   string
}{
	{"官方示例成功", "aabcc", "dbbca", "aadbbcbcac"},
	{"官方示例失败", "aabcc", "dbbca", "aadbbbaccc"},
	{"百级长度全相同字符", strings.Repeat("a", 100), strings.Repeat("a", 100), strings.Repeat("a", 200)},
	{"百级长度交替字符", strings.Repeat("ab", 50), strings.Repeat("ba", 50), strings.Repeat("abba", 50)},
}

func BenchmarkIsInterleave(b *testing.B) {
	for _, bm := range isInterleaveBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsInterleave(bm.s1, bm.s2, bm.s3)
			}
		})
	}
}

func BenchmarkIsInterleaveOptimized(b *testing.B) {
	for _, bm := range isInterleaveBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsInterleaveOptimized(bm.s1, bm.s2, bm.s3)
			}
		})
	}
}
