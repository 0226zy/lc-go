package findfirstoccurrence

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		// LeetCode 官方示例
		{"示例1: sadbutsad 中找 sad", "sadbutsad", "sad", 0},
		{"示例2: leetcode 中找 leeto", "leetcode", "leeto", -1},

		// 边界：单字符
		{"单字符相等", "a", "a", 0},
		{"单字符不等", "a", "b", -1},
		{"文本单字符模式多字符", "a", "aa", -1},

		// 边界：模式串等于文本串
		{"模式串等于文本串", "abc", "abc", 0},
		{"模式串等于文本串且更长", "abc", "abcd", -1},

		// 边界：匹配出现在开头 / 结尾
		{"匹配在开头", "hello", "he", 0},
		{"匹配在结尾", "hello", "lo", 3},
		{"匹配在中间", "mississippi", "ssis", 2},

		// 边界：多次出现，返回第一个
		{"多次出现返回第一个", "sadbutsad", "sad", 0},
		{"ababab 中找 aba", "ababab", "aba", 0},
		{"aaaaaa 中找 aa", "aaaaaa", "aa", 0},

		// 边界：需要利用 next 数组回退的复杂失配场景
		{"经典 KMP 回退场景", "ababcabcacbab", "abcac", 5},
		{"部分前缀重叠失配", "aabaaabaaac", "aabaaac", 4},

		// 边界：几乎匹配但最后失配
		{"末尾失配", "abcababcab", "abcabd", -1},

		// 较长文本
		{"长文本随机匹配", "qweasdzxcqweasdzxcqweasdzxcfghjkl", "fghjkl", 27},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("StrStr(%q, %q) = %d, want %d", tt.haystack, tt.needle, got, tt.want)
			}
		})
	}
}

// TestStrStrConsistency 校验 KMP、标准库与暴力三种实现结果一致
func TestStrStrConsistency(t *testing.T) {
	haystacks := []string{"sadbutsad", "leetcode", "aaaaaa", "ababcabcacbab", "mississippi", "", "a", "abababab"}
	needles := []string{"sad", "leeto", "aa", "abcac", "issip", "", "a", "abab"}

	for _, h := range haystacks {
		for _, n := range needles {
			got, want := StrStr(h, n), StrStrBuiltin(h, n)
			if got != want {
				t.Errorf("StrStr(%q, %q) = %d, StrStrBuiltin = %d", h, n, got, want)
			}
			if got != StrStrBrute(h, n) {
				t.Errorf("StrStr(%q, %q) = %d, StrStrBrute = %d", h, n, got, StrStrBrute(h, n))
			}
		}
	}
}

// 生成最坏情况输入：文本全为 'a'，模式串末尾一个 'b'
func generateWorstCase(m, n int) (string, string) {
	haystack := make([]byte, m)
	for i := range haystack {
		haystack[i] = 'a'
	}
	needle := make([]byte, n)
	for i := range needle {
		needle[i] = 'a'
	}
	needle[n-1] = 'b'
	return string(haystack), string(needle)
}

func BenchmarkStrStr(b *testing.B) {
	sizes := []struct {
		name string
		m, n int
	}{
		{"短文本 KMP", 9, 3}, // haystack = "sadbutsad", needle = "sad"
		{"最坏情况 m=n=100", 100, 100},
		{"最坏情况 m=n=10000", 10000, 10000},
	}

	for _, bm := range sizes {
		b.Run(bm.name, func(b *testing.B) {
			if bm.m == 9 {
				for i := 0; i < b.N; i++ {
					StrStr("sadbutsad", "sad")
				}
				return
			}
			haystack, needle := generateWorstCase(bm.m, bm.n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				StrStr(haystack, needle)
			}
		})
	}
}

func BenchmarkStrStrBrute(b *testing.B) {
	haystack, needle := generateWorstCase(10000, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrStrBrute(haystack, needle)
	}
}
