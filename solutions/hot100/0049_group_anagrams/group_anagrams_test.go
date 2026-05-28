package groupanagrams

import (
	"sort"
	"testing"
)

// equalGroupAnagramsResult 比较两个字母异位词分组结果是否等价
// 由于返回的组间顺序和组内顺序均不固定，需对每个组排序后再比较
func equalGroupAnagramsResult(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// 对每个组内部排序
	normalize := func(groups [][]string) [][]string {
		res := make([][]string, len(groups))
		for i, g := range groups {
			copyGroup := make([]string, len(g))
			copy(copyGroup, g)
			sort.Strings(copyGroup)
			res[i] = copyGroup
		}
		// 对组间排序：以组内第一个元素为 key
		sort.Slice(res, func(i, j int) bool {
			if len(res[i]) == 0 && len(res[j]) == 0 {
				return false
			}
			if len(res[i]) == 0 {
				return true
			}
			if len(res[j]) == 0 {
				return false
			}
			return res[i][0] < res[j][0]
		})
		return res
	}

	aNorm := normalize(a)
	bNorm := normalize(b)

	for i := range aNorm {
		if len(aNorm[i]) != len(bNorm[i]) {
			return false
		}
		for j := range aNorm[i] {
			if aNorm[i][j] != bNorm[i][j] {
				return false
			}
		}
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "官方示例",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "空输入",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "单元素",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "无字母异位词",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "全部为同一组异位词",
			strs: []string{"abc", "bca", "cab"},
			want: [][]string{
				{"abc", "bca", "cab"},
			},
		},
		{
			name: "包含空字符串",
			strs: []string{"", "", "a"},
			want: [][]string{
				{"", ""},
				{"a"},
			},
		},
		{
			name: "重复元素",
			strs: []string{"aaa", "aaa", "aaa"},
			want: [][]string{
				{"aaa", "aaa", "aaa"},
			},
		},
		{
			name: "多个不同分组",
			strs: []string{"ab", "ba", "cd", "dc", "ef"},
			want: [][]string{
				{"ab", "ba"},
				{"cd", "dc"},
				{"ef"},
			},
		},
		{
			name: "较长异位词",
			strs: []string{"listen", "silent", "enlist", "hello", "world"},
			want: [][]string{
				{"listen", "silent", "enlist"},
				{"hello"},
				{"world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.strs)
			if !equalGroupAnagramsResult(got, tt.want) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for i := 0; i < b.N; i++ {
		GroupAnagrams(strs)
	}
}
