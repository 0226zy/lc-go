package sorts

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "普通数组",
			nums: []int{3, 1, 4, 1, 5, 9, 2, 6},
			want: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name: "已排序数组",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "逆序数组",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "包含负数",
			nums: []int{-3, 0, 2, -1, 5},
			want: []int{-3, -1, 0, 2, 5},
		},
		{
			name: "空数组",
			nums: []int{},
			want: []int{},
		},
		{
			name: "单元素",
			nums: []int{42},
			want: []int{42},
		},
		{
			name: "全部相同",
			nums: []int{2, 2, 2, 2},
			want: []int{2, 2, 2, 2},
		},
		{
			name: "两个元素",
			nums: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "大数组",
			nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.nums)
			if !utils.EqualIntSlice(tt.nums, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func TestMergeSortDesc(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "普通数组",
			nums: []int{3, 1, 4, 1, 5, 9, 2, 6},
			want: []int{9, 6, 5, 4, 3, 2, 1, 1},
		},
		{
			name: "已排序数组",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "逆序数组",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "包含负数",
			nums: []int{-3, 0, 2, -1, 5},
			want: []int{5, 2, 0, -1, -3},
		},
		{
			name: "空数组",
			nums: []int{},
			want: []int{},
		},
		{
			name: "单元素",
			nums: []int{42},
			want: []int{42},
		},
		{
			name: "全部相同",
			nums: []int{2, 2, 2, 2},
			want: []int{2, 2, 2, 2},
		},
		{
			name: "两个元素",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "大数组",
			nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortDesc(tt.nums)
			if !utils.EqualIntSlice(tt.nums, tt.want) {
				t.Errorf("MergeSortDesc() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 8, 4, 2, 7, 1, 6}
		MergeSort(nums)
	}
}

func BenchmarkMergeSortDesc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{5, 3, 8, 4, 2, 7, 1, 6}
		MergeSortDesc(nums)
	}
}
