package longestsubstringwithatmostkdistinctcharacters

import "testing"

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: eceba k=2", "eceba", 2, 3},
		{"示例2: aa k=1", "aa", 1, 2},

		// 边界：k 为 0
		{"k为0", "eceba", 0, 0},
		{"k为0且单字符", "a", 0, 0},

		// 边界：单元素
		{"单字符 k=1", "a", 1, 1},

		// 边界：k 不小于不同字符总数，整个字符串都合法
		{"k等于不同字符总数", "abcde", 5, 5},
		{"k大于不同字符总数", "aaabbb", 10, 6},

		// 典型场景
		{"最优解在中间", "abaccc", 2, 4},
		{"最优解在开头", "aabbccdd", 2, 4},
		{"最优解在结尾", "abcdaaa", 2, 4},
		{"交替字符", "abababab", 2, 8},
		{"交错三种字符", "abcadcacac", 3, 8},
		{"收缩后重新扩张", "abbbbcdddd", 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringKDistinct(tt.s, tt.k); got != tt.want {
				t.Errorf("LengthOfLongestSubstringKDistinct(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstringKDistinct(b *testing.B) {
	// 构造一个包含大量重复字符的长字符串：abc...z 循环
	long := make([]byte, 0, 50000)
	for len(long) < 50000 {
		for c := byte('a'); c <= byte('z'); c++ {
			long = append(long, c)
		}
	}
	long = long[:50000]

	benchmarks := []struct {
		name string
		s    string
		k    int
	}{
		{"len=10,k=2", "ecebaeceba", 2},
		{"len=100,k=2", string(long[:100]), 2},
		{"len=1000,k=5", string(long[:1000]), 5},
		{"len=50000,k=10", string(long), 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstringKDistinct(bm.s, bm.k)
			}
		})
	}
}
