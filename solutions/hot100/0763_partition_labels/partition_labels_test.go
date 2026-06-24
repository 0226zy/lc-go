package partitionlabels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "示例1：基本用例",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name: "示例2：全部相同字母",
			s:    "aaaa",
			want: []int{4},
		},
		{
			name: "边界：单字符",
			s:    "a",
			want: []int{1},
		},
		{
			name: "每个字母独立",
			s:    "abcde",
			want: []int{1, 1, 1, 1, 1},
		},
		{
			name: "两端相同字母",
			s:    "abca",
			want: []int{4},
		},
		{
			name: "交叉出现",
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartitionLabels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionLabels(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkPartitionLabels(b *testing.B) {
	s := "ababcbacadefegdehijhklij"
	for i := 0; i < b.N; i++ {
		PartitionLabels(s)
	}
}
