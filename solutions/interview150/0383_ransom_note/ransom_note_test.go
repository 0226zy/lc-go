package ransomnote

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		// LeetCode 官方示例
		{"示例1: ransomNote 所需字符不足", "a", "b", false},
		{"示例2: ransomNote 所需字符不足", "aa", "ab", false},
		{"示例3: ransomNote 可完整构造", "aa", "aab", true},

		// 边界：空串
		{"空 ransomNote", "", "abc", true},
		{"空 magazine 且非空 ransomNote", "a", "", false},

		// 边界：相等串
		{"完全相同", "abc", "abc", true},
		{"完全相同的重复字符", "aaa", "aaa", true},

		// 边界：字符种类不同
		{"magazine 含 ransomNote 没有的字符", "abc", "abcde", true},
		{"ransomNote 含 magazine 没有的字符", "abcdef", "abcde", false},

		// 边界：字符数量恰好
		{"数量恰好够用", "aab", "aba", true},
		{"数量差一个", "aabb", "abac", false},

		// 边界：单字符
		{"单字符可构造", "z", "z", true},
		{"单字符不可构造", "z", "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("CanConstruct(%q, %q) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}

func BenchmarkCanConstruct(b *testing.B) {
	benchmarks := []struct {
		name       string
		ransomNote string
		magazine   string
	}{
		{"短串", "aa", "aab"},
		{"长串可构造", generateRepeat('a', 500) + generateRepeat('b', 500), generateRepeat('b', 500) + generateRepeat('a', 500)},
		{"长串不可构造", generateRepeat('a', 1000), generateRepeat('b', 1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanConstruct(bm.ransomNote, bm.magazine)
			}
		})
	}
}

// generateRepeat 生成 n 个重复字符 c 组成的字符串
func generateRepeat(c byte, n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = c
	}
	return string(s)
}
