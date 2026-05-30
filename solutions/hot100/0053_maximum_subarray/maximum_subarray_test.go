package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  int
	}{
		// LeetCode 官方示例
		{"示例1: [-2,1,-3,4,-1,2,1,-5,4]", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"示例2: [1]", []int{1}, 1},
		{"示例3: [5,4,-1,7,8]", []int{5, 4, -1, 7, 8}, 23},

		// 边界：全为负数
		{"全为负数", []int{-1, -2, -3}, -1},
		{"单个负数", []int{-5}, -5},

		// 边界：全为正数
		{"全为正数", []int{1, 2, 3}, 6},
		{"单个正数", []int{5}, 5},

		// 边界：正负交替
		{"正负交替", []int{-1, 2, -1, 2, -1}, 3},
		{"以负数开头和结尾", []int{-2, 1, -3, 4}, 4},

		// 边界：含0
		{"含0", []int{-1, 0, -2}, 0},
		{"全为0", []int{0, 0, 0}, 0},

		// 较大数组
		{"递增序列", []int{1, 2, 3, 4, 5}, 15},
		{"递减序列", []int{-1, -2, -3, -4, -5}, -1},
		{"正负混合-最大在中间", []int{-2, -1, 3, 4, 5, -10, 2}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("MaxSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums  []int
	}{
		{"len=10", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 3}},
		{"len=100", generateArray(100)},
		{"len=1000", generateArray(1000)},
		{"len=10000", generateArray(10000)},
		{"len=100000", generateArray(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubArray(bm.nums)
			}
		})
	}
}

// generateArray 生成长度为 n 的数组 [-n/2, -n/2+1, ..., n/2]
func generateArray(n int) []int {
	arr := make([]int, n)
	start := -n / 2
	for i := 0; i < n; i++ {
		arr[i] = start + i
	}
	return arr
}
