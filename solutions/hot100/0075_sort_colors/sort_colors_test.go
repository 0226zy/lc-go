package sortcolors

import (
	"reflect"
	"testing"
)

func clone(nums []int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	return cpy
}

func TestSortColors_OfficialExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "官方示例1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "官方示例2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := clone(tt.nums)
			SortColors(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func TestSortColors_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "单元素-0",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "单元素-1",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "单元素-2",
			nums: []int{2},
			want: []int{2},
		},
		{
			name: "已全部有序",
			nums: []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "全部相同-0",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "全部相同-1",
			nums: []int{1, 1, 1},
			want: []int{1, 1, 1},
		},
		{
			name: "全部相同-2",
			nums: []int{2, 2, 2},
			want: []int{2, 2, 2},
		},
		{
			name: "逆序",
			nums: []int{2, 2, 1, 1, 0, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "2 在前",
			nums: []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := clone(tt.nums)
			SortColors(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func TestSortColors_AllSolutions(t *testing.T) {
	solutions := []func([]int){
		SortColors,
		SortColorsCounting,
	}

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "官方示例1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "官方示例2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			name: "逆序",
			nums: []int{2, 2, 1, 0},
			want: []int{0, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				nums := clone(tt.nums)
				fn(nums)
				if !reflect.DeepEqual(nums, tt.want) {
					t.Errorf("solution[%d](%v) = %v, want %v", i, tt.nums, nums, tt.want)
				}
			}
		})
	}
}

func BenchmarkSortColors_ThreePointers(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0}
	for i := 0; i < b.N; i++ {
		SortColors(clone(nums))
	}
}

func BenchmarkSortColors_Counting(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0}
	for i := 0; i < b.N; i++ {
		SortColorsCounting(clone(nums))
	}
}
