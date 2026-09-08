package kthlargestelementinanarray

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [3,2,1,5,6,4] k=2", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"示例2: 含重复元素", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},

		// 边界：单元素
		{"单元素", []int{1}, 1, 1},
		{"单元素负数", []int{-1}, 1, -1},

		// 边界：k=1 即最大值，k=n 即最小值
		{"k=1 最大值", []int{1, 2, 3, 4, 5}, 1, 5},
		{"k=n 最小值", []int{1, 2, 3, 4, 5}, 5, 1},

		// 边界：全部相同
		{"全部相同", []int{7, 7, 7, 7}, 3, 7},

		// 边界：负数混合
		{"负数混合", []int{-3, -1, -2, -4}, 2, -2},
		{"正负混合", []int{-1, 5, 0, -3, 2}, 3, 0},

		// 边界：已排序数组
		{"升序数组", []int{1, 2, 3, 4, 5}, 2, 4},
		{"降序数组", []int{5, 4, 3, 2, 1}, 2, 4},

		// 边界：重复元素居多
		{"重复元素居多", []int{2, 2, 2, 1, 1, 1}, 4, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest(tt.nums, tt.k); got != tt.want {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// TestFindKthLargestBySort 验证排序对照解法与堆解法结果一致
func TestFindKthLargestBySort(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	if got := FindKthLargestBySort(nums, 4); got != 4 {
		t.Errorf("FindKthLargestBySort = %d, want 4", got)
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}
	for _, n := range sizes {
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rand.Intn(2*n) - n
		}
		k := n / 10
		b.Run("堆解法 n="+strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindKthLargest(nums, k)
			}
		})
		b.Run("排序解法 n="+strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindKthLargestBySort(nums, k)
			}
		})
	}
}
