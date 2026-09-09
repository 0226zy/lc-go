package editdistance

import (
	"strings"
	"testing"
)

var minDistanceCases = []struct {
	name  string
	word1 string
	word2 string
	want  int
}{
	// LeetCode 官方示例
	{name: "示例1：horse转ros", word1: "horse", word2: "ros", want: 3},
	{name: "示例2：intention转execution", word1: "intention", word2: "execution", want: 5},

	// 边界：空串
	{name: "两个都是空串", word1: "", word2: "", want: 0},
	{name: "word1为空串需全部插入", word1: "", word2: "abc", want: 3},
	{name: "word2为空串需全部删除", word1: "abc", word2: "", want: 3},

	// 边界：单字符
	{name: "相同单字符", word1: "a", word2: "a", want: 0},
	{name: "不同单字符需替换一次", word1: "a", word2: "b", want: 1},

	// 完全相同的字符串
	{name: "相同字符串无需操作", word1: "leetcode", word2: "leetcode", want: 0},

	// 无公共字符，需要全部替换或增删
	{name: "完全不同且等长", word1: "abc", word2: "def", want: 3},
	{name: "完全不同且不等长", word1: "ab", word2: "xyzw", want: 4},

	// 只需插入或只需删除
	{name: "只需插入一个字符", word1: "abc", word2: "abcd", want: 1},
	{name: "只需删除一个字符", word1: "abcd", word2: "abc", want: 1},
	{name: "只需在中间插入", word1: "ac", word2: "abc", want: 1},

	// 只需替换
	{name: "只需替换一个字符", word1: "abc", word2: "adc", want: 1},

	// 删除加插入
	{name: "删除加插入", word1: "ab", word2: "bc", want: 2},

	// 长公共子串交错
	{name: "经典例子dinitrophenylhydrazine", word1: "dinitrophenylhydrazine", word2: "acetylphenylhydrazine", want: 6},
	{name: "distance转springbok", word1: "distance", word2: "springbok", want: 9},
}

func TestMinDistance(t *testing.T) {
	for _, tt := range minDistanceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func TestMinDistanceOptimized(t *testing.T) {
	for _, tt := range minDistanceCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistanceOptimized(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistanceOptimized(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

// TestMinDistanceStress 压力场景：两个长度均为约束上限 500 的字符串
func TestMinDistanceStress(t *testing.T) {
	same := strings.Repeat("a", 500)
	stressCases := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		// 完全相同的 500 个字符，编辑距离为 0
		{name: "500个相同字符距离为0", word1: same, word2: same, want: 0},
		// 全 a 与全 b，需要替换 500 次
		{name: "500个a转500个b需500次替换", word1: same, word2: strings.Repeat("b", 500), want: 500},
		// 全 a 与 a 前缀加 b 后缀：只需在末尾插入一个 b
		{name: "只需在末尾插入一个字符", word1: same, word2: same + "b", want: 1},
	}
	for _, tt := range stressCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistance = %d, want %d", got, tt.want)
			}
			if got := MinDistanceOptimized(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistanceOptimized = %d, want %d", got, tt.want)
			}
		})
	}
}

var minDistanceBenchmarks = []struct {
	name  string
	word1 string
	word2 string
}{
	{"官方示例intention转execution", "intention", "execution"},
	{"100个a转100个b", strings.Repeat("a", 100), strings.Repeat("b", 100)},
	{"500个a转500个b_约束上限", strings.Repeat("a", 500), strings.Repeat("b", 500)},
}

func BenchmarkMinDistance(b *testing.B) {
	for _, bm := range minDistanceBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinDistance(bm.word1, bm.word2)
			}
		})
	}
}

func BenchmarkMinDistanceOptimized(b *testing.B) {
	for _, bm := range minDistanceBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinDistanceOptimized(bm.word1, bm.word2)
			}
		})
	}
}
