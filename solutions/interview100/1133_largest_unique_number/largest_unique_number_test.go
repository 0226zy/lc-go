package largestuniquenumber

import "testing"

func TestLargestUniqueNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 最大值不唯一", []int{5, 7, 3, 9, 4, 9, 8, 3, 1}, 8},
		{"示例2: 无唯一数", []int{9, 9, 8, 8}, -1},

		// 边界：单元素
		{"单元素", []int{7}, 7},
		{"单元素为0", []int{0}, 0},

		// 边界：答案是 0（不能误返回 -1）
		{"答案是0", []int{1, 1, 0}, 0},

		// 边界：全部元素相同
		{"全部相同", []int{5, 5, 5, 5}, -1},

		// 边界：值域端点
		{"值域上界1000唯一", []int{1000, 999, 999}, 1000},

		// 典型场景
		{"最大唯一数在末尾", []int{2, 2, 1, 1, 9}, 9},
		{"多个唯一数取最大", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestUniqueNumber(tt.nums); got != tt.want {
				t.Errorf("LargestUniqueNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkLargestUniqueNumber(b *testing.B) {
	// 构造长度为 2000 的数组，值在 [0, 1000] 内循环
	nums := make([]int, 2000)
	for i := range nums {
		nums[i] = i % 1001
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LargestUniqueNumber(nums)
	}
}
