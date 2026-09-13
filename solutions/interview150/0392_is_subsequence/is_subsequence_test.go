package issubsequence

import (
	"strings"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: abc 是 ahbgdc 的子序列", "abc", "ahbgdc", true},
		{"示例2: axc 不是 ahbgdc 的子序列", "axc", "ahbgdc", false},

		// 边界：空串
		{"空 s 是任意 t 的子序列", "", "ahbgdc", true},
		{"空 s 与空 t", "", "", true},
		{"非空 s 与空 t", "a", "", false},

		// 边界：完全相等 / 单字符
		{"两串完全相等", "abc", "abc", true},
		{"单字符匹配", "a", "a", true},
		{"单字符不匹配", "a", "b", false},
		{"s 比 t 长", "abcd", "abc", false},

		// 常规：字符分散匹配
		{"首尾字符匹配", "ac", "abc", true},
		{"连续子串也是子序列", "bcd", "abcdef", true},
		{"跳过中间字符", "ace", "abcde", true},
		{"相对顺序错误 aec", "aec", "abcde", false},
		{"重复字符需要足够次数", "aaa", "aba", false},
		{"重复字符刚好够", "aaa", "aaaba", true},

		// 全部小写字母覆盖
		{"s 用尽 t 的部分字符", "leetcode", "leetcoding", false},
		{"末尾才匹配完", "c", "abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.s, tt.t); got != tt.want {
				t.Errorf("IsSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestIsSubsequenceStress(t *testing.T) {
	longT := strings.Repeat("a", 10000)
	t.Run("一万个a中匹配aaa", func(t *testing.T) {
		if !IsSubsequence("aaa", longT) {
			t.Errorf("IsSubsequence(\"aaa\", 一万a) = false, want true")
		}
	})
	t.Run("一万个a中匹配b失败", func(t *testing.T) {
		if IsSubsequence("b", longT) {
			t.Errorf("IsSubsequence(\"b\", 一万a) = true, want false")
		}
	})
}

func BenchmarkIsSubsequence(b *testing.B) {
	longT := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 400) // 约 10400 字符
	benchmarks := []struct {
		name string
		s, t string
	}{
		{"短串匹配", "abc", "ahbgdc"},
		{"短串不匹配", "axc", "ahbgdc"},
		{"长t匹配末尾", "z", longT},
		{"长t整串作为子序列", "abcdefghijklmnopqrstuvwxyz", longT},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsSubsequence(bm.s, bm.t)
			}
		})
	}
}
