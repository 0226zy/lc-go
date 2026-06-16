package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
)

func TestSearchRange_OfficialExamples(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "官方示例1-目标存在且重复",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "官方示例2-目标不存在",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "官方示例3-空数组",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchRange_EdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "单元素命中",
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			name:   "单元素未命中",
			nums:   []int{1},
			target: 2,
			want:   []int{-1, -1},
		},
		{
			name:   "目标小于最小值",
			nums:   []int{2, 2},
			target: 1,
			want:   []int{-1, -1},
		},
		{
			name:   "目标大于最大值",
			nums:   []int{2, 2},
			target: 3,
			want:   []int{-1, -1},
		},
		{
			name:   "全部元素相同",
			nums:   []int{2, 2, 2, 2, 2},
			target: 2,
			want:   []int{0, 4},
		},
		{
			name:   "目标在数组开头",
			nums:   []int{1, 1, 2, 3, 4},
			target: 1,
			want:   []int{0, 1},
		},
		{
			name:   "目标在数组末尾",
			nums:   []int{1, 2, 3, 4, 4},
			target: 4,
			want:   []int{3, 4},
		},
		{
			name:   "含负数",
			nums:   []int{-5, -3, -3, -3, 0, 1},
			target: -3,
			want:   []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchRange_AllSolutions(t *testing.T) {
	solutions := []func([]int, int) []int{
		SearchRange,
		SearchRangeBinarySearchAndExpand,
	}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "官方示例1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "官方示例2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "全部相同",
			nums:   []int{2, 2, 2},
			target: 2,
			want:   []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				got := fn(tt.nums, tt.target)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("solution[%d](%v, %d) = %v, want %v", i, tt.nums, tt.target, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkSearchRange_Main(b *testing.B) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	for i := 0; i < b.N; i++ {
		SearchRange(nums, target)
	}
}

func BenchmarkSearchRange_BinarySearchAndExpand(b *testing.B) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	for i := 0; i < b.N; i++ {
		SearchRangeBinarySearchAndExpand(nums, target)
	}
}
