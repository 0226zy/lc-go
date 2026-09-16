package longestrepeatingcharacterreplacement

import (
	"strings"
	"testing"
)

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: ABAB k=2",
			"ABAB", 2,
			4,
		},
		{
			"示例2: AABABBA k=1",
			"AABABBA", 1,
			4,
		},

		// 边界：k 为 0，只能统计原有最长重复子串
		{
			"k为0",
			"AABABBA", 0,
			2,
		},

		// 边界：k 等于字符串长度，整串可替换为同一字符
		{
			"k等于串长",
			"ABCDE", 5,
			5,
		},

		// 边界：k 大于需要替换的次数
		{
			"k大于需要替换次数",
			"ABAB", 3,
			4,
		},

		// 边界：全部字符相同
		{
			"全部字符相同",
			"AAAAA", 2,
			5,
		},

		// 边界：全部字符互不相同
		{
			"全部字符互不相同",
			"ABCDE", 2,
			3,
		},

		// 边界：单字符
		{
			"单字符",
			"A", 0,
			1,
		},

		// 边界：最优窗口不在串首
		{
			"最优窗口在串尾",
			"BAAAB", 2,
			5,
		},

		// 边界：答案受限于 k，替换次数刚好用尽
		{
			"替换次数刚好用尽",
			"ABBB", 1,
			4,
		},

		// 较长输入：交替字符，答案为 2k+1 或串长
		{
			"长交替串",
			strings.Repeat("AB", 50), 10,
			21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestRepeatingCharacterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("LongestRepeatingCharacterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestRepeatingCharacterReplacement(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		k    int
	}{
		{"100交替字符", strings.Repeat("AB", 50), 10},
		{"10000交替字符", strings.Repeat("AB", 5000), 100},
		{"100000全相同字符", strings.Repeat("A", 100000), 100000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestRepeatingCharacterReplacement(bm.s, bm.k)
			}
		})
	}
}
