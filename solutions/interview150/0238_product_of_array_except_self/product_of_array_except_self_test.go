package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4]", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"示例2: [-1,1,0,-3,3]", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},

		// 边界：长度为 2
		{"长度2", []int{2, 3}, []int{3, 2}},
		{"长度2含0", []int{0, 4}, []int{4, 0}},

		// 边界：含 0
		{"含一个0", []int{1, 0, 3}, []int{0, 3, 0}},
		{"含两个0", []int{0, 2, 0}, []int{0, 0, 0}},

		// 边界：含负数
		{"全为负数", []int{-1, -2, -3}, []int{6, 3, 2}},
		{"正负混合", []int{-1, 2, -3, 4}, []int{-24, 12, -8, 6}},

		// 边界：含 1 和 -1
		{"含1", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"含-1", []int{-1, -1, -1}, []int{1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"len=100", generateArray(100)},
		{"len=1000", generateArray(1000)},
		{"len=10000", generateArray(10000)},
		{"len=100000", generateArray(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProductExceptSelf(bm.nums)
			}
		})
	}
}

// generateArray 生成长度为 n 的数组，元素在 [-3, 10] 之间，避免全 0
func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i%13 - 3
		if arr[i] == 0 {
			arr[i] = 1
		}
	}
	return arr
}
