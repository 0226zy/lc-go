package singlenumber

import "testing"

func TestSingleNumber_OfficialExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "官方示例2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "官方示例3-单元素",
			nums: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSingleNumber_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "负数",
			nums: []int{-1, -2, -1},
			want: -2,
		},
		{
			name: "全负数",
			nums: []int{-4, -3, -4, -3, -5},
			want: -5,
		},
		{
			name: "大数组-目标在开头",
			nums: []int{0, 1, 1, 2, 2, 3, 3},
			want: 0,
		},
		{
			name: "大数组-目标在末尾",
			nums: []int{1, 1, 2, 2, 3, 3, 0},
			want: 0,
		},
		{
			name: "目标为 0",
			nums: []int{7, 7, 0, 1, 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSingleNumber_AllSolutions(t *testing.T) {
	solutions := []func([]int) int{
		SingleNumber,
		SingleNumberHashMap,
	}

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "官方示例2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "官方示例3",
			nums: []int{1},
			want: 1,
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

func BenchmarkSingleNumber_XOR(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	for i := 0; i < b.N; i++ {
		SingleNumber(nums)
	}
}

func BenchmarkSingleNumber_HashMap(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	for i := 0; i < b.N; i++ {
		SingleNumberHashMap(nums)
	}
}
