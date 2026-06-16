package searchinrotatedsortedarray

import "testing"

func TestSearchInRotatedSortedArray_OfficialExamples(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "官方示例1-目标存在",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "官方示例2-目标不存在",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "官方示例3-单元素未命中",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInRotatedSortedArray(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchInRotatedSortedArray_EdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "单元素命中",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "未旋转数组",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "目标为旋转点",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			want:   3,
		},
		{
			name:   "目标在左半部分",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			want:   1,
		},
		{
			name:   "目标在右半部分",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 1,
			want:   5,
		},
		{
			name:   "目标小于最小值",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: -1,
			want:   -1,
		},
		{
			name:   "目标大于最大值",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 8,
			want:   -1,
		},
		{
			name:   "含负数",
			nums:   []int{2, 3, -2, -1, 0, 1},
			target: -2,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInRotatedSortedArray(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestSearchInRotatedSortedArray_AllSolutions(t *testing.T) {
	solutions := []func([]int, int) int{
		SearchInRotatedSortedArray,
		SearchInRotatedSortedArrayFindPivot,
	}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "官方示例1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "官方示例2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "未旋转",
			nums:   []int{1, 2, 3, 4, 5},
			target: 2,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				got := fn(tt.nums, tt.target)
				if got != tt.want {
					t.Errorf("solution[%d](%v, %d) = %d, want %d", i, tt.nums, tt.target, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkSearchInRotatedSortedArray_Main(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	for i := 0; i < b.N; i++ {
		SearchInRotatedSortedArray(nums, target)
	}
}

func BenchmarkSearchInRotatedSortedArray_FindPivot(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	for i := 0; i < b.N; i++ {
		SearchInRotatedSortedArrayFindPivot(nums, target)
	}
}
