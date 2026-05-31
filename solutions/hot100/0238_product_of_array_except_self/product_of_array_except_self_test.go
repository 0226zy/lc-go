package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4]", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"示例2: [-1,1,0,-3,3]", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},

		// 边界：全为正数
		{"全为正数", []int{1, 2, 3}, []int{6, 3, 2}},
		{"单个元素（最小长度）", []int{1, 2}, []int{2, 1}},

		// 边界：全为负数
		{"全为负数", []int{-1, -2, -3}, []int{6, 3, 2}},
		{"含负数", []int{-1, 2, -3, 4}, []int{-24, 12, -8, 6}},

		// 边界：含0
		{"单个0", []int{1, 0, 3}, []int{0, 3, 0}},
		{"多个0", []int{0, 1, 0}, []int{0, 0, 0}},
		{"0在开头", []int{0, 2, 3}, []int{6, 0, 0}},
		{"0在结尾", []int{1, 2, 0}, []int{0, 0, 2}},

		// 边界：含1和-1
		{"含1和-1", []int{1, -1, 1, -1}, []int{1, -1, 1, -1}},
		{"全为1", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},

		// 较大数组
		{"长数组", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{"含大数", []int{10, 20, 30}, []int{600, 300, 200}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums  []int
	}{
		{"len=4", []int{1, 2, 3, 4}},
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

// generateArray 生成长度为 n 的数组 [1,2,...,n]
func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}
