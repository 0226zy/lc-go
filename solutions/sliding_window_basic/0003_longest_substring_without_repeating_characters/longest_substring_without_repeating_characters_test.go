package longestsubstringwithoutrepeatingcharacters

import "testing"

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

		// 边界：空串与单字符
		{"空字符串", "", 0},
		{"单字符", "a", 1},

		// 边界：无重复字符
		{"全部不重复", "abcdefg", 7},
		{"数字与符号混合不重复", "12345!@#$%", 10},

		// 边界：重复字符在两端
		{"首尾相同", "abca", 3},
		{"首尾相同更长", "abcdaefgh", 8},

		// 边界：空格也算字符，"a b"（含空格）无重复
		{"空格参与去重", "a a b", 3},

		// 典型场景
		{"重复字符后接新序列", "dvdf", 3},
		{"交错重复", "abba", 2},
		{"中间断开", "tmmzuxt", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	// 构造一个包含重复字符的长字符串：abc...z 循环
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
	}{
		{"len=10", "abcabcbbab"},
		{"len=100", string(long[:100])},
		{"len=1000", string(long[:1000])},
		{"len=50000", string(long)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstring(bm.s)
			}
		})
	}
}
