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
		{"示例1: [1,2,3,4,5,6,7], k=3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"示例2: [-1,-100,3,99], k=2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},

		// 边界：k=0
		{"k=0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"k=0 单个元素", []int{5}, 0, []int{5}},

		// 边界：k 等于数组长度
		{"k等于数组长度", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k大于数组长度", []int{1, 2, 3}, 4, []int{3, 1, 2}},

		// 边界：单元素数组
		{"单元素数组", []int{1}, 1, []int{1}},
		{"单元素数组 k=0", []int{1}, 0, []int{1}},

		// 边界：双元素数组
		{"双元素数组 k=1", []int{1, 2}, 1, []int{2, 1}},
		{"双元素数组 k=2", []int{1, 2}, 2, []int{1, 2}},

		// 边界：全为负数
		{"全为负数", []int{-1, -2, -3, -4}, 2, []int{-3, -4, -1, -2}},

		// 边界：含0
		{"含0", []int{0, 1, 2, 3}, 1, []int{3, 0, 1, 2}},
		{"全为0", []int{0, 0, 0}, 1, []int{0, 0, 0}},

		// 较大数组
		{"长数组 k=1", []int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{"长数组 k=n-1", []int{1, 2, 3, 4, 5}, 4, []int{2, 3, 4, 5, 1}},
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

func BenchmarkRotate(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
		k    int
	}{
		{"len=7,k=3", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"len=100,k=30", generateArray(100), 30},
		{"len=1000,k=300", generateArray(1000), 300},
		{"len=10000,k=3000", generateArray(10000), 3000},
		{"len=100000,k=30000", generateArray(100000), 30000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nums := make([]int, len(bm.nums))
				copy(nums, bm.nums)
				Rotate(nums, bm.k)
			}
		})
	}
}

// generateArray 生成长度为 n 的数组 [1,2,...,n]
func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}
