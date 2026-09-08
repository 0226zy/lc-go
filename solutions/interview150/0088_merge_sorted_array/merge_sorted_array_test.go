package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: 标准合并", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"示例2: nums2为空", []int{1}, 1, []int{}, 0, []int{1}},
		{"示例3: nums1有效部分为空", []int{0}, 0, []int{1}, 1, []int{1}},

		// 边界：全部来自 nums2
		{"全部来自nums2", []int{0, 0, 0}, 0, []int{1, 2, 3}, 3, []int{1, 2, 3}},
		// 边界：nums1 元素全部大于 nums2
		{"nums1全部更大", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		// 边界：nums1 元素全部小于 nums2
		{"nums1全部更小", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		// 边界：大量重复元素
		{"大量重复元素", []int{1, 1, 2, 0, 0, 0}, 3, []int{1, 2, 2}, 3, []int{1, 1, 1, 2, 2, 2}},
		// 边界：单元素相等
		{"单元素相等", []int{1, 0}, 1, []int{1}, 1, []int{1, 1}},
		// 边界：包含负数
		{"包含负数", []int{-1, 0, 0, 0}, 1, []int{-2, -1, 3}, 3, []int{-2, -1, -1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := make([]int, len(tt.nums1))
			copy(nums1, tt.nums1)
			nums2 := make([]int, len(tt.nums2))
			copy(nums2, tt.nums2)
			Merge(nums1, tt.m, nums2, tt.n)
			if !reflect.DeepEqual(nums1, tt.want) {
				t.Errorf("Merge(%v, %d, %v, %d) = %v, want %v",
					tt.nums1, tt.m, tt.nums2, tt.n, nums1, tt.want)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(sizeName(size), func(b *testing.B) {
			nums2 := make([]int, size)
			for i := range nums2 {
				nums2[i] = 2 * i
			}
			for i := 0; i < b.N; i++ {
				nums1 := make([]int, 2*size)
				for j := 0; j < size; j++ {
					nums1[j] = 2*j + 1
				}
				Merge(nums1, size, nums2, size)
			}
		})
	}
}

func sizeName(n int) string {
	switch n {
	case 100:
		return "len=100"
	case 1000:
		return "len=1000"
	default:
		return "len=10000"
	}
}
