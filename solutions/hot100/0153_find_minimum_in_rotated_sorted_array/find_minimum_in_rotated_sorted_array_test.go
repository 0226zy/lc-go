package findminimuminrotatedsortedarray

import "testing"

func TestFindMin_OfficialExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "官方示例2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "官方示例3-未旋转",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMin(tt.nums)
			if got != tt.want {
				t.Errorf("FindMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFindMin_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "单元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "两个元素-旋转",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "两个元素-未旋转",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "最小值在末尾",
			nums: []int{3, 4, 5, 6, 1, 2},
			want: 1,
		},
		{
			name: "最小值在开头",
			nums: []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "含负数",
			nums: []int{0, 1, 2, -5, -3, -1},
			want: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMin(tt.nums)
			if got != tt.want {
				t.Errorf("FindMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFindMin_AllSolutions(t *testing.T) {
	solutions := []func([]int) int{
		FindMin,
		FindMinLinear,
	}

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "官方示例2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "官方示例3",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				got := fn(tt.nums)
				if got != tt.want {
					t.Errorf("solution[%d](%v) = %d, want %d", i, tt.nums, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkFindMin_Main(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		FindMin(nums)
	}
}

func BenchmarkFindMin_Linear(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		FindMinLinear(nums)
	}
}
