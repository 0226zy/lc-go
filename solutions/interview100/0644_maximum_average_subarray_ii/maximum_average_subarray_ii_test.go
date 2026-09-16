package maximumaveragesubarrayii

import (
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	const eps = 1e-5
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		// LeetCode 官方示例
		{"示例1: 含正负混合", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"示例2: 单元素", []int{5}, 1, 5.0},

		// 边界：k 等于数组长度，只能取整个数组
		{"k等于数组长度", []int{1, 2, 3, 4}, 4, 2.5},

		// 边界：全部元素相同
		{"全部元素相同", []int{3, 3, 3, 3}, 2, 3.0},

		// 边界：全负数，答案为负
		{"全负数取最大者", []int{-5, -2, -3}, 1, -2.0},
		{"全负数k大于1", []int{-5, -2, -3, -1}, 3, -2.0},

		// 边界：k 为 1，答案就是最大元素
		{"k为1取最大元素", []int{4, 0, 9, 1}, 1, 9.0},

		// 典型场景：更长子数组拉低平均，短子数组更优
		{"短子数组更优", []int{1, 1, 100, 1, 1}, 2, 50.5},
		{"答案在数组中部", []int{-1, -1, 8, 8, -1}, 2, 8.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxAverage(tt.nums, tt.k); math.Abs(got-tt.want) > eps {
				t.Errorf("FindMaxAverage(%v, %d) = %v, want %v (误差超过 %v)", tt.nums, tt.k, got, tt.want, eps)
			}
		})
	}
}

func BenchmarkFindMaxAverage(b *testing.B) {
	// 构造长度为 10^4 的伪随机数组
	nums := make([]int, 10000)
	seed := 12345
	for i := range nums {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		nums[i] = seed%20001 - 10000
	}

	benchmarks := []struct {
		name string
		nums []int
		k    int
	}{
		{"n=100,k=10", nums[:100], 10},
		{"n=1000,k=100", nums[:1000], 100},
		{"n=10000,k=1", nums, 1},
		{"n=10000,k=5000", nums, 5000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMaxAverage(bm.nums, bm.k)
			}
		})
	}
}
