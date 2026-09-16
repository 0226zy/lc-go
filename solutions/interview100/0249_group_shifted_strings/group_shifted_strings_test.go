package groupshiftedstrings

import (
	"sort"
	"testing"
)

// equal2DStringSliceUnordered 比较两个二维字符串切片是否相等（分组顺序、组内顺序均无关）
func equal2DStringSliceUnordered(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	normalize := func(groups [][]string) []string {
		keys := make([]string, len(groups))
		for i, g := range groups {
			sorted := make([]string, len(g))
			copy(sorted, g)
			sort.Strings(sorted)
			key := ""
			for _, s := range sorted {
				key += s + "\x00"
			}
			keys[i] = key
		}
		sort.Strings(keys)
		return keys
	}
	ka, kb := normalize(a), normalize(b)
	for i := range ka {
		if ka[i] != kb[i] {
			return false
		}
	}
	return true
}

func TestGroupStrings(t *testing.T) {
	tests := []struct {
		name    string
		strings []string
		want    [][]string
	}{
		// LeetCode 官方示例
		{
			"示例1: 多组混合",
			[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			[][]string{{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}},
		},
		{
			"示例2: 单字符",
			[]string{"a"},
			[][]string{{"a"}},
		},

		// 边界：全部单字符聚为一组
		{
			"多个单字符同组",
			[]string{"a", "b", "z"},
			[][]string{{"a", "b", "z"}},
		},

		// 边界：长度不同不可能互为移位字符串
		{
			"长度不同必不同组",
			[]string{"ab", "abc"},
			[][]string{{"ab"}, {"abc"}},
		},

		// 典型：环绕移位 'z'->'a'，差值序列为 (1,24) 的归为一组
		{
			"跨 z-a 环绕",
			[]string{"abz", "bca", "yza"},
			[][]string{{"abz", "bca"}, {"yza"}},
		},

		// 典型：回文式差值
		{
			"差值回文",
			[]string{"aba", "cdc", "abc"},
			[][]string{{"aba", "cdc"}, {"abc"}},
		},

		// 全部相同字符串
		{
			"全部相同",
			[]string{"dd", "dd"},
			[][]string{{"dd", "dd"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupStrings(tt.strings)
			if !equal2DStringSliceUnordered(got, tt.want) {
				t.Errorf("GroupStrings(%v) = %v, want %v", tt.strings, got, tt.want)
			}
		})
	}
}

func BenchmarkGroupStrings(b *testing.B) {
	// 构造 200 个长度为 50 的字符串
	strs := make([]string, 0, 200)
	base := []byte("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx")
	for i := 0; i < 200; i++ {
		strs = append(strs, string(base[:50]))
		base[0] = 'a' + byte(i%26)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupStrings(strs)
	}
}
