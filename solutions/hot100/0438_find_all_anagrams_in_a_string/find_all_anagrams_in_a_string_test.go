package findallanagramsinastring

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		{"示例1", "cbaebabacd", "abc", []int{0, 6}},
		{"示例2", "abab", "ab", []int{0, 1, 2}},
		{"无匹配", "abcdefg", "xyz", []int{}},
		{"s长度小于p", "ab", "abc", []int{}},
		{"单个字符匹配", "aa", "a", []int{0, 1}},
		{"全部相同字符", "aaaaa", "aa", []int{0, 1, 2, 3}},
		{"空p", "abc", "", []int{}},
		{"边界情况-长度相等", "abc", "abc", []int{0}},
		{"边界情况-长度相等不匹配", "abc", "def", []int{}},
		{"复杂情况", "abacbabc", "abc", []int{1, 2, 3, 5}},
		{"重复字符", "baa", "aa", []int{1}},
		{"长字符串", "ababababab", "ab", []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.s, tt.p)
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("FindAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}

func BenchmarkFindAnagrams(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		p    string
	}{
		{"小规模", "cbaebabacd", "abc"},
		{"中等规模", "abababababababababab", "ab"},
		{"大规模", "aaaaaaaaaaaaaaaaaaaa", "aaa"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindAnagrams(bm.s, bm.p)
			}
		})
	}
}
