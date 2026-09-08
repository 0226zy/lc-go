package groupanagrams

import (
	"sort"
	"testing"
)

// sortResult 对分组结果排序，使分组间、组内元素顺序确定，便于断言比较
func sortResult(result [][]string) [][]string {
	for _, group := range result {
		sort.Strings(group)
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) != len(result[j]) {
			return len(result[i]) < len(result[j])
		}
		return result[i][0] < result[j][0]
	})
	return result
}

// expectedGroups 构造期望结果并排序
func expectedGroups(groups ...[]string) [][]string {
	return sortResult(groups)
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		// LeetCode 官方示例
		{
			"示例1: 多组异位词混合",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expectedGroups(
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			),
		},
		{"示例2: 空串一组", []string{""}, expectedGroups([]string{""})},
		{"示例3: 单个字符一组", []string{"a"}, expectedGroups([]string{"a"})},

		// 边界：空输入
		{"空输入", []string{}, [][]string{}},

		// 边界：全是异位词
		{"全是异位词", []string{"abc", "bca", "cab"},
			expectedGroups([]string{"abc", "bca", "cab"})},

		// 边界：两两不同，无法分组
		{"两两不同", []string{"ab", "cd", "ef"},
			expectedGroups([]string{"ab"}, []string{"cd"}, []string{"ef"})},

		// 边界：重复字符串
		{"含重复字符串", []string{"aa", "aa", "bb"},
			expectedGroups([]string{"aa", "aa"}, []string{"bb"})},

		// 边界：不同长度
		{"不同长度", []string{"a", "ab", "ba"},
			expectedGroups([]string{"a"}, []string{"ab", "ba"})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortResult(GroupAnagrams(tt.strs))
			if !equalGroups(got, tt.want) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}

// equalGroups 比较两个分组结果是否完全一致（均已排序）
func equalGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func BenchmarkGroupAnagrams(b *testing.B) {
	benchmarks := []struct {
		name string
		strs []string
	}{
		{"官方示例", []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		{"100个短串", generateStrs(100)},
		{"1000个短串", generateStrs(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GroupAnagrams(bm.strs)
			}
		})
	}
}

// generateStrs 生成 n 个长度为 5 的随机风格字符串（由 abc 组成）
func generateStrs(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		bytes := make([]byte, 5)
		for j := range bytes {
			bytes[j] = byte('a' + (i+j)%3)
		}
		strs[i] = string(bytes)
	}
	return strs
}
