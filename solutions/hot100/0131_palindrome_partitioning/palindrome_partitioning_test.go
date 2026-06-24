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
			name: "全部相同字符",
			s:    "aaa",
			want: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "无重复",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "两个字符回文",
			s:    "aa",
			want: [][]string{{"a", "a"}, {"aa"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Partition(tt.s)
			if !equalPartitions(got, tt.want) {
				t.Errorf("Partition(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestPartitionWithTable(t *testing.T) {
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
			name: "全部相同字符",
			s:    "aaa",
			want: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "无重复",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "两个字符回文",
			s:    "aa",
			want: [][]string{{"a", "a"}, {"aa"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionWithTable(tt.s)
			if !equalPartitions(got, tt.want) {
				t.Errorf("PartitionWithTable(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// equalPartitions 比较两个分割方案是否相同（顺序无关）
func equalPartitions(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	sa := make([][]string, len(a))
	copy(sa, a)
	sb := make([][]string, len(b))
	copy(sb, b)
	sort.Slice(sa, func(i, j int) bool {
		for k := 0; k < len(sa[i]) && k < len(sa[j]); k++ {
			if sa[i][k] != sa[j][k] {
				return sa[i][k] < sa[j][k]
			}
		}
		return len(sa[i]) < len(sa[j])
	})
	sort.Slice(sb, func(i, j int) bool {
		for k := 0; k < len(sb[i]) && k < len(sb[j]); k++ {
			if sb[i][k] != sb[j][k] {
				return sb[i][k] < sb[j][k]
			}
		}
		return len(sb[i]) < len(sb[j])
	})
	for i := range sa {
		if len(sa[i]) != len(sb[i]) {
			return false
		}
		for k := range sa[i] {
			if sa[i][k] != sb[i][k] {
				return false
			}
		}
	}
	return true
}

func BenchmarkPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Partition("aaaaaaaaaaaaaaaa")
	}
}

func BenchmarkPartitionWithTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartitionWithTable("aaaaaaaaaaaaaaaa")
	}
}
