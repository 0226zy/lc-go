package longestsubstring

import (
	"strings"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: abcabcbb", "abcabcbb", 3},
		{"示例2: bbbbb", "bbbbb", 1},
		{"示例3: pwwkew", "pwwkew", 3},

		// 边界：空字符串和单字符
		{"空字符串", "", 0},
		{"单字符", "a", 1},

		// 边界：全相同字符
		{"全相同2个", "aa", 1},
		{"全相同10个", "aaaaaaaaaa", 1},

		// 边界：全不同字符
		{"全不同", "abcdefg", 7},
		{"全不同2个", "ab", 2},

		// 重复字符在首部
		{"重复在首部", "aab", 2},
		// 重复字符在尾部
		{"重复在尾部", "abb", 2},
		// 重复字符在中间
		{"重复在中间", "abba", 2},

		// 含空格
		{"含空格", "a b c a b c", 3},
		{"纯空格", "   ", 1},

		// 含数字和符号
		{"数字和字母混合", "abc123abc", 6},
		{"含特殊符号", "a!@#a!@#", 4},

		// 滑动窗口跳跃场景
		{"窗口跳跃", "abcdeafghij", 10},
		{"交替重复", "ababababab", 2},
		{"首尾相同中间不同", "axyza", 4},

		// 较长字符串
		{"26个小写字母", "abcdefghijklmnopqrstuvwxyz", 26},
		{"重复的26字母", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"len=10_全不同", "abcdefghij"},
		{"len=100_重复模式", strings.Repeat("abcdefghij", 10)},
		{"len=1000_重复模式", strings.Repeat("abcdefghij", 100)},
		{"len=10000_重复模式", strings.Repeat("abcdefghij", 1000)},
		{"len=1000_全相同", strings.Repeat("a", 1000)},
		{"len=1000_全不同ASCII", genASCIIString(1000)},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstring(bm.s)
			}
		})
	}
}

func genASCIIString(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(33 + i%94)) // 可打印 ASCII 33~126
	}
	return sb.String()
}
