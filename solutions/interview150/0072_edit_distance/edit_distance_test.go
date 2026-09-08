package editdistance

import (
	"strings"
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: horse转ros", "horse", "ros", 3},
		{"示例2: intention转execution", "intention", "execution", 5},

		// 边界：空串
		{"两个都是空串", "", "", 0},
		{"word1为空串需全部插入", "", "abc", 3},
		{"word2为空串需全部删除", "abc", "", 3},

		// 边界：单字符
		{"相同单字符", "a", "a", 0},
		{"不同单字符需替换一次", "a", "b", 1},

		// 完全相同的字符串
		{"相同字符串无需操作", "leetcode", "leetcode", 0},

		// 无公共字符，需要全部替换或增删
		{"完全不同且等长", "abc", "def", 3},
		{"完全不同且不等长", "ab", "xyzw", 4},

		// 只需插入或只需删除
		{"只需插入一个字符", "abc", "abcd", 1},
		{"只需删除一个字符", "abcd", "abc", 1},
		{"只需在中间插入", "ac", "abc", 1},

		// 只需替换
		{"只需替换一个字符", "abc", "adc", 1},

		// 长公共子串交错
		{"经典例子dinitrophenylhydrazine", "dinitrophenylhydrazine", "acetylphenylhydrazine", 6},
		{"distance转springbok", "distance", "springbok", 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("MinDistance(%q, %q) = %v, want %v", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

// TestMinDistanceStress 压力场景：两个长度均为约束上限 500 的字符串
func TestMinDistanceStress(t *testing.T) {
	// 完全相同的 500 个字符，编辑距离为 0
	same := strings.Repeat("a", 500)
	t.Run("500个相同字符距离为0", func(t *testing.T) {
		if got := MinDistance(same, same); got != 0 {
			t.Errorf("MinDistance(500个a, 500个a) = %v, want 0", got)
		}
	})

	// 全 a 与全 b，需要替换 500 次
	t.Run("500个a转500个b需500次替换", func(t *testing.T) {
		if got := MinDistance(same, strings.Repeat("b", 500)); got != 500 {
			t.Errorf("MinDistance(500个a, 500个b) = %v, want 500", got)
		}
	})

	// 全 a 与 a 前缀加 b 后缀：只需在末尾插入一个 b
	t.Run("只需在末尾插入一个字符", func(t *testing.T) {
		if got := MinDistance(same, same+"b"); got != 1 {
			t.Errorf("MinDistance(500个a, 500个a加b) = %v, want 1", got)
		}
	})
}

func BenchmarkMinDistance(b *testing.B) {
	benchmarks := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"官方示例intention转execution", "intention", "execution"},
		{"100个a转100个b", strings.Repeat("a", 100), strings.Repeat("b", 100)},
		{"500个a转500个b_约束上限", strings.Repeat("a", 500), strings.Repeat("b", 500)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinDistance(bm.word1, bm.word2)
			}
		})
	}
}
