package interleavingstring

import (
	"strings"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		s3   string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: aabcc与dbbca交错成aadbbcbcac", "aabcc", "dbbca", "aadbbcbcac", true},
		{"示例2: aabcc与dbbca无法交错成aadbbbaccc", "aabcc", "dbbca", "aadbbbaccc", false},
		{"示例3: 三个空串", "", "", "", true},

		// 边界：空字符串参与
		{"s1为空s2等于s3", "", "abc", "abc", true},
		{"s1为空s2不等于s3", "", "abc", "abd", false},
		{"s2为空s1等于s3", "abc", "", "abc", true},
		{"s2为空s1不等于s3", "abc", "", "abd", false},
		{"s1为空s3为空s2非空", "", "a", "", false},

		// 边界：长度不匹配
		{"长度不匹配直接失败", "a", "b", "abc", false},
		{"s3太短", "ab", "cd", "abc", false},

		// 边界：单字符
		{"单字符交错成功", "a", "b", "ab", true},
		{"单字符交错交换顺序", "a", "b", "ba", true},
		{"单字符交错失败", "a", "b", "aa", false},

		// 易错：字符顺序必须保持
		{"破坏s1内部顺序", "ab", "c", "bac", false},
		{"破坏s2内部顺序", "c", "ab", "bac", false},

		// 易错：相同字符导致多种取法
		{"相同字符可交错", "aa", "aa", "aaaa", true},
		{"相同字符交错失败", "aa", "aa", "aaab", false},
		{"全相同字符数量不匹配", "aaa", "aa", "aaaaa", true},

		// 一般场景
		{"交替取字符", "abc", "def", "adbecf", true},
		{"先取完s1再取s2", "abc", "def", "abcdef", true},
		{"先取完s2再取s1", "abc", "def", "defabc", true},
		{"字符对不上", "abc", "def", "adbecx", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("IsInterleave(%q, %q, %q) = %v, want %v",
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
	})

	t.Run("百级长度末位字符不匹配", func(t *testing.T) {
		s3 := strings.Repeat("a", 199) + "b"
		if IsInterleave(s1, s2, s3) {
			t.Errorf("IsInterleave 期望 false, 得到 true")
		}
	})
}

func BenchmarkIsInterleave(b *testing.B) {
	benchmarks := []struct {
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

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsInterleave(bm.s1, bm.s2, bm.s3)
			}
		})
	}
}
