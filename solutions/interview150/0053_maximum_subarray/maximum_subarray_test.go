package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [-2,1,-3,4,-1,2,1,-5,4]", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"示例2: [1]", []int{1}, 1},
		{"示例3: [5,4,-1,7,8]", []int{5, 4, -1, 7, 8}, 23},

		// 边界：单元素
		{"单元素为负数", []int{-1}, -1},
		{"单元素为最小值", []int{-10000}, -10000},

		// 边界：全负数
		{"全负数", []int{-2, -1, -3, -4}, -1},
		{"全负数且递减", []int{-5, -4, -3, -2, -1}, -1},

		// 边界：全正数
		{"全正数", []int{1, 2, 3, 4}, 10},

		// 边界：正负交错
		{"正负交错", []int{1, -1, 1, -1, 1}, 1},
		{"先正后负拖尾", []int{3, -2, -1, -10}, 3},
		{"中间低谷后反弹", []int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},

		// 边界：答案就是整个数组
		{"累加最大", []int{1, 2, -1, 2, 3}, 7},
		{"极大值混入", []int{-1, 0, -2, 10000, -10000}, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.nums); got != tt.want {
				t.Errorf("MaxSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 3}},
		{"len=100", generateNums(100)},
		{"len=1000", generateNums(1000)},
		{"len=10000", generateNums(10000)},
		{"len=100000", generateNums(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubArray(bm.nums)
			}
		})
	}
}

// generateNums 生成长度为 n 的正负交错数组
func generateNums(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = i % 97
		} else {
			arr[i] = -(i % 89)
		}
	}
	return arr
}
