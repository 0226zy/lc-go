package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,1,2]", []int{1, 1, 2}, 2, []int{1, 2}},
		{"示例2: 多重复元素", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},

		// 边界：空数组
		{"空数组", []int{}, 0, []int{}},
		// 边界：单元素
		{"单元素", []int{1}, 1, []int{1}},
		// 边界：全部相同
		{"全部相同", []int{2, 2, 2, 2}, 1, []int{2}},
		// 边界：全部不同
		{"全部不同", []int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		// 边界：含负数
		{"含负数", []int{-3, -3, -2, -1, -1}, 3, []int{-3, -2, -1}},
		// 边界：大量重复
		{"大量重复", []int{1, 1, 1, 1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := RemoveDuplicates(nums)
			if got != tt.wantLen {
				t.Errorf("RemoveDuplicates(%v) 长度 = %d, want %d", tt.nums, got, tt.wantLen)
			}
			gotNums := make([]int, got)
			copy(gotNums, nums[:got])
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("RemoveDuplicates(%v) 去重结果 = %v, want %v", tt.nums, gotNums, tt.wantNums)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(benchName(size), func(b *testing.B) {
			base := make([]int, size)
			for i := range base {
				base[i] = i / 3 // 制造重复
			}
			for i := 0; i < b.N; i++ {
				nums := make([]int, size)
				copy(nums, base)
				RemoveDuplicates(nums)
			}
		})
	}
}

func benchName(n int) string {
	switch n {
	case 100:
		return "len=100"
	case 1000:
		return "len=1000"
	default:
		return "len=10000"
	}
}
