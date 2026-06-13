package subsets

import (
	"sort"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "官方示例1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "官方示例2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "单元素负数",
			nums: []int{-1},
			want: [][]int{{}, {-1}},
		},
		{
			name: "两个元素",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subsets(tt.nums)
			sortSubsets(got)
			sortSubsets(tt.want)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("Subsets(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSubsetsBitmask(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "官方示例1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "官方示例2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "单元素负数",
			nums: []int{-1},
			want: [][]int{{}, {-1}},
		},
		{
			name: "两个元素",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubsetsBitmask(tt.nums)
			sortSubsets(got)
			sortSubsets(tt.want)
			if !utils.Equal2DIntSlice(got, tt.want) {
				t.Errorf("SubsetsBitmask(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// sortSubsets 对子集集合排序，便于比较
func sortSubsets(sets [][]int) {
	sort.Slice(sets, func(i, j int) bool {
		li, lj := len(sets[i]), len(sets[j])
		if li != lj {
			return li < lj
		}
		for k := 0; k < li; k++ {
			if sets[i][k] != sets[j][k] {
				return sets[i][k] < sets[j][k]
			}
		}
		return false
	})
}

func BenchmarkSubsets(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Subsets(nums)
	}
}

func BenchmarkSubsetsBitmask(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		SubsetsBitmask(nums)
	}
}
