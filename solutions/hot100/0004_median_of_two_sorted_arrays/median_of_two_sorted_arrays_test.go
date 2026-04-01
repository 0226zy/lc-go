package medianoftwosortedarrays

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	const eps = 1e-5

	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		// LeetCode 官方示例
		{"示例1: [1,3]+[2]", []int{1, 3}, []int{2}, 2.0},
		{"示例2: [1,2]+[3,4]", []int{1, 2}, []int{3, 4}, 2.5},

		// 边界：一个数组为空
		{"nums1为空", []int{}, []int{1}, 1.0},
		{"nums2为空", []int{2}, []int{}, 2.0},
		{"nums1为空_偶数", []int{}, []int{1, 2}, 1.5},
		{"nums2为空_奇数", []int{1, 2, 3}, []int{}, 2.0},

		// 边界：单元素数组
		{"单元素+单元素", []int{1}, []int{2}, 1.5},
		{"单元素+单元素_相同", []int{3}, []int{3}, 3.0},

		// 奇偶长度组合
		{"奇+奇=偶", []int{1, 3, 5}, []int{2, 4, 6}, 3.5},
		{"偶+偶=偶", []int{1, 2}, []int{3, 4}, 2.5},
		{"奇+偶=奇", []int{1, 3}, []int{2, 4, 5}, 3.0},

		// 完全不重叠
		{"nums1全小于nums2", []int{1, 2, 3}, []int{4, 5, 6}, 3.5},
		{"nums1全大于nums2", []int{4, 5, 6}, []int{1, 2, 3}, 3.5},

		// 完全重叠（相同数组）
		{"两个相同数组_奇", []int{1, 2, 3}, []int{1, 2, 3}, 2.0},
		{"两个相同数组_偶", []int{1, 2}, []int{1, 2}, 1.5},

		// 含负数
		{"含负数", []int{-3, -1, 0}, []int{-2, 1, 2}, -0.5},
		{"全负数", []int{-5, -3, -1}, []int{-4, -2}, -3.0},

		// 含重复元素
		{"大量重复", []int{1, 1, 1}, []int{1, 1, 1}, 1.0},
		{"部分重复", []int{1, 1, 2}, []int{1, 2, 2}, 1.5},

		// 大小差异悬殊
		{"长短悬殊_短在前", []int{1}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
		{"长短悬殊_短在后", []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1}, 5.5},
		{"单元素插中间", []int{5}, []int{1, 2, 3, 4, 6, 7, 8, 9}, 5.0},

		// 交错分布
		{"交错分布", []int{1, 3, 5, 7}, []int{2, 4, 6, 8}, 4.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(got-tt.want) > eps {
				t.Errorf("FindMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	genSorted := func(n int, offset int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = i*2 + offset
		}
		return s
	}

	benchmarks := []struct {
		name string
		size int
	}{
		{"len=10", 10},
		{"len=100", 100},
		{"len=1000", 1000},
		{"len=10000", 10000},
	}
	for _, bm := range benchmarks {
		nums1 := genSorted(bm.size, 0)
		nums2 := genSorted(bm.size, 1)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMedianSortedArrays(nums1, nums2)
			}
		})
	}
}

func BenchmarkFindMedianSortedArrays_Asymmetric(b *testing.B) {
	short := []int{500}
	long := make([]int, 10000)
	for i := range long {
		long[i] = i
	}

	b.Run("1+10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindMedianSortedArrays(short, long)
		}
	})
}
