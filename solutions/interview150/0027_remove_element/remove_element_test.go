package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantLen  int
		wantNums []int
	}{
		// LeetCode 官方示例
		{"示例1: [3,2,2,3]移除3", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"示例2: 混合数组移除2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},

		// 边界：全部移除
		{"全部移除", []int{1, 1, 1}, 1, 0, []int{}},
		// 边界：一个都不移除
		{"一个都不移除", []int{1, 2, 3}, 4, 3, []int{1, 2, 3}},
		// 边界：单元素需移除
		{"单元素需移除", []int{1}, 1, 0, []int{}},
		// 边界：单元素保留
		{"单元素保留", []int{1}, 2, 1, []int{1}},
		// 边界：空数组
		{"空数组", []int{}, 1, 0, []int{}},
		// 边界：值在开头和结尾
		{"首尾为移除值", []int{2, 1, 2}, 2, 1, []int{1}},
		// 边界：包含负数与0
		{"含负数与0", []int{-1, 0, -1, 2}, -1, 2, []int{0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := RemoveElement(nums, tt.val)
			if got != tt.wantLen {
				t.Errorf("RemoveElement(%v, %d) 长度 = %d, want %d", tt.nums, tt.val, got, tt.wantLen)
			}
			// 校验保留下来的元素集合正确（只关心前 got 个）
			gotNums := make([]int, got)
			copy(gotNums, nums[:got])
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("RemoveElement(%v, %d) 保留元素 = %v, want %v", tt.nums, tt.val, gotNums, tt.wantNums)
			}
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(benchName(size), func(b *testing.B) {
			base := make([]int, size)
			for i := range base {
				base[i] = i % 10
			}
			for i := 0; i < b.N; i++ {
				nums := make([]int, size)
				copy(nums, base)
				RemoveElement(nums, 5)
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
