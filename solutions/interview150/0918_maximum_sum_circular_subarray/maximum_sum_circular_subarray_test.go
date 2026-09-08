package maximumsumcircularsubarray

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,-2,3,-2]", []int{1, -2, 3, -2}, 3},
		{"示例2: [5,-3,5]", []int{5, -3, 5}, 10},
		{"示例3: [-3,-2,-3]", []int{-3, -2, -3}, -2},

		// 边界：单元素
		{"单元素为正", []int{3}, 3},
		{"单元素为负", []int{-5}, -5},

		// 边界：全负数（不能返回 0）
		{"全负数", []int{-2, -3, -1}, -1},
		{"全负数含0", []int{-1, -2, 0, -3}, 0},

		// 边界：全正数（最优就是不跨环）
		{"全正数", []int{1, 2, 3}, 6},

		// 边界：跨环取两段
		{"跨环两段", []int{3, -1, 2, -1}, 4},
		{"跨环两端各一个", []int{1, -1, 1}, 2},
		{"最优在环尾和环首", []int{8, -4, 3, -8, 8}, 16},

		// 边界：普通 Kadane 更优
		{"中间一段最优", []int{-2, 4, -5, 4, -5, 9, -20}, 9},
		{"含极大值", []int{-1, 30000, -2, -3}, 30000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubarraySumCircular(tt.nums); got != tt.want {
				t.Errorf("MaxSubarraySumCircular(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSubarraySumCircular(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", []int{5, -3, 5, -2, 1, -2, 3, -2, 4, -1}},
		{"len=100", generateCircularNums(100)},
		{"len=1000", generateCircularNums(1000)},
		{"len=10000", generateCircularNums(10000)},
		{"len=30000", generateCircularNums(30000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxSubarraySumCircular(bm.nums)
			}
		})
	}
}

// generateCircularNums 生成长度为 n 的正负交错数组
func generateCircularNums(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = i % 101
		} else {
			arr[i] = -(i % 83)
		}
	}
	return arr
}
