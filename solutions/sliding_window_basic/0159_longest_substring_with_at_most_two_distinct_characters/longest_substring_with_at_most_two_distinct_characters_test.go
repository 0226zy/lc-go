package longestsubstringwithatmosttwodistinctcharacters

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: eceba", "eceba", 3},
		{"示例2: ccaabbb", "ccaabbb", 5},

		// 边界：空串与单字符
		{"空字符串", "", 0},
		{"单字符", "a", 1},

		// 边界：窗口内只有一种字符
		{"全部相同", "aaaaa", 5},

		// 边界：恰好两种字符
		{"两种字符交替", "ababab", 6},
		{"两种字符分段", "aabbbaa", 7},

		// 典型场景
		{"三种字符取最长", "abaccc", 4},
		{"开头即第三种", "abcabcabc", 2},
		{"最优解在中间", "xyzabbbaz", 5},
		{"交替后追加同字符", "abba", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringTwoDistinct(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstringTwoDistinct(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstringTwoDistinct(b *testing.B) {
	// 构造一个包含多种字符的长字符串：abcd 循环
	long := make([]byte, 0, 100000)
	for len(long) < 100000 {
		for c := byte('a'); c <= byte('d'); c++ {
			long = append(long, c)
		}
	}
	long = long[:100000]

	benchmarks := []struct {
		name string
		s    string
	}{
		{"len=10", "ecebaecbab"},
		{"len=100", string(long[:100])},
		{"len=10000", string(long[:10000])},
		{"len=100000", string(long)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstringTwoDistinct(bm.s)
			}
		})
	}
}
