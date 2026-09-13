package palindromepartitioning

import (
	"sort"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "官方示例1",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "官方示例2",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "全部字符相同",
			s:    "aaa",
			want: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "无任何回文子串",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "两个字符构成回文",
			s:    "aa",
			want: [][]string{{"a", "a"}, {"aa"}},
		},
		{
			name: "整体即为回文串",
			s:    "aba",
			want: [][]string{{"a", "b", "a"}, {"aba"}},
		},
		{
			name: "较长混合字符串",
			s:    "aabaa",
			want: [][]string{
				{"a", "a", "b", "a", "a"},
				{"a", "a", "b", "aa"},
				{"a", "aba", "a"},
				{"aa", "b", "a", "a"},
				{"aa", "b", "aa"},
				{"aabaa"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Partition(tt.s)
			if !equalPartitions(normalize(got), normalize(tt.want)) {
				t.Errorf("Partition(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// normalize 将每个分割方案序列化后整体排序，使二维切片的比较与方案顺序无关
func normalize(partitions [][]string) []string {
	keys := make([]string, len(partitions))
	for i, p := range partitions {
		key := ""
		for _, sub := range p {
			key += sub + "\x00" // 用不可见字符分隔，避免拼接歧义
		}
		keys[i] = key
	}
	sort.Strings(keys)
	return keys
}

// equalPartitions 比较两组序列化后的分割方案是否完全一致
func equalPartitions(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Partition("aaaaaaaaaaaaaaaa")
	}
}
