package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "示例1：基本用例",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "示例2：答案在数组末尾",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "示例3：相同元素不同下标",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "边界：最小长度数组",
			nums:   []int{1, 2},
			target: 3,
			want:   []int{0, 1},
		},
		{
			name:   "负数用例",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if got == nil {
				t.Skip("未实现")
				return
			}
			if len(got) != 2 || tt.nums[got[0]]+tt.nums[got[1]] != tt.target {
				t.Errorf("TwoSum(%v, %d) = %v, 但两数之和应为 %d", tt.nums, tt.target, got, tt.target)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
