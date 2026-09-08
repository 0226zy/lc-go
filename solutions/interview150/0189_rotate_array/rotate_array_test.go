package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5,6,7]轮转3位", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"示例2: [-1,-100,3,99]轮转2位", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},

		// 边界：k=0 不轮转
		{"k=0不轮转", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		// 边界：k 等于数组长度（轮转一整圈）
		{"k等于长度", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		// 边界：k 大于数组长度（需取模）
		{"k大于长度", []int{1, 2, 3, 4, 5, 6, 7}, 10, []int{5, 6, 7, 1, 2, 3, 4}},
		// 边界：k 为长度的倍数
		{"k为长度倍数", []int{1, 2, 3, 4}, 8, []int{1, 2, 3, 4}},
		// 边界：单元素
		{"单元素", []int{1}, 5, []int{1}},
		// 边界：全部相同
		{"全部相同", []int{2, 2, 2, 2}, 1, []int{2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			Rotate(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("Rotate(%v, %d) = %v, want %v", tt.nums, tt.k, nums, tt.want)
			}
		})
	}
}

func TestRotateCopy(t *testing.T) {
	// 对照解法跑同一组核心用例
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"示例2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{"k大于长度", []int{1, 2, 3, 4, 5, 6, 7}, 10, []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			RotateCopy(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("RotateCopy(%v, %d) = %v, want %v", tt.nums, tt.k, nums, tt.want)
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	for _, size := range []int{100, 1000, 10000} {
		b.Run(benchName(size), func(b *testing.B) {
			base := make([]int, size)
			for i := range base {
				base[i] = i
			}
			b.Run("翻转法", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					nums := make([]int, size)
					copy(nums, base)
					Rotate(nums, 3)
				}
			})
			b.Run("辅助数组法", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					nums := make([]int, size)
					copy(nums, base)
					RotateCopy(nums, 3)
				}
			})
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
