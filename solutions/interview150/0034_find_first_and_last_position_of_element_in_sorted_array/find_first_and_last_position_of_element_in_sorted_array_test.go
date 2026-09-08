package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		// LeetCode 官方示例
		{"示例1: 命中中间重复元素", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"示例2: 目标不存在", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"示例3: 空数组", []int{}, 0, []int{-1, -1}},

		// 边界：单元素
		{"单元素命中", []int{1}, 1, []int{0, 0}},
		{"单元素未命中", []int{1}, 2, []int{-1, -1}},

		// 边界：全部元素都等于 target
		{"全部等于target", []int{2, 2, 2, 2}, 2, []int{0, 3}},
		{"两元素都等于target", []int{3, 3}, 3, []int{0, 1}},

		// 边界：target 在数组头部
		{"target在头部且重复", []int{5, 5, 7, 7, 8, 8, 10}, 5, []int{0, 1}},
		{"target在头部单个", []int{5, 7, 7, 8, 8, 10}, 5, []int{0, 0}},

		// 边界：target 在数组尾部
		{"target在尾部且重复", []int{5, 7, 7, 8, 8, 10, 10}, 10, []int{5, 6}},
		{"target在尾部单个", []int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},

		// 边界：target 小于最小值 / 大于最大值
		{"target小于最小值", []int{5, 7, 7, 8, 8, 10}, 4, []int{-1, -1}},
		{"target大于最大值", []int{5, 7, 7, 8, 8, 10}, 11, []int{-1, -1}},

		// 边界：只出现一次的中间元素
		{"target仅出现一次", []int{1, 2, 3, 4, 5}, 3, []int{2, 2}},

		// 边界：两个相同元素相邻
		{"相邻重复元素", []int{1, 4, 4, 5}, 4, []int{1, 2}},

		// 边界：负数
		{"含负数命中", []int{-5, -3, -3, 0, 2}, -3, []int{1, 2}},
		{"含负数未命中", []int{-5, -3, -3, 0, 2}, -4, []int{-1, -1}},

		// 较大数组
		{"长数组命中重复段", []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7}, 4, []int{3, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkSearchRange(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=10", 10},
		{"n=1000", 1000},
		{"n=100000", 100000},
		{"n=1000000", 1000000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			nums := make([]int, bm.n)
			for i := 0; i < bm.n; i++ {
				nums[i] = i / 2 // 制造大量重复元素
			}
			target := bm.n / 4
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				SearchRange(nums, target)
			}
		})
	}
}

func BenchmarkLowerBound(b *testing.B) {
	nums := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums[i] = i / 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lowerBound(nums, 500000)
	}
}
