package blockplacementqueries

import (
	"reflect"
	"testing"
)

func TestGetResults(t *testing.T) {
	tests := []struct {
		name    string
		queries [][]int
		want    []bool
	}{
		{
			name:    "官方示例1",
			queries: [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			want:    []bool{false, true, true},
		},
		{
			name:    "官方示例2",
			queries: [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			want:    []bool{true, true, false},
		},
		{
			name:    "无障碍物查询",
			queries: [][]int{{2, 5, 5}, {2, 5, 6}},
			want:    []bool{true, false},
		},
		{
			name:    "单个障碍物",
			queries: [][]int{{1, 3}, {2, 5, 3}, {2, 5, 2}, {2, 2, 2}},
			want:    []bool{true, true, true},
		},
		{
			name:    "障碍物在边界",
			queries: [][]int{{1, 1}, {2, 1, 1}, {2, 1, 2}},
			want:    []bool{true, false},
		},
		{
			name:    "多个障碍物紧密排列",
			queries: [][]int{{1, 2}, {1, 4}, {1, 6}, {2, 6, 2}, {2, 6, 1}},
			want:    []bool{true, true},
		},
		{
			name:    "查询x小于障碍物",
			queries: [][]int{{1, 5}, {2, 3, 3}, {2, 3, 2}},
			want:    []bool{true, true},
		},
		{
			name:    "大间隔",
			queries: [][]int{{1, 2}, {1, 10}, {2, 10, 8}, {2, 10, 9}},
			want:    []bool{true, false},
		},
		{
			name:    "sz等于间隔",
			queries: [][]int{{1, 3}, {2, 3, 3}, {2, 3, 4}},
			want:    []bool{true, false},
		},
		{
			name:    "物块恰好接触障碍物",
			queries: [][]int{{1, 4}, {2, 4, 4}, {2, 4, 5}},
			want:    []bool{true, false},
		},
		{
			name:    "先查询后放置障碍物",
			queries: [][]int{{2, 9, 9}, {1, 5}, {2, 9, 6}},
			want:    []bool{true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetResults(tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResults(%v) = %v, want %v", tt.queries, got, tt.want)
			}
		})
	}
}

func BenchmarkGetResults(b *testing.B) {
	queries := [][]int{
		{1, 2},
		{2, 3, 3},
		{2, 3, 1},
		{2, 2, 2},
	}
	for i := 0; i < b.N; i++ {
		GetResults(queries)
	}
}
